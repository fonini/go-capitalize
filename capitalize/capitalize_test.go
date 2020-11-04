package capitalize

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// test that input matches the value we want. If not, report an error on t.
func testValue(t *testing.T, input, want string) {
	v, err := Capitalize(input)

	if err != nil {
		t.Errorf("Capitalize(%q) returned error: %v", input, err)
	}

	if diff := cmp.Diff(want, v); diff != "" {
		t.Errorf("Capitalize(%q) mismatch:\n%s", input, diff)
	}
}

// test that input matches the value we want. If not, report an error on t.
func testValueWithOptions(t *testing.T, input, want string, options Options) {
	v, err := Capitalize(input, options)

	if err != nil {
		t.Errorf("Capitalize(%q) returned error: %v", input, err)
	}

	if diff := cmp.Diff(want, v); diff != "" {
		t.Errorf("Capitalize(%q) mismatch:\n%s", input, diff)
	}
}

func TestValues_BrazilianNames(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"jonnas fonini", "Jonnas Fonini"},
		{"joão DA Silva", "João da Silva"},
		{"dom JOÃO vi", "Dom João VI"},
		{"maria helena dos santos", "Maria Helena dos Santos"},
		{"SANT'ANA DO LIVRAMENTO", "Sant'Ana do Livramento"},
		{"pedro dos passos", "Pedro dos Passos"},
		{"sebastião DE bourBOn E Bragança", "Sebastião de Bourbon e Bragança"},
		{"maria   da     conceição", "Maria da Conceição"},
	}

	for _, tt := range tests {
		testValue(t, tt.input, tt.want)
	}
}

func TestValues_SpanishNames(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"miguel del castillo", "Miguel del Castillo"},
		{"DOM QUIXOTE DE LA MANCHA", "Dom Quixote de la Mancha"},
		{"gimnasia y  ESGrima", "Gimnasia y Esgrima"},
	}

	for _, tt := range tests {
		testValue(t, tt.input, tt.want)
	}
}

func TestValues_ItalianNames(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"martino iii della torre", "Martino III della Torre"},
		{"sant'ilARio dello IONIO", "Sant'Ilario dello Ionio"},
	}

	for _, tt := range tests {
		testValue(t, tt.input, tt.want)
	}
}

func TestValues_DutchNames(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"EDDIE VAN HALEN", "Eddie van Halen"},
	}

	for _, tt := range tests {
		testValue(t, tt.input, tt.want)
	}
}

func TestValues_GermanNames(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"wernher von braun", "Wernher von Braun"},
	}

	for _, tt := range tests {
		testValue(t, tt.input, tt.want)
	}
}

func TestValues_WithOptions(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"gørvel fadersdotter of giske", "Gørvel Fadersdotter of Giske"},
		{"BREDO VON MUNTHE AF MORGENSTIERNE", "Bredo von Munthe af Morgenstierne"},
	}

	options := Options{
		Exceptions: []string{"of", "af"},
	}

	for _, tt := range tests {
		testValueWithOptions(t, tt.input, tt.want, options)
	}
}

func ExampleCapitalize() {
	name, _ := Capitalize("jonnas fonini")
	fmt.Print(name)

	// Output: Jonnas Fonini
}
