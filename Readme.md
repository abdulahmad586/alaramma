Welcome to Alaramma! "Quranic Guide"
Your personal Alaramma companion app designed to enhance your journey in understanding, reciting, and connecting with the profound teachings of the Quran. This app serves as your virtual mentor, offering comprehensive support in your quest to explore and deepen your relationship with the sacred text.

**Database structure**
------------------

**chapters**
id, surahNumber, totalVerses, englishName, arabicName. pageStart, pageEnd, totalPages, revelation

**verses**
id, verseNumber, surahNumber, pageNumber, arabicText, englishText, keywords

**translation**
id, translatorName, language, content, verseNumber, surahNumber

**tafseer**
id, scholarName, language, content, verseNumber, surahNumber

**audioVerses**
id, reciterName, audioUrl, verseNumber, surahNumber