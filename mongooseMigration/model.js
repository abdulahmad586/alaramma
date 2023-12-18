const mongoose = require('mongoose');

const ChapterSchema = new mongoose.Schema({
    chapterNumber: { type: Number, required: true, index: { unique: true } },
    totalVerses: { type: Number, required: true },
    englishName: { type: String, required: true },
    arabicName: { type: String, required: true },
    pageStart: Number,
    pageEnd: Number,
    totalPages: Number,
    revelationPlace: String
}, { timestamps: true });

const VerseSchema = new mongoose.Schema({
    verseNumberGlobal: { type: Number, required: true, },
    verseNumber: { type: Number, required: true, },
    chapterNumber: { type: Number, ref: 'Chapter', required: true },
    juzNumber: { type: Number },
    pageNumber: { type: Number },
    arabicText: { type: String },
    englishText: { type: String },
    keywords: [{ type: Number }],
    keywordsEmbedding: [{ type: String }],
}, { timestamps: true });

const TranslationSchema = new mongoose.Schema({
    // verseId: { type: mongoose.Schema.Types.ObjectId, ref: 'Verse', required: true },
    translatorName: { type: String, required: true },
    language: { type: String, required: true },
    content: { type: String, required: true },
    contentEmbedding: { type: [Number] },
    verseNumber: { type: Number, ref: 'Verse', required: true },
    chapterNumber: { type: Number, ref: 'Chapter', required: true }
}, { timestamps: true });

const TafseerSchema = new mongoose.Schema({
    // verseId: { type: mongoose.Schema.Types.ObjectId, ref: 'Verse', required: true },
    scholarName: { type: String, required: true },
    language: { type: String, required: true },
    content: { type: String },
    contentEmbedding: [{ type: Number }],
    verseNumber: { type: Number, ref: 'Verse', required: true },
    chapterNumber: { type: Number, ref: 'Chapter', required: true }
}, { timestamps: true });

const AudioVerseSchema = new mongoose.Schema({
    reciterName: { type: String },
    audioUrl: { type: String, required: true },
    verseNumber: { type: Number, ref: 'Verse', required: true },
    chapterNumber: { type: Number, ref: 'Chapter', required: true }
}, { timestamps: true });

// Indexing
ChapterSchema.index({ chapterNumber: 1 }); // Assuming frequent queries by chapterNumber
VerseSchema.index({ chapterNumber: 1, verseNumber: 1 }); // Composite index for frequent querying

const Chapter = mongoose.model('Chapter', ChapterSchema);
const Verse = mongoose.model('Verse', VerseSchema);
const Translation = mongoose.model('Translation', TranslationSchema);
const Tafseer = mongoose.model('Tafseer', TafseerSchema);
const AudioVerse = mongoose.model('AudioVerse', AudioVerseSchema);

module.exports = { Chapter, Verse, Translation, Tafseer, AudioVerse };
