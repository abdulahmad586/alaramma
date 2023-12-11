import "package:quran/quran.dart";
import "package:sqlite_async/sqlite_async.dart";

List<String> tableCreation = [
  'CREATE TABLE chapters (id INTEGER PRIMARY KEY AUTOINCREMENT, chapterNumber INTEGER, totalVerses INTEGER, englishName TEXT, arabicName TEXT, pageStart INTEGER, pageEnd INTEGER, totalPages INTEGER, revelationPlace TEXT)',
  'CREATE TABLE verses (id INTEGER PRIMARY KEY AUTOINCREMENT, verseNumber INTEGER, chapterNumber INTEGER, juzNumber INTEGER, pageNumber INTEGER, arabicText TEXT, englishText TEXT, keywords TEXT)',
  'CREATE TABLE translations (id INTEGER PRIMARY KEY AUTOINCREMENT, translatorName TEXT, language TEXT, content TEXT, verseNumber INTEGER, chapterNumber INTEGER)',
  'CREATE TABLE tafseer (id INTEGER PRIMARY KEY AUTOINCREMENT, scholarName TEXT, language TEXT, content TEXT, verseNumber INTEGER, chapterNumber INTEGER)',
  'CREATE TABLE audioVerses (id INTEGER PRIMARY KEY AUTOINCREMENT, reciterName, audioUrl, verseNumber, chapterNumber)'
];

void main(List<String> args) async {
  SqliteMigration createTableMigrations = SqliteMigration(1, (tx) async {
    for (final tableString in tableCreation) {
      await tx.execute(tableString);
    }
    print("Finished preparing tables");
  });

  final db = SqliteDatabase(path: 'quran_database/test.db');
  await createTableMigrations.fn(db);
  await populateChapters(db);
  await populateVerses(db);

  db.close();
}

Future<void> populateChapters(SqliteDatabase db) async {
  for (int i = 1; i < 115; i++) {
    String sql =
        'INSERT INTO chapters (chapterNumber, totalVerses, englishName, arabicName, pageStart, pageEnd, totalPages, revelationPlace) VALUES (?,?,?,?,?,?,?,?)';
    await db.execute(sql, [
      i,
      getVerseCount(i),
      getSurahNameEnglish(i),
      getSurahNameArabic(i),
      getSurahPages(i).first,
      getSurahPages(i).last,
      getSurahPages(i).length,
      getPlaceOfRevelation(i)
    ]);
    print("Finished ${getSurahName(i)}");
  }
}

Future<void> populateVerses(SqliteDatabase db) async {
  for (int i = 1; i < 115; i++) {
    int chapterNumber = i;
    for (int ayahNumber = 1;
        ayahNumber < getVerseCount(chapterNumber) + 1;
        ayahNumber++) {
      String sqlVerse =
          'INSERT INTO verses (verseNumber, chapterNumber, juzNumber, pageNumber, arabicText, englishText, keywords) VALUES (?, ?, ?, ?, ?, ?, ?)';
      await db.execute(sqlVerse, [
        ayahNumber,
        chapterNumber,
        getJuzNumber(chapterNumber, ayahNumber),
        getPageNumber(chapterNumber, ayahNumber),
        getVerse(chapterNumber, ayahNumber),
        '',
        ''
      ]);

      String sqlTranslations =
          'INSERT INTO translations (verseNumber, chapterNumber, content, language, translatorName) VALUES (?,?,?,?,?)';
      await db.execute(sqlTranslations, [
        ayahNumber,
        chapterNumber,
        getVerseTranslation(chapterNumber, ayahNumber),
        "English",
        "Saheeh International"
      ]);

      String sqlAudioVerses =
          'INSERT INTO audioVerses (verseNumber, chapterNumber, audioUrl, reciterName) VALUES (?,?,?,?)';
      await db.execute(sqlAudioVerses, [
        ayahNumber,
        chapterNumber,
        getAudioURLByVerse(chapterNumber, ayahNumber),
        ''
      ]);
    }
  }
}
