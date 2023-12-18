const sqlite3 = require('sqlite3').verbose();
const mongoose = require('mongoose');
const { Chapter, Verse, Translation, Tafseer, AudioVerse } = require('./model'); // Import your Mongoose models

// MongoDB connection
mongoose.connect('mongodb+srv://ddld93:1234567890@cluster0.fljiocn.mongodb.net/qurandb');
const dbMongo = mongoose.connection;
dbMongo.on('error', console.error.bind(console, 'MongoDB connection error:'));

// SQLite connection
const dbSQLite = new sqlite3.Database('../quran_database/quran.db', sqlite3.OPEN_READONLY, (err) => {
    if (err) {
        console.error(err.message);
    }
    console.log('Connected to the SQLite database.');
});

// Helper function to map chapter number to MongoDB ObjectId
const chapterNumberToIdMap = new Map();

// Function to migrate data
const migrateData = async () => {
    // Migrate Chapters
    await new Promise((resolve, reject) => {
        dbSQLite.each("SELECT * FROM chapters", async (err, row) => {
            if (err) {
                console.error(err);
                reject(err);
                return;
            }
            const chapter = new Chapter({
                chapterNumber: row.chapterNumber,
                totalVerses: row.totalVerses,
                englishName: row.englishName,
                arabicName: row.arabicName,
                pageStart: row.pageStart,
                pageEnd: row.pageEnd,
                totalPages: row.totalPages,
                revelationPlace: row.revelationPlace
            });

            const savedChapter = await chapter.save();
            // chapterNumberToIdMap.set(row.chapterNumber, savedChapter._id);
            resolve();
        });
    });

    // Migrate Verses
    await new Promise((resolve, reject) => {
        dbSQLite.each("SELECT * FROM verses", async (err, row) => {
            if (err) {
                console.error(err);
                reject(err);
                return;
            }
            const verse = new Verse({
                verseNumberGlobal: row.id,
                verseNumber: row.verseNumber,
                chapterNumber: row.chapterNumber,
                juzNumber: row.juzNumber,
                pageNumber: row.pageNumber,
                arabicText: row.arabicText,
                englishText: row.englishText,
                keywords: row.keywords ? row.keywords.split(',') : [],
                // Assuming keywordsEmbedding is an array of numbers
                // keywordsEmbedding: row.keywordsEmbedding ? JSON.parse(row.keywordsEmbedding) : []
            });

            await verse.save();
            resolve();
        });
    });

    // Migrate Translation
    await new Promise((resolve, reject) => {
        dbSQLite.each("SELECT * FROM translations", async (err, row) => {
            if (err) {
                console.error(err);
                reject(err);
                return;
            }
            const translation = new Translation({
                translatorName: row.translatorName,
                language: row.language,
                content: row.content,
                verseNumber: row.verseNumber,
                chapterNumber: row.chapterNumber
                // Assuming keywordsEmbedding is an array of numbers
                // keywordsEmbedding: row.keywordsEmbedding ? JSON.parse(row.keywordsEmbedding) : []
            });

            await translation.save();
            resolve();
        });
    });

     // Migrate AudioVerse
     await new Promise((resolve, reject) => {
        dbSQLite.each("SELECT * FROM audioVerses", async (err, row) => {
            if (err) {
                console.error(err);
                reject(err);
                return;
            }
            const audioVerse = new AudioVerse({
                reciterName: row.reciterName,
                audioUrl: row.audioUrl,
                verseNumber: row.verseNumber,
                chapterNumber: row.chapterNumber
                // Assuming keywordsEmbedding is an array of numbers
                // keywordsEmbedding: row.keywordsEmbedding ? JSON.parse(row.keywordsEmbedding) : []
            });

            await audioVerse.save();
            resolve();
        });
    });




    console.log('Migration completed.');
};

migrateData().then(() => {
    dbSQLite.close(); // Close SQLite connection
    // dbMongo.close(); // Close MongoDB connection
}).catch(console.error);
