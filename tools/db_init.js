// A file to initialize mongodb database.
// Creates database and collections, and adds initial documents,
// also creates indexes.

// run `mongo <database> db_init.js`

try {
	// Check if database kails exists
	// drop it and create a new clean
	// database.
	db._dropDatabase("kails");
	db._createDatabase("kails");
	db._useDatabase("kails");

} catch(e) {
	if (e == "[ArangoError 1228: database not found]") {
		db._createDatabase("kails");
		db._useDatabase("kails");
	} else {
		throw "Some unknown error" + e + "happened";
	}
}

db._create("users");
db._create("languages");
db._create("topics");
db._create("privileges");
// User = user
// Password = password
default_user = {
	"Username" : "user",
	"Email" : "user@email.com",
	"Password" : "RurG60nC/Kx9N0MumOq74K7tNwebAjWC9AYXhJOqTFY=",
	"Salt" : "Li6QELxiH4vcqg==",
	"Language" : "es-MX",
	"StudyLanguage": "english",
	"Since" : new Date(),
	"Levels" : {
		"1": {
			"Unlocked": true,
			// last review
			"Last":  new Date(),
		}
	},
	"Topics" : [
	{
		"Topic": "sports",
		"Subtopics": ["soccer", "baseball"]
	},
	{
		"Topic": "entertainment",
		"Subtopics": ["videogames", "movies", "theater"]
	},
	{
		"Topic": "vehicles",
		"Subtopics": ["motorcycles"]
	}
	],
};
db.users.save(default_user);

english_program = {
	"Type": "program",
	"Language": "english",
	"Levels": [
		{
			"Id": 1,
			"Description": "some sort of description",
		},{
			"Id": 2,
			"Description": "some other sort of description",
		},
	]
};
db.languages.save(english_program);

// English language collection
// contains levels and words
level_1 = {
	"Id": 1,
	"Type": "level",
	"Language": "english",
	// description
	"Description": "some sort of description",
	"Version": 1,
	"Words": [
		{
			"Word": "word1",
			"Translation": "blah blah blah",
			"Class": "verb",
			// rename this to challenge?
			// fuck it, so much to change T.T
			"Sentences": [
				{
					// .......Dropping the thought here.........
					// perhaps it's better "question/answer"
					// instead of native/translation
					// .........................................
					"Native": "english sentence here",
					"Translation": "translation here"
				},{
					"Native": "other english sentence here",
					"Translation": "other translation here"
				},
			]
		},
	]
};

// level_2
level_2 = {
	"Id": 2,
	"Type": "level",
	"Language": "english",
	// description
	"Desc": "some sort of description",
	"Version": 1,
	"Words": [
		{
			"Word": "word1",
			"Translation": "blah blah blah",
			"Class": "verb",
			// rename this to challenge?
			// fuck it, so much to change T.T
			"Sentences": [
				{
					// .......Dropping the thought here.........
					// perhaps it's better "question/answer"
					// instead of native/translation
					// .........................................
					"Native": "blah blah",
					"Translation": "blah blah"
				},{
					"Native": "other english sentence here",
					"Translation": "other translation here"
				},
			]
		},
	]
};
db.languages.save(level_1);
db.languages.save(level_2);

word1 = {
	"Word": "word1",
	"Type": "word",
	"Level": 1,
	"Class": "verb",
	"Language": "english",
	"Translation": "blah blah blah blah",
	"Sentences": [
		{
			"Native": "english sentence here",
			"Translation": "translation here"
		},{
			"Native": "other english sentence here",
			"Translation": "other translation here"
		},
	],

};
db.languages.save(word1);

topic = {
	"Id": 1,
	"Name": "sports",
	"Subtopics": [
		{
			"name": "soccer",
			"noUsers": 0
		},
		{
			"name": "football",
			"noUsers": 0
		}
	]
}
db.topics.save(topic);

privilege = {
	"Level": 1,
	"Yopics": 1,
	"Friends": 5,
	"Features": ["chat"],
	"Time": 10 //Time in minutes
}
db.privileges.save(privilige);

// Add indexes.
// user collection indexes
db.users.ensureFulltextIndex("Username");
db.users.ensureUniqueConstraint("Name");
db.users.ensureUniqueConstraint("Email");

// languages collection indexes
db.languages.ensureHashIndex("Language");
db.languages.ensureHashIndex("Level");
db.languages.ensureHashIndex("Type");
db.languages.ensureHashIndex("Id");
db.languages.ensureHashIndex("Word");

print("done!");
