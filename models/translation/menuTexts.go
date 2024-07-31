package translation

import "religare/models"

const (
	WelcomeMessageMenu TextIndex = iota
	SelectGeneratorMenu
	BypassValidationMenu
	StartMessageMenu
)

var MenuTexts = LanguageMap{
	models.English: Texts{
		WelcomeMessageMenu: "Welcome to Religare! This application aims to allow communication " +
			"with entities capable of manipulating the eletomagnectic spectrum. Once you've selected " +
			"the correct generator and validation mode, continue your communication experiments as usual. " +
			"Messages should appear on the screen as soon as a valid communication is received. \n" +
			"For more information, please refer to our usage manual (portuguese only) https://gentle-aura-fd4.notion.site/Manual-de-Uso-Religare-7e145f899f944cbc9ccc713ee65bb772?pvs=74",
		SelectGeneratorMenu: "Please select one of the supported generators:\n" +
			"1- Generate using Wifi Signal (Default)\n" +
			"2- Generate using random signals\n" +
			"3- Use text input (WARNING: This is for tests only)\n",
		BypassValidationMenu: "Should Religare validate the incomming signal?\n" +
			"(Bypassing signal validation means that all data received will be displayed on screen)\n" +
			"1- No (Default)\n" +
			"2- Yes (Only recommended if validation is not producing results)",
		StartMessageMenu: "Ok! Religare will now begin!\n" +
			"If you chose to validate the signal (no bypass), the screen will stay still untill a valid " +
			"message appears. \n" +
			"To stop the application and return to the mode selection menu, press CTRL + C",
	},
	models.Portuguese: Texts{
		WelcomeMessageMenu: "Bem-vindo ao Religare! Este aplicativo tem como objetivo permitir a comunicação " +
			"com entidades capazes de manipular o espectro eletromagnético. Depois de selecionar " +
			"o gerador correto e o modo de validação, continue seus experimentos de comunicação como de costume. " +
			"As mensagens devem aparecer na tela assim que uma comunicação válida for recebida. \n" +
			"Para mais informações, consulte nosso manual de uso em https://gentle-aura-fd4.notion.site/Manual-de-Uso-Religare-7e145f899f944cbc9ccc713ee65bb772?pvs=74",
		SelectGeneratorMenu: "Por favor, selecione um dos geradores suportados:\n" +
			"1- Gerar usando sinal Wifi (Padrão)\n" +
			"2- Gerar usando sinais aleatórios\n" +
			"3- Usar entrada de texto (ATENÇÃO: Isso é apenas para testes)\n",
		BypassValidationMenu: "O Religare deve validar o sinal recebido?\n" +
			"(Bypassing a validação do sinal significa que todos os dados recebidos serão exibidos na tela)\n" +
			"1- Não (Padrão)\n" +
			"2- Sim (Recomendado apenas se a validação não estiver produzindo resultados)",
		StartMessageMenu: "Ok! O Religare começará agora!\n" +
			"Se você escolheu validar o sinal (sem bypass), a tela permanecerá parada até que uma mensagem válida " +
			"apareça. \n" +
			"Para interromper o aplicativo e retornar ao menu de seleção de modo, pressione CTRL + C",
	},
}
