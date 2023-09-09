# api-matriks

api-matriks

# API OPERASI MATRIKS

## Deskripsi

API ini menyediakan layanan untuk melakukan operasi antara dua matriks. Operasi-operasi ini dapat digunakan untuk berbagai keperluan perhitungan matriks.

## Penggunaan

### Pertambahan Matriks

Untuk melakukan pertambahan antara dua matriks, Anda dapat menggunakan endpoint berikut:

### endpoint

POST http://localhost:3000/api/add

```json
{
	"matrixA": [
		[1, 2],
		[3, 4]
	],
	"matrixB": [
		[5, 6],
		[7, 8]
	]
}
```

### response

```json
{
	"result": [
		[6, 8],
		[10, 12]
	]
}
```

### Pengurangan Matriks

Untuk melakukan pengurangan antara dua matriks, Anda dapat menggunakan endpoint berikut:

### endpoint

POST http://localhost:3000/api/subtract

```json
{
	"matrixA": [
		[1, 2],
		[3, 4]
	],
	"matrixB": [
		[5, 6],
		[7, 8]
	]
}
```

### response

```json
{
	"result": [
		[-4, -4],
		[-4, -4]
	]
}
```

### Perkalian Matriks

Untuk melakukan perkalian antara dua matriks, Anda dapat menggunakan endpoint berikut:

### endpoint

POST http://localhost:3000/api/multiply

```json
{
	"matrixA": [
		[1, 2],
		[3, 4]
	],
	"matrixB": [
		[5, 6],
		[7, 8]
	]
}
```

### response

```json
{
	"result": [
		[19, 22],
		[43, 50]
	]
}
```

### Transpose Matriks

Untuk melakukan transpose matriks, Anda dapat menggunakan endpoint berikut:

### endpoint

POST http://localhost:3000/api/transpose

```json
{
	"matrix": [
		[1, 2],
		[3, 4]
	]
}
```

### response

```json
{
	"result": [
		[1, 3],
		[2, 4]
	]
}
```

### Inverse Matriks

Untuk melakukan invers matriks, Anda dapat menggunakan endpoint berikut:

### endpoint

POST http://localhost:3000/api/inverse

```json
{
	"matrix": [
		[1, 2],
		[3, 4]
	]
}
```

### response

```json
{
	"result": [
		[-2, 1],
		[1.5, -0.5]
	]
}
```
