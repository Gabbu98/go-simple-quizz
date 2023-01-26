package cmd

type question struct {
	Id       int      `json:"id"`
	Question string   `json:"question"`
	Choices  []Choice `json:"choices"`
}

type option struct {
	QuestionId int `json:"questionId"`
}

type Choice struct {
	Id     int    `json:"id"`
	Choice string `json:"choice"`
}

type correctAnswer struct {
	QuestionId int `json:"questionId"`
	ChoiceId   int `json:"choiceId"`
}

type answer struct {
	QuestionId int `json:"questionId"`
	ChoiceId   int `json:"choiceId"`
}

var database = []question{
	{Id: 1,
		Question: "With respect to the certification of airmen, which is a category of aircraft?", Choices: []Choice{
			{Id: 1, Choice: "A. Gyroplane, helicopter, airship, free balloon."},
			{Id: 2, Choice: "B. Airplane, rotorcraft, glider, lighter-than-air."},
			{Id: 3, Choice: "C. Single-engine land and sea, multiengine land and sea."},
		},
	
	},
	{Id: 2,
		Question: "With respect to the certification of airmen, which is a category of aircraft?", Choices: []Choice{
			{Id: 1, Choice: "A. Gyroplane, helicopter, airship, free balloon."},
			{Id: 2, Choice: "B. Airplane, rotorcraft, glider, lighter-than-air."},
			{Id: 3, Choice: "C. Single-engine land and sea, multiengine land and sea."},
		},
	
	},
	{Id: 3,
		Question: "With respect to the certification of airmen, which is a category of aircraft?", Choices: []Choice{
			{Id: 1, Choice: "A. Gyroplane, helicopter, airship, free balloon."},
			{Id: 2, Choice: "B. Airplane, rotorcraft, glider, lighter-than-air."},
			{Id: 3, Choice: "C. Single-engine land and sea, multiengine land and sea."},
		},
	
	},
	{Id: 4,
		Question: "With respect to the certification of airmen, which is a category of aircraft?", Choices: []Choice{
			{Id: 1, Choice: "A. Gyroplane, helicopter, airship, free balloon."},
			{Id: 2, Choice: "B. Airplane, rotorcraft, glider, lighter-than-air."},
			{Id: 3, Choice: "C. Single-engine land and sea, multiengine land and sea."},
		},
	
	},
	{Id: 5,
		Question: "With respect to the certification of airmen, which is a category of aircraft?", Choices: []Choice{
			{Id: 1, Choice: "A. Gyroplane, helicopter, airship, free balloon."},
			{Id: 2, Choice: "B. Airplane, rotorcraft, glider, lighter-than-air."},
			{Id: 3, Choice: "C. Single-engine land and sea, multiengine land and sea."},
		},
	
	},
	{Id: 6,
		Question: "With respect to the certification of airmen, which is a category of aircraft?", Choices: []Choice{
			{Id: 1, Choice: "A. Gyroplane, helicopter, airship, free balloon."},
			{Id: 2, Choice: "B. Airplane, rotorcraft, glider, lighter-than-air."},
			{Id: 3, Choice: "C. Single-engine land and sea, multiengine land and sea."},
		},
	
	},
	{Id: 7,
		Question: "With respect to the certification of airmen, which is a category of aircraft?", Choices: []Choice{
			{Id: 1, Choice: "A. Gyroplane, helicopter, airship, free balloon."},
			{Id: 2, Choice: "B. Airplane, rotorcraft, glider, lighter-than-air."},
			{Id: 3, Choice: "C. Single-engine land and sea, multiengine land and sea."},
		},
	
	},
	{Id: 8,
		Question: "With respect to the certification of airmen, which is a category of aircraft?", Choices: []Choice{
			{Id: 1, Choice: "A. Gyroplane, helicopter, airship, free balloon."},
			{Id: 2, Choice: "B. Airplane, rotorcraft, glider, lighter-than-air."},
			{Id: 3, Choice: "C. Single-engine land and sea, multiengine land and sea."},
		},
	
	},
	{Id: 9,
		Question: "With respect to the certification of airmen, which is a category of aircraft?", Choices: []Choice{
			{Id: 1, Choice: "A. Gyroplane, helicopter, airship, free balloon."},
			{Id: 2, Choice: "B. Airplane, rotorcraft, glider, lighter-than-air."},
			{Id: 3, Choice: "C. Single-engine land and sea, multiengine land and sea."},
		},
	
	},
	{Id: 10,
		Question: "With respect to the certification of airmen, which is a category of aircraft?", Choices: []Choice{
			{Id: 1, Choice: "A. Gyroplane, helicopter, airship, free balloon."},
			{Id: 2, Choice: "B. Airplane, rotorcraft, glider, lighter-than-air."},
			{Id: 3, Choice: "C. Single-engine land and sea, multiengine land and sea."},
		},
	
	},
}

var options = []option{
	{QuestionId: 1}}

var correctAnswers = []answer{
	{QuestionId: 1, ChoiceId: 2},
}

var answers = []answer{}