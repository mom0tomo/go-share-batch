package main

import (
	"fmt"
	"os"

	"gopkg.in/gographics/imagick.v2/imagick"
)

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	// base
	mw1 := imagick.NewMagickWand()
	defer mw1.Destroy()
	base, err := mw1.ReadImage(os.Args[1])
	if err != nil {
		return fmt.Errorf("画像ファイルが読み込めませんでした%s", base)
	}

	// Book image
	mw2 := imagick.NewMagickWand()
	defer mw2.Destroy()
	book, err = mw2.ReadImage(os.Args[2])
	if err != nil {
		return fmt.Errorf("画像ファイルが読み込めませんでした%s", base)
	}

	// Music image
	mw3 := imagick.NewMagickWand()
	defer mw3.Destroy()
	music, err = mw3.ReadImage(os.Args[3])
	if err != nil {
		return fmt.Errorf("画像ファイルが読み込めませんでした%s", music)
	}

	// Movie image
	mw4 := imagick.NewMagickWand()
	defer mw4.Destroy()
	movie, err = mw4.ReadImage(os.Args[4])
	if err != nil {
		return fmt.Errorf("画像ファイルが読み込めませんでした%s", movie)
	}

	// Profile image
	mw5 := imagick.NewMagickWand()
	defer mw5.Destroy()
	prof, err = mw5.ReadImage(os.Args[5])
	if err != nil {
		return fmt.Errorf("画像ファイルが読み込めませんでした%s", prof)
	}

	// TODO: 縦横比をいい感じにする
	// リサイズする
	rbook, err = mw2.ResizeImage(100, 100, imagick.FILTER_POINT, 0)
	if err != nil {
		fmt.Errorf("画像をリサイズできませんでした%s", rbook)
	}

	rmovie, err = mw3.ResizeImage(100, 100, imagick.FILTER_POINT, 0)
	if err != nil {
		fmt.Errorf("画像をリサイズできませんでした%s", rmovie)
	}

	rmusic, err = mw4.ResizeImage(100, 100, imagick.FILTER_POINT, 0)
	if err != nil {
		fmt.Errorf("画像をリサイズできませんでした%s", rmusic)
	}

	rprof, err = mw5.ResizeImage(100, 100, imagick.FILTER_POINT, 0)
	if err != nil {
		fmt.Errorf("画像をリサイズできませんでした%s", rprof)
	}
	//w1 := mw1.GetImageWidth()
	//h1 := mw1.GetImageHeight()

	//w2 := mw2.GetImageWidth()
	//h2 := mw2.GetImageHeight()

	// TODO: 切り抜き変更する

	// 画像を重ねる
	// ベースの画像の上に2枚目を乗せる
	pbook, err = mw1.CompositeImage(mw2, mw2.GetImageCompose(), 0, 0)
	if err != nil {
		fmt.Errorf("画像を合成できませんでした%s", pbook)
	}

	// ベースの画像の上に3枚目を乗せる
	pmovie, err = mw1.CompositeImage(mw3, mw3.GetImageCompose(), 0, 100)
	if err != nil {
		fmt.Errorf("画像を合成できませんでした%s", pmovie)
	}

	// ベースの画像の上に4枚目を乗せる
	pmusic, err = mw1.CompositeImage(mw4, mw4.GetImageCompose(), 100, 0)
	if err != nil {
		fmt.Errorf("画像を合成できませんでした%s", pmusic)
	}

	// 1枚目の画像の上に5枚目を乗せる
	pprof, err = mw1.CompositeImage(mw5, mw5.GetImageCompose(), 100, 100)
	if err != nil {
		fmt.Errorf("画像を合成できませんでした%s", pprof)
	}

	// 合成したイメージをファイルにする
	mw1.WriteImage("out.jpg")

	fmt.Fprintln(os.Stderr, "FINISH!!")
}
