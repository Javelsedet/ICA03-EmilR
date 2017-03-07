package ascii

import "fmt"

const ascii = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f" +
	"\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f" +
	` !"#$%&'()*+,-./0123456789:;<=>?` +
	`@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_` +
	"`abcdefghijklmnopqrstuvwxyz{|}~\x7f"

const hello = "\x48\x65\x6c\x6c\x6f\x20\x3a\x2d\x29"

// Funksjon tar en streng med kun ASCII tegn og lager en utskrift på
// følgende format:
// [ascii-kode heksadesimalt med store bokstaver A-F][mellomrom]
// [symbol for ascii-kode][mellomrom][ascii-kode binært][linjeskift]
// Eksempel (bruk denne funksjonen i en main.go fil):
//	…
// 3E > 111110
// 3F ? 111111
// 40 @ 1000000
// ...
func IterateOverASCIIStringLiteral(sl string) {
	// Kode for Oppgave 1a

}

// Funksjonen skal generere en utskrift fra en sekvens av bytes,
// dvs. av typen []bytes (det betyr at du må finne den heksadesimale
// eller binære representasjonen av alle tegn i strengen,
// som skal skrives ut (inkludert anførselstegn eller
// “double quotes” på engelsk).
// Funksjonen greetingASCII() returnerer en variabel av typen string,
// som inneholder kun ASCII tegn (ikke utvidet ASCII).
// Gjelder oppgave 1b
func GreetingASCII() string {
	t1 := []byte("\x48\x65\x6c\x6c\x6f\x20\x3a\x2d\x29")
	for i := 0; i < len(t1); i++ {
		fmt.Printf("%c", t1[i])
	}
	t2 := "\x48\x65\x6c\x6c\x6f\x20\x3a\x2d\x29"
	return t2

}
