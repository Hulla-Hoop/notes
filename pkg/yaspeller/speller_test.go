package yaspeller_test

import (
	"testing"

	"github.com/jMurad/notes/pkg/yaspeller"
)

func Test_RightText(t *testing.T) {
	tests := []struct {
		name      string
		wrongText string
		rightText string
	}{
		{"1", "Бабушка выстовила зимнии рамы. И сразу в избе стало светло весела. За окном в кустах вороби чирикают ласточки щибечут ребетишки кричат и смеюца.",
			"Бабушка выставила зимние рамы. И сразу в избе стало светло весело. За окном в кустах воробьи чирикают ласточки щебечут ребятишки кричат и смеются."},
		{"2", "Дедушка в агороди капал землю. Он аперся на заступ и по сматрел на небо. Ана небе уже сново сеяла сонце. И только две маленькии лиловые тучьки летели по ветру и словна до ганяли друг друга. Дедушка разеснил Тане что это не гроза, проста сталкнулись молодые тучки ударились молнию высикли.",
			"Дедушка в огороде копал землю. Он оперся на заступ и посмотрел на небо. А на небе уже снова сияло солнце. И только две маленькие лиловые тучки летели по ветру и словно догоняли друг друга. Дедушка разъяснил Тане что это не гроза, просто столкнулись молодые тучки ударились молнию высекли."},
		{"3", "Сутра марасит дождь. Дорога от школы до деревьне идет около опушке леса. Со всех сторон вода капает тичет льеца. Зашол я под деревья. Трехнула березка ветками и струйка мне за варатник папала.",
			"С утра моросит дождь. Дорога от школы до деревне идет около опушке леса. Со всех сторон вода капает течет льется. Зашел я под деревья. Тряхнула березка ветками и струйка мне за воротник попала."},
		{"4", "Вот стоит старая ель. Нижнее лапы унее шатром весят. Под этим шатром сухо струиться там теплый воздух. Залес я подплотные лапы сумку слажил сежу и наблюдаю.",
			"Вот стоит старая ель. Нижнее лапы у нее шатром висят. Под этим шатром сухо струится там теплый воздух. Залез я под плотные лапы сумку сложил сижу и наблюдаю."},
		{"5", "Летам мы чясто ходим в лес за грибами и ягадами. В лису типло и светло. Дует лёхкий ветирок. Громко пают птицы. Вот зашуршала под кустом лисная мыш. Там сычь затинул сваю песьню. Громко застучал па бальному дериву Дятел. Он-настоящий лесной доктор. Уж тихо зашуршал в сухой траве. Лес наполнен удевительными звуками прероды.",
			"Летом мы часто ходим в лес за грибами и ягодами. В лесу тепло и светло. Дует легкий ветерок. Громко поют птицы. Вот зашуршала под кустом лесная мышь. Там сыч затянул свою песню. Громко застучал по больному дереву Дятел. Он-настоящий лесной доктор. Уж тихо зашуршал в сухой траве. Лес наполнен удивительными звуками природы."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := yaspeller.CheckText(tt.wrongText)
			if err != nil {
				t.Errorf("Test [%s] is failed.\nError: %v\n", tt.name, err)
			}
			right := res.RightText(tt.wrongText)
			if right != tt.rightText {
				t.Errorf("Test [%s] is failed.\nWrong text: {%s}\nResult text: {%s}\n", tt.name, tt.wrongText, right)
			}
		})
	}
}

func Test_IsCorrect(t *testing.T) {
	testCases := []struct {
		name    string
		text    string
		isValid bool
	}{
		{
			name:    "1",
			text:    "Бабушка выстовила зимнии рамы. И сразу в избе стало светло весела. За окном в кустах вороби чирикают ласточки щибечут ребетишки кричат и смеюца.",
			isValid: false,
		},
		{
			name:    "2",
			text:    "Дедушка в огороде копал землю. Он оперся на заступ и посмотрел на небо. А на небе уже снова сияло солнце. И только две маленькие лиловые тучки летели по ветру и словно догоняли друг друга. Дедушка разъяснил Тане что это не гроза, просто столкнулись молодые тучки ударились молнию высекли.",
			isValid: true,
		},
		{
			name:    "3",
			text:    "Сутра марасит дождь. Дорога от школы до деревьне идет около опушке леса. Со всех сторон вода капает тичет льеца. Зашол я под деревья. Трехнула березка ветками и струйка мне за варатник папала.",
			isValid: false,
		},
		{
			name:    "4",
			text:    "Вот стоит старая ель. Нижнее лапы у нее шатром висят. Под этим шатром сухо струится там теплый воздух. Залез я под плотные лапы сумку сложил сижу и наблюдаю.",
			isValid: true,
		},
		{
			name:    "5",
			text:    "Летам мы чясто ходим в лес за грибами и ягадами. В лису типло и светло. Дует лёхкий ветирок. Громко пают птицы. Вот зашуршала под кустом лисная мыш. Там сычь затинул сваю песьню. Громко застучал па бальному дериву Дятел. Он-настоящий лесной доктор. Уж тихо зашуршал в сухой траве. Лес наполнен удевительными звуками прероды.",
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := yaspeller.CheckText(tc.text)
			if err != nil {
				t.Errorf("Test [%s] is failed.\nError: %v\n", tc.name, err)
			}
			right := res.IsCorrect()
			if right != tc.isValid {
				t.Errorf("Test [%s] is failed. The 'right' the contains errors\n", tc.name)
			}
		})
	}
}