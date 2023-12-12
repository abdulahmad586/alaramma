package query

import (
    "encoding/json"
    "log"
)

// sqlQueries is the exported variable containing the SQL query statements.
var SqlQueries sqlQueriesStruct

// sqlQueriesStruct is the internal structure that holds the SQL queries and manipulations.
type sqlQueriesStruct struct {
    queries       queryStatements    `json:"queries"`
    manipulations manipulationStatements `json:"manipulations"`
}

type queryStatements struct {
    listAllChapters        string `json:"listAllChapters"`
    getChapterByNumber     string `json:"getChapterByNumber"`
    listVersesInChapter    string `json:"listVersesInChapter"`
    getSpecificVerse       string `json:"getSpecificVerse"`
    searchVersesByKeywords string `json:"searchVersesByKeywords"`
    listTranslationsForVerse string `json:"listTranslationsForVerse"`
    getTranslationByLanguage string `json:"getTranslationByLanguage"`
    listTafseerForVerse    string `json:"listTafseerForVerse"`
    getTafseerByScholar    string `json:"getTafseerByScholar"`
    listAudioVersionsOfVerse string `json:"listAudioVersionsOfVerse"`
    getAudioByReciter      string `json:"getAudioByReciter"`
}

type manipulationStatements struct {
    insertNewChapter      string `json:"insertNewChapter"`
    updateChapterDetails  string `json:"updateChapterDetails"`
    deleteChapter         string `json:"deleteChapter"`
    insertNewVerse        string `json:"insertNewVerse"`
    updateVerseDetails    string `json:"updateVerseDetails"`
    deleteVerse           string `json:"deleteVerse"`
}

func init() {
    jsonStr := `{
		"queries": {
		  "listAllChapters": "SELECT * FROM chapters;",
		  "getChapterByNumber": "SELECT * FROM chapters WHERE chapterNumber = ?;",
		  "listVersesInChapter": "SELECT * FROM verses WHERE chapterNumber = ?;",
		  "getSpecificVerse": "SELECT * FROM verses WHERE verseNumber = ? AND chapterNumber = ?;",
		  "searchVersesByKeywords": "SELECT * FROM verses WHERE keywords LIKE '%?%';",
		  "listTranslationsForVerse": "SELECT * FROM translations WHERE verseNumber = ? AND chapterNumber = ?;",
		  "getTranslationByLanguage": "SELECT * FROM translations WHERE verseNumber = ? AND chapterNumber = ? AND language = ?;",
		  "listTafseerForVerse": "SELECT * FROM tafseer WHERE verseNumber = ? AND chapterNumber = ?;",
		  "getTafseerByScholar": "SELECT * FROM tafseer WHERE verseNumber = ? AND chapterNumber = ? AND scholarName = ?;",
		  "listAudioVersionsOfVerse": "SELECT * FROM audioVerses WHERE verseNumber = ? AND chapterNumber = ?;",
		  "getAudioByReciter": "SELECT * FROM audioVerses WHERE verseNumber = ? AND chapterNumber = ? AND reciterName = ?;"
		},
		"manipulations": {
		  "insertNewChapter": "INSERT INTO chapters (chapterNumber, totalVerses, englishName, arabicName, pageStart, pageEnd, totalPages, revelationPlace) VALUES (?, ?, ?, ?, ?, ?, ?, ?);",
		  "updateChapterDetails": "UPDATE chapters SET totalVerses = ?, englishName = ?, arabicName = ?, pageStart = ?, pageEnd = ?, totalPages = ?, revelationPlace = ? WHERE chapterNumber = ?;",
		  "deleteChapter": "DELETE FROM chapters WHERE chapterNumber = ?;",
		  "insertNewVerse": "INSERT INTO verses (verseNumber, chapterNumber, juzNumber, pageNumber, arabicText, englishText, keywords) VALUES (?, ?, ?, ?, ?, ?, ?);",
		  "updateVerseDetails": "UPDATE verses SET juzNumber = ?, pageNumber = ?, arabicText = ?, englishText = ?, keywords = ? WHERE verseNumber = ? AND chapterNumber = ?;",
		  "deleteVerse": "DELETE FROM verses WHERE verseNumber = ? AND chapterNumber = ?;"
		}
	  }
	  `

    err := json.Unmarshal([]byte(jsonStr), &SqlQueries)
    if err != nil {
        log.Fatalf("Error unmarshalling JSON: %v", err)
    }
}
