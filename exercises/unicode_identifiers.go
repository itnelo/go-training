package exercises

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Go supports Unicode identifiers (lol)
// Okay, if you want so... :D

//
// A sign enum
//

type Знак byte

const (
	Аард Знак = iota
	Аксий
	Гелиотроп
	Игни
	Ирден
	Квен
	Сомн
)

var знаки = []string{
	"Аард",
	"Аксий",
	"Гелиотроп",
	"Игни",
	"Ирден",
	"Квен",
	"Сомн",
}

// A stringer for a enumeration type
func (рунический Знак) String() string {
	if рунический > Сомн {
		panic("Такого знака нет. В енуме Сомн является последним знаком!")
	}

	return знаки[рунический]
}

//
// An object state
//

type Самочувствие struct {
	здоровье uint
}

//
// Actor behaviors
//

type Чувствующий interface {
	оценитьСамочувствие() (*Самочувствие, error)
}

//
// Actors on scene
//

type Ведьмак struct {
	имя        string
	рунический Знак
	ухудшаемое Самочувствие
}

// implementing "Чувствующий" interface
func (чувствующий *Ведьмак) оценитьСамочувствие() (ведьмачье *Самочувствие, нежданчик error) {
	if nil == чувствующий {
		нежданчик = fmt.Errorf("<nil> pointer")

		return
	}

	ведьмачье = &чувствующий.ухудшаемое

	return
}

// stringer (string serialization)
func (сериализующийся *Ведьмак) String() string {
	if nil == сериализующийся {
		return "<nil>"
	}

	return fmt.Sprintf(
		"Ведьмак %v (%v здоровья, выбран %v)",
		сериализующийся.имя,
		сериализующийся.ухудшаемое.здоровье,
		сериализующийся.рунический)
}

func (подготавливающийся *Ведьмак) выбратьЗнак(рунический Знак) {
	if nil == подготавливающийся {
		return
	}

	подготавливающийся.рунический = рунический

	fmt.Printf("%v выбрал знак %v\n", подготавливающийся, рунический)
}

// empty interfaces are used by code that handles values of unknown type
func (кастующий *Ведьмак) ухудшитьСамочувствие(цель interface{}) (нежданчик error) {
	if nil == кастующий {
		нежданчик = fmt.Errorf("<nil> pointer")

		return
	}

	if сейчасПолучит, можноВтащить := цель.(Чувствующий); можноВтащить {
		var самочувствие, _ = сейчасПолучит.оценитьСамочувствие()
		var генераторУрона = генераторУронаОтЗнаков()

		урон := генераторУрона()

		if самочувствие.здоровье >= урон {
			самочувствие.здоровье -= урон
		} else {
			самочувствие.здоровье = 0

			defer func() {
				fmt.Printf("%v погибает\n", цель)
			}()
		}

		fmt.Printf("%v получает %v урона\n", цель, урон)

		return
	}

	нежданчик = Оправдаться("Нельзя втащить этой цели, она же ничего не чувствует!", цель)

	return
}

func (убегающий *Ведьмак) убежать() (нежданчик error) {
	if nil == убегающий {
		нежданчик = fmt.Errorf("<nil> pointer")

		return
	}

	fmt.Printf("%v убежал, т.к. не смог втащить\n", убегающий)

	return
}

//
// A ghoul
//

type Гуль struct {
	ухудшаемое Самочувствие
}

// stringer
func (сериализующийся *Гуль) String() string {
	if nil == сериализующийся {
		return "<nil>"
	}

	return fmt.Sprintf("Гуль (%v здоровья)", сериализующийся.ухудшаемое.здоровье)
}

func (чувствующий *Гуль) оценитьСамочувствие() (чудовищное *Самочувствие) {
	if nil == чувствующий {
		return
	}

	чудовищное = &чувствующий.ухудшаемое

	return
}

//
// Helpers
//

func генераторУронаОтЗнаков() func() uint {
	var generatorSeed int64 = time.Now().UnixNano() / int64(time.Nanosecond)
	fmt.Println("Generator seed:", generatorSeed)

	rand.Seed(generatorSeed)

	return func() (урон uint) {
		урон = uint(rand.Intn(3000))

		return
	}
}

//
// Errors
//

type НеПолучилосьВтащить struct {
	оправдание string
	кому       interface{}
}

func (нежданчик *НеПолучилосьВтащить) Error() string {
	if nil == нежданчик {
		return "<nil>"
	}

	var stringsBuilder strings.Builder

	stringsBuilder.WriteString(нежданчик.оправдание)
	stringsBuilder.WriteString(fmt.Sprintf(" Цель: %v", нежданчик.кому))

	return stringsBuilder.String()
}

func Оправдаться(оправдание string, кому interface{}) *НеПолучилосьВтащить {
	return &НеПолучилосьВтащить{
		оправдание: оправдание,
		кому:       кому,
	}
}

//
// The scene
//

func светКамераМотор() {
	Ламберт := &Ведьмак{
		имя:        "Ламберт",
		рунический: Аард,
		ухудшаемое: Самочувствие{4500},
	}

	fmt.Printf("Ведьмак добавлен на сцену:\n%v\n", Ламберт)

	var недоброжелатели = []Гуль{
		Гуль{Самочувствие{1500}},
		Гуль{Самочувствие{700}},
		Гуль{Самочувствие{800}},
	}

	fmt.Printf("Гули добавлены на сцену:\n%+v\n", недоброжелатели)

	Ламберт.выбратьЗнак(Игни)

	for порядковыйНомер, _ := range недоброжелатели {
		var следующий = &недоброжелатели[порядковыйНомер]

		if нежданчик := Ламберт.ухудшитьСамочувствие(следующий); нежданчик != nil {
			fmt.Println("Произошла неожиданная ситуация.", нежданчик)
			Ламберт.убежать()

			break
		}
	}

	// Ведьмак добавлен на сцену:
	// Ведьмак Ламберт (4500 здоровья, выбран Аард)
	// Гули добавлены на сцену:
	// [{ухудшаемое:{здоровье:1500}} {ухудшаемое:{здоровье:700}} {ухудшаемое:{здоровье:800}}]
	// Ведьмак Ламберт (4500 здоровья, выбран Игни) выбрал знак Игни
	// Generator seed: 1555520466156621595
	// Гуль (1006 здоровья) получает 494 урона
	// Generator seed: 1555520466156634636
	// Гуль (0 здоровья) получает 1791 урона
	// Гуль (0 здоровья) погибает
	// Generator seed: 1555520466156647227
	// Гуль (0 здоровья) получает 2250 урона
	// Гуль (0 здоровья) погибает
}
