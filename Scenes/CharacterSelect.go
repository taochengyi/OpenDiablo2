package Scenes

import (
	"github.com/essial/OpenDiablo2/Common"
	"github.com/essial/OpenDiablo2/Palettes"
	"github.com/essial/OpenDiablo2/ResourcePaths"
	"github.com/essial/OpenDiablo2/Sound"
	"github.com/essial/OpenDiablo2/UI"
	"github.com/hajimehoshi/ebiten"
)

type CharacterSelect struct {
	uiManager         *UI.Manager
	soundManager      *Sound.Manager
	fileProvider      Common.FileProvider
	sceneProvider     SceneProvider
	background        *Common.Sprite
	newCharButton     *UI.Button
	convertCharButton *UI.Button
	deleteCharButton  *UI.Button
	exitButton        *UI.Button
	okButton          *UI.Button
}

func CreateCharacterSelect(
	fileProvider Common.FileProvider,
	sceneProvider SceneProvider,
	uiManager *UI.Manager,
	soundManager *Sound.Manager,
) *CharacterSelect {
	result := &CharacterSelect{
		uiManager:     uiManager,
		sceneProvider: sceneProvider,
		fileProvider:  fileProvider,
		soundManager:  soundManager,
	}
	return result
}

func (v *CharacterSelect) Load() []func() {
	v.soundManager.PlayBGM(ResourcePaths.BGMTitle)
	return []func(){
		func() {
			v.background = v.fileProvider.LoadSprite(ResourcePaths.CharacterSelectionBackground, Palettes.Sky)
			v.background.MoveTo(0, 0)
		},
		func() {
			v.newCharButton = UI.CreateButton(UI.ButtonTypeTall, v.fileProvider, Common.CombineStrings(Common.SplitIntoLinesWithMaxWidth(Common.TranslateString("#831"), 13)))
			v.newCharButton.MoveTo(33, 468)
			v.newCharButton.OnActivated(func() { v.onNewCharButtonClicked() })
			v.uiManager.AddWidget(v.newCharButton)
		},
		func() {
			v.convertCharButton = UI.CreateButton(UI.ButtonTypeTall, v.fileProvider, Common.CombineStrings(Common.SplitIntoLinesWithMaxWidth(Common.TranslateString("#825"), 13)))
			v.convertCharButton.MoveTo(233, 468)
			v.convertCharButton.SetEnabled(false)
			v.uiManager.AddWidget(v.convertCharButton)
		},
		func() {
			v.deleteCharButton = UI.CreateButton(UI.ButtonTypeTall, v.fileProvider, Common.CombineStrings(Common.SplitIntoLinesWithMaxWidth(Common.TranslateString("#832"), 13)))
			v.deleteCharButton.MoveTo(433, 468)
			v.deleteCharButton.SetEnabled(false)
			v.uiManager.AddWidget(v.deleteCharButton)
		},
		func() {
			v.exitButton = UI.CreateButton(UI.ButtonTypeMedium, v.fileProvider, Common.TranslateString("#970"))
			v.exitButton.MoveTo(33, 537)
			v.exitButton.OnActivated(func() { v.onExitButtonClicked() })
			v.uiManager.AddWidget(v.exitButton)
		},
		func() {
			v.okButton = UI.CreateButton(UI.ButtonTypeMedium, v.fileProvider, Common.TranslateString("#971"))
			v.okButton.MoveTo(625, 537)
			v.okButton.SetEnabled(false)
			v.uiManager.AddWidget(v.okButton)
		},
	}
}

func (v *CharacterSelect) onNewCharButtonClicked() {
	v.sceneProvider.SetNextScene(CreateSelectHeroClass(v.fileProvider, v.sceneProvider, v.uiManager, v.soundManager))
}

func (v *CharacterSelect) onExitButtonClicked() {
	mainMenu := CreateMainMenu(v.fileProvider, v.sceneProvider, v.uiManager, v.soundManager)
	mainMenu.ShowTrademarkScreen = false
	v.sceneProvider.SetNextScene(mainMenu)
}

func (v *CharacterSelect) Unload() {
}

func (v *CharacterSelect) Render(screen *ebiten.Image) {
	v.background.DrawSegments(screen, 4, 3, 0)
}

func (v *CharacterSelect) Update(tickTime float64) {
}
