package main

func TestSoma(t *testing.T){

	total := Soma(15, 15)

	fi total != 30 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado %d", total, 30)
	}

}