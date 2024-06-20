package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Determina el directorio desde donde se llama el programa
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error obteniendo el directorio actual:", err)
		return
	}

	fmt.Println("Directorio actual:", dir)

	// Nombre del archivo de salida
	outputFilename := filepath.Join(dir, "mobspecc.go")

	// Abre el archivo CSV de entrada
	inputFilename := filepath.Join("prebuild", "mobspec.csv")
	fmt.Println("Archivo de entrada:", inputFilename)
	file, err := os.Open(inputFilename)
	if err != nil {
		fmt.Println("Error abriendo el archivo CSV:", err)
		return
	}
	defer file.Close()

	// Lee el archivo CSV
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error leyendo el archivo CSV:", err)
		return
	}

	// Crea el archivo de salida
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		fmt.Println("Error creando el archivo de salida:", err)
		return
	}
	defer outputFile.Close()

	// Escribe en el archivo de salida
	fmt.Fprintln(outputFile, "package mobspec") // Ajusta el paquete según sea necesario
	fmt.Fprintln(outputFile)
	fmt.Fprintln(outputFile, "var MobileSpecsData = []MobileSpecs{")

	for _, record := range records {
		fmt.Fprintf(outputFile, "\t{\n")
		fmt.Fprintf(outputFile, "\t\tBuildID:                 \"%s\",\n", record[0])
		fmt.Fprintf(outputFile, "\t\tBuildDisplayID:          \"%s\",\n", record[1])
		fmt.Fprintf(outputFile, "\t\tProductName:             \"%s\",\n", record[2])
		fmt.Fprintf(outputFile, "\t\tProductDevice:           \"%s\",\n", record[3])
		fmt.Fprintf(outputFile, "\t\tProductBoard:            \"%s\",\n", record[4])
		fmt.Fprintf(outputFile, "\t\tProductManufacturer:     \"%s\",\n", record[5])
		fmt.Fprintf(outputFile, "\t\tProductBrand:            \"%s\",\n", record[6])
		fmt.Fprintf(outputFile, "\t\tProductModel:            \"%s\",\n", record[7])
		fmt.Fprintf(outputFile, "\t\tBootloader:              \"%s\",\n", record[8])
		fmt.Fprintf(outputFile, "\t\tHardware:                \"%s\",\n", record[9])
		fmt.Fprintf(outputFile, "\t\tBuildType:               \"%s\",\n", record[10])
		fmt.Fprintf(outputFile, "\t\tBuildTags:               \"%s\",\n", record[11])
		fmt.Fprintf(outputFile, "\t\tBuildFingerprint:        \"%s\",\n", record[12])
		fmt.Fprintf(outputFile, "\t\tBuildUser:               \"%s\",\n", record[13])
		fmt.Fprintf(outputFile, "\t\tBuildHost:               \"%s\",\n", record[14])
		fmt.Fprintf(outputFile, "\t\tBuildVersionIncremental: \"%s\",\n", record[15])
		fmt.Fprintf(outputFile, "\t\tBuildVersionRelease:     \"%s\",\n", record[16])
		fmt.Fprintf(outputFile, "\t\tBuildVersionSDK:         \"%s\",\n", record[17])
		fmt.Fprintf(outputFile, "\t\tBuildVersionCodename:    \"%s\",\n", record[18])
		fmt.Fprintf(outputFile, "\t\tScreenHeight:            \"%s\",\n", record[19])
		fmt.Fprintf(outputFile, "\t\tScreenWidth:             \"%s\",\n", record[20])
		fmt.Fprintf(outputFile, "\t\tScreenDensity:           \"%s\",\n", record[21])
		fmt.Fprintf(outputFile, "\t\tScreenXDPI:              \"%s\",\n", record[22])
		fmt.Fprintf(outputFile, "\t\tScreenYDPI:              \"%s\",\n", record[23])
		fmt.Fprintf(outputFile, "\t\tScreenDPI:               \"%s\",\n", record[24])
		fmt.Fprintf(outputFile, "\t\tScreenScaledDensity:     \"%s\",\n", record[25])
		fmt.Fprintf(outputFile, "\t},\n")
	}

	fmt.Fprintln(outputFile, "}")

	fmt.Println("Archivo generado:", outputFilename)
}
