package game

import (
	"image/color"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	slime "gitlab.utc.fr/vilecler/musicslime/slime"
	types "gitlab.utc.fr/vilecler/musicslime/types"
)

type Game struct {
	agents   [types.AgentsNum]slime.Agent
	trailMap types.TrailMap
}

func NewGame() *Game {
	g := &Game{}

	//Init TrailMap
	g.trailMap = make([][][]float64, int(types.GetWindowDefault().Width))

	for i := 0; i < len(g.trailMap); i++ {
		g.trailMap[i] = make([][]float64, int(types.GetWindowDefault().Height))
		for j := 0; j < len(g.trailMap[i]); j++ {
			g.trailMap[i][j] = make([]float64, 4)
		}
	}

	//Starting simulation
	for i := 0; i < types.AgentsNum; i++ {
		g.agents[i].Start(i)
	}
	return g
}

func (g *Game) Update() error {
	var wg sync.WaitGroup
	wg.Add(types.AgentsNum)

	g.trailMap = slime.Diffuse(g.trailMap)

	for i := 0; i < types.AgentsNum/1000; i++ {
		go func(index int, delim int) {
			for j := index; j < (index + delim); j++ {
				g.agents[j].Update(&g.trailMap)
				wg.Done()
			}
		}(i, 1000)
	}

	wg.Wait()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := 0; i < types.AgentsNum; i++ {
		g.agents[i].Draw(screen)
	}


	firstColor := slime.GetSpeciesSettingsByID(0).Colour
	secondColor := slime.GetSpeciesSettingsByID(1).Colour
	thirdColor := slime.GetSpeciesSettingsByID(2).Colour
	fourthColor := slime.GetSpeciesSettingsByID(3).Colour

	for i := 0; i < len(g.trailMap); i++ {
		for j := 0; j < len(g.trailMap[i]); j++ {
			if g.trailMap[i][j][0] > 0.0{
				var color color.RGBA
				color.R = uint8(firstColor[0] * 255)
				color.G = uint8(firstColor[1] * 255)
				color.B = uint8(firstColor[2] * 255)
				color.A = uint8(g.trailMap[i][j][0] * 255)
				ebitenutil.DrawRect(screen, float64(i), float64(j), 1, 1, color)
			} else if(g.trailMap[i][j][1] > 0.0){
				var color color.RGBA
				color.R = uint8(secondColor[0] * 255)
				color.G = uint8(secondColor[1] * 255)
				color.B = uint8(secondColor[2] * 255)
				color.A = uint8(g.trailMap[i][j][1] * 255)
				ebitenutil.DrawRect(screen, float64(i), float64(j), 1, 1, color)
			} else if(g.trailMap[i][j][2] > 0.0){
				var color color.RGBA
				color.R = uint8(thirdColor[0] * 255)
				color.G = uint8(thirdColor[1] * 255)
				color.B = uint8(thirdColor[2] * 255)
				color.A = uint8(g.trailMap[i][j][2] * 255)
				ebitenutil.DrawRect(screen, float64(i), float64(j), 1, 1, color)
			} else if(g.trailMap[i][j][3] > 0.0){
				var color color.RGBA
				color.R = uint8(fourthColor[0] * 255)
				color.G = uint8(fourthColor[1] * 255)
				color.B = uint8(fourthColor[2] * 255)
				color.A = uint8(g.trailMap[i][j][3] * 255)
				ebitenutil.DrawRect(screen, float64(i), float64(j), 1, 1, color)
			}


		}
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(types.GetWindowDefault().Width), int(types.GetWindowDefault().Height)
}
