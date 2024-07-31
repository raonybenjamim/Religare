package translation

import "religare/customTypes"

const (
	WelcomeMessageMenu TextIndex = iota
	SelectGeneratorMenu
	BypassValidationMenu
	StartMessageMenu
)

var MenuTexts = LanguageMap{
	customTypes.English: Texts{
		WelcomeMessageMenu: "Welcome to Religare! This application aims to allow communication " +
			"with entities capable of manipulating the eletomagnectic spectrum. Once you've selected " +
			"the correct generator and validation mode, continue your communication experiments as usual. " +
			"Messages should appear on the screen as soon as a valid communication is received. \n" +
			"For more information, please refer to our usage manual (portuguese only) https://gentle-aura-fd4.notion.site/Manual-de-Uso-Religare-7e145f899f944cbc9ccc713ee65bb772?pvs=74",
		SelectGeneratorMenu: "Please select one of the supported generators:\n" +
			"0- Generate using Wifi Signal (Default)\n" +
			"1- Generate using random signals\n" +
			"2- Use text input (WARNING: This is for tests only)\n",
		BypassValidationMenu: "Should Religare validate the incomming signal?\n" +
			"(Bypassing signal validation means that all data received will be displayed on screen)\n" +
			"0- Yes (Default)\n" +
			"1- No (Only recommended if validation is not producing results)",
		StartMessageMenu: "Ok! Religare will now begin!\n" +
			"If you chose to validate the signal (no bypass), the screen will stay still untill a valid " +
			"message appears. \n" +
			"To stop the application and return to the mode selection menu, press CTRL + C",
	},
	customTypes.Portuguese: Texts{
		WelcomeMessageMenu: "Bem-vindo ao Religare! Este aplicativo tem como objetivo permitir a comunicação " +
			"com entidades capazes de manipular o espectro eletromagnético. Depois de selecionar " +
			"o gerador correto e o modo de validação, continue seus experimentos de comunicação como de costume. " +
			"As mensagens devem aparecer na tela assim que uma comunicação válida for recebida. \n" +
			"Para mais informações, consulte nosso manual de uso em https://gentle-aura-fd4.notion.site/Manual-de-Uso-Religare-7e145f899f944cbc9ccc713ee65bb772?pvs=74",
		SelectGeneratorMenu: "Por favor, selecione um dos geradores suportados:\n" +
			"0- Gerar usando sinal Wifi (Padrão)\n" +
			"1- Gerar usando sinais aleatórios\n" +
			"2- Usar entrada de texto (ATENÇÃO: Isso é apenas para testes)\n",
		BypassValidationMenu: "O Religare deve validar o sinal recebido?\n" +
			"(Bypassing a validação do sinal significa que todos os dados recebidos serão exibidos na tela)\n" +
			"0- Sim (Padrão)\n" +
			"1- Não (Recomendado apenas se a validação não estiver produzindo resultados)",
		StartMessageMenu: "Ok! O Religare começará agora!\n" +
			"Se você escolheu validar o sinal (sem bypass), a tela permanecerá parada até que uma mensagem válida " +
			"apareça. \n" +
			"Para interromper o aplicativo e retornar ao menu de seleção de modo, pressione CTRL + C",
	},
}
