
# Memo

- Go言語ポインタ
  - C言語でおなじみのポインタですが、Go言語にもポインタがあります。 宣言する方法もC言語と同じで、たとえばint型変数のポインタは、「*int」と記述します。変数のアドレスを取得するときも、C言語と同様に「&」をつけ、アドレスから変数の中身へアクセスする時は 「`*`」を使います。

- 値渡し
  - ある変数を関数の引数として渡す場合、値のコピーが渡されます。そのため呼び出された関数内で変数の値を変更しても、元の値には影響がありません。これを「値渡し」と呼びます。  
- ポインタ渡し
  - ポインタ変数を関数に渡した場合は、ポインタが指し示す値が、呼び出し元も、呼び出された方も同じものであるため、呼び出された関数内で元の値を変更することができます。これを「ポインタ渡し（または参照渡し）」と呼びます。

(参考： https://goo.gl/QgMk8a)


- ポインタ配列
  - Go 言語の配列名は配列の先頭アドレスを表していません。また、配列にアドレス演算子を適用すると、配列へのポインタが生成されますが、それは配列の先頭アドレスを表しているわけではありません。Go 言語の配列はひとつの「値」なので、配列へのポインタは配列そのものを指し示すことになります。


### メソッドの `ポイントレシーバ` と `変数レシーバ`

```
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

```

- ポイントレシーバ
  - 呼び出し時に、変数、または、ポインタのいずれかのレシーバとして取ることができる

```
var v Vertex
ScaleFunc(v)  // Compile error!
ScaleFunc(&v) // OK

var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
```

- 変数レシーバ
  - メソッドが変数レシーバである場合、呼び出し時に、変数、または、ポインタのいずれかのレシーバとして取ることができる

```
var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // Compile error!

var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
```
