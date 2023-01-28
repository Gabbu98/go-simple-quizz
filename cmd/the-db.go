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
	ChoiceId   string `json:"choiceId"`
}

type playerStruct struct {
	Name string `json:"name"`
	Score   int `json:"score"`
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
		Question: "The definition of nighttime is?", Choices: []Choice{
			{Id: 1, Choice: "A. the time between the end of evening civil twilight and the beginning of morning civil twilight."},
			{Id: 2, Choice: "B. sunset to sunrise."},
			{Id: 3, Choice: "C. 1 hour after sunset to 1 hour before sunrise."},
		},
	
	},
	{Id: 3,
		Question: "Which V-speed represents maneuvering speed?", Choices: []Choice{
			{Id: 1, Choice: "A. VLO"},
			{Id: 2, Choice: "B. VNE"},
			{Id: 3, Choice: "C. VA"},
		},
	
	},
	{Id: 4,
		Question: "Which V-speed represents maximum flap extended speed?", Choices: []Choice{
			{Id: 1, Choice: "A.VFE"},
			{Id: 2, Choice: "B. VLOF"},
			{Id: 3, Choice: "C. VFC"},
		},
	
	},
	{Id: 5,
		Question: "VNO is defined as the", Choices: []Choice{
			{Id: 1, Choice: "A. normal operating range."},
			{Id: 2, Choice: "B. maximum structural cruising speed."},
			{Id: 3, Choice: "C. never-exceed speed."},
		},
	
	},
	{Id: 6,
		Question: "When must a current pilot certificate be in the pilot's personal possession or readily accessible in the aircraft?", Choices: []Choice{
			{Id: 1, Choice: "A. Only when passengers are carried."},
			{Id: 2, Choice: "B. Anytime when acting as pilot in command or as a required crewmember."},
			{Id: 3, Choice: "C. When acting as a crew chief during launch and recovery."},
		},
	
	},
	{Id: 7,
		Question: "With respect to daylight hours, what is the earliest time a recreational pilot may take off", Choices: []Choice{
			{Id: 1, Choice: "A. At the beginning of morning civil twilight."},
			{Id: 2, Choice: "B. One hour before sunrise."},
			{Id: 3, Choice: "C. At sunrise."},
		},
	
	},
	{Id: 8,
		Question: "What does PIC stand for?", Choices: []Choice{
			{Id: 1, Choice: "A. Person in Command"},
			{Id: 2, Choice: "B. Penguin in Command"},
			{Id: 3, Choice: "C. Pilot in Command"},
		},
	
	},
	{Id: 9,
		Question: "Which cloud is associated with thunderstorms?", Choices: []Choice{
			{Id: 1, Choice: "A. Cumulus"},
			{Id: 2, Choice: "B. Stratus"},
			{Id: 3, Choice: "C. Cumulonimbus"},
		},
	
	},
	{Id: 10,
		Question: "What does NDB mean?", Choices: []Choice{
			{Id: 1, Choice: "A. Non-Directional Beacon"},
			{Id: 2, Choice: "B. Non-Distant Bearing"},
			{Id: 3, Choice: "C. Near Distant Beacon"},
		},
	
	},
}

var correctAnswers = []answer{
	{QuestionId: 1, ChoiceId: "2"},
	{QuestionId: 2, ChoiceId: "1"},
	{QuestionId: 3, ChoiceId: "3"},
	{QuestionId: 4, ChoiceId: "1"},
	{QuestionId: 5, ChoiceId: "1"},
	{QuestionId: 6, ChoiceId: "2"},
	{QuestionId: 7, ChoiceId: "3"},
	{QuestionId: 8, ChoiceId: "3"},
	{QuestionId: 9, ChoiceId: "3"},
	{QuestionId: 10, ChoiceId: "1"},
}

var standings = []playerStruct{
	{Name: "Geoff", Score: 60},
	{Name: "Julian", Score: 70},
	{Name: "Tina", Score: 80},
	{Name: "Mary", Score: 10},
	{Name: "Simon", Score: 40},
}