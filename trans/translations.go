package trans

import (
	"github.com/DebtsTracker/translations/emoji"
)

const adsCommandTitle = "\xE2\xAD\x90\xE2\xAD\x90\xE2\xAD\x90"

// TODO: Row 1704

// TRANS - translation string
var TRANS = map[string]map[string]string{
	"EXAMPLE": {
		"ru-RU": "ПРИМЕР",
		"en-US": "SAMPLE",
		"it-IT": "ESEMPIO",
	},

	"Jan": {
		"en-US": "Jan",
		"ru-RU": "Янв.",
		"it-IT": "Gen",
	},
	"Feb": {
		"en-US": "Feb",
		"ru-RU": "Фев.",
		"it-IT": "Feb",
	},
	"Mar": {
		"en-US": "Mar",
		"ru-RU": "Мрт.",
		"it-IT": "Mar",
	},
	"Apr": {
		"en-US": "Apr",
		"ru-RU": "Апр.",
		"it-IT": "Apr",
	},
	"May": {
		"en-US": "May",
		"ru-RU": "Май ",
		"it-IT": "Mag",
	},
	"Jun": {
		"en-US": "Jun",
		"ru-RU": "Июнь",
		"it-IT": "Giu",
	},
	"Jul": {
		"en-US": "Jul",
		"ru-RU": "Июль",
		"it-IT": "Lug",
	},
	"Aug": {
		"en-US": "Aug",
		"ru-RU": "Авг.",
		"it-IT": "Ago",
	},
	"Sep": {
		"en-US": "Sep",
		"ru-RU": "Сен.",
		"it-IT": "Sett",
	},
	"Oct": {
		"en-US": "Oct",
		"ru-RU": "Окт.",
		"it-IT": "Ott",
	},
	"Nov": {
		"en-US": "Nov",
		"ru-RU": "Нбр.",
		"it-IT": "Nov",
	},
	"Dec": {
		"en-US": "Dec",
		"ru-RU": "Дек.",
		"it-IT": "Dic",
	},
	COMMAND_START: {
		"en-US": "start",
		"ru-RU": "старт",
		"it-IT": "inizio",
	},
	COMMAND_MENU: {
		"en-US": "menu",
		"ru-RU": "меню",
	},
	COMMAND_GAVE: {
		"en-US": "gave",
		"ru-RU": "дал",
		"it-IT": "debito",
	},
	COMMAND_GOT: {
		"en-US": "got",
		"ru-RU": "взял",
		"it-IT": "credito",
	},
	COMMAND_RETURNED: {
		"en-US": "returned",
		"ru-RU": "вернул",
		"it-IT": "rientra",
	},
	COMMAND_BALANCE: {
		"en-US": "balance",
		"ru-RU": "баланс",
		"it-IT": "bilancio",
	},
	COMMAND_HISTORY: {
		"en-US": "history",
		"ru-RU": "история",
		"it-IT": "cronologia",
	},
	COMMAND_SETTINGS: {
		"en-US": "settings",
		"ru-RU": "настройки",
		"it-IT": "impostazioni",
	},
	COMMAND_HELP: {
		"en-US": "help",
		"ru-RU": "помощь",
		"it-IT": "aiuto",
	},
	COMMAND_CANCEL: {
		"en-US": "cancel",
		"ru-RU": "/отменить",
		"it-IT": "annulla",
	},
	COMMAND_CLEAR: {
		"en-US": "clear",
		"ru-RU": "очистить",
		"it-IT": "chiaro",
	},
	adsCommandTitle: {
		"ru-RU": adsCommandTitle,
		"en-US": adsCommandTitle,
		"it-IT": adsCommandTitle,
	},
	" and ": {
		"ru-RU": " и ",
		"en-US": " and ",
		"it-IT": " e ",
	},
	"MESSAGE_TEXT_OOPS_SOMETHING_WENT_WRONG": {
		"ru-RU": "Упс, что-то пошло не так... \xF0\x9F\x98\xB3",
		"en-US": "Oops, something went wrong... \xF0\x9F\x98\xB3",
		"it-IT": "Ops, qualcosa e' andato storto... \xF0\x9F\x98\xB3",
	},
	MESSAGE_TEXT_ASK_DUE: {
		"ru-RU": "Когда дата возврата?",
		"en-US": "When is the due date?",
		"it-IT": "Quand'e' la data di scadenza?",
	},
	MESSAGE_TEXT_ASK_DATE_TO_REMIND: {
		"ru-RU": `Чтобы задать дату напопинания напишите её в формате <i>ДД.MM.ГОД</i>.
<b>Например</b> для 20 января 2017 г. отправьте:
    <i>20.01.2017</i>`,

		"en-US": `To set date for next reminder please send it as a text in format of <i>DD.MM.YEAR</i>.
<b>For example</b> for 20th of January 2017 submit:
    <i>20.01.2017</i>`,

		"it-IT": `Per impostare la data per il promemoria successivo invia il messaggio con la data nel seguente formato <i>GG.MM.ANNO</i>.
<b>Esempio</b> per indicare la data 20 Gennaio 2017 inserisci:
    <i>20.01.2017</i>`,
	},
	MESSAGE_TEXT_ASK_DUE_DATE: {
		"ru-RU": `Чтобы задать дату возврата напишите её в формате <i>ДД.MM.ГОД</i>.
<b>Например</b> для 20 января 2017 г. отправьте:
    <i>20.01.2017</i>`,

		"en-US": `To set due date please send it as a text in format of <i>DD.MM.YEAR</i>.
<b>For example</b> for 20th of January 2017 submit:
    <i>20.01.2017</i>`,

		"it-IT": `Per impostare la data di scadenza invia il messaggio con la data nel seguente formato <i>GG.MM.ANNO</i>.
<b>Esempio</b> per indicare la data 20 Gennaio 2017 inserisci:
    <i>20.01.2017</i>`,
	},
	MESSAGE_TEXT_WRONG_DATE: {
		"ru-RU": "Извините, что-то не так с датой которую вы отправили.",
		"en-US": "Sorry, there is something wrong with the date you've provided.",
		"it-IT": "Mi spiace, ma c'e' qualcosa di sbagliato nella data che hai inserito.",
	},
	COMMAND_TEXT_DISABLE_REMINDER: {
		"ru-RU": "Не напоминать",
		"en-US": "No reminder",
		"it-IT": "Nessun promemoria",
	},
	COMMAND_TEXT_TOMORROW: {
		"ru-RU": "Завтра",
		"en-US": "Tomorrow",
		"it-IT": "Domani",
	},
	COMMAND_TEXT_DAY_AFTER_TOMORROW: {
		"ru-RU": "Послезавтра",
		"en-US": "Day after tomorrow",
		"it-IT": "Dopo domani",
	},
	COMMAND_TEXT_THIS_WEEK: {
		"ru-RU": "На этой неделе",
		"en-US": "This week",
		"it-IT": "Questa settimana",
	},
	COMMAND_TEXT_YES_IT_HAS_RETURN_DEADLINE: {
		"ru-RU": "Да, есть срок возврата!",
		"en-US": "Yes, it has a deadline!",
		"it-IT": "Si, c'e' una data di scadenza",
	},
	COMMAND_TEXT_NO_IT_CAN_BE_RETURNED_ANYTIME: {
		"ru-RU": "Нет, срок возврата не важен.",
		"en-US": "No, whenever is fine.",
		"it-IT": "No, quando gli pare",
	},
	COMMAND_TEXT_IT_CAN_BE_RETURNED_ANYTIME: {
		"ru-RU": "Когда-нибудь",
		"en-US": "Whenever is fine",
		"it-IT": "No, quando gli pare",
	},
	COMMAND_TEXT_IN_FEW_MINUTES: {
		"ru-RU": "Через минуту",
		"en-US": "In few minutes",
		"it-IT": "Fra pochi minuti",
	},
	COMMAND_TEXT_IN_1_WEEK: {
		"ru-RU": "Через неделю",
		"en-US": "In 1 week",
		"it-IT": "Fra una settimana",
	},
	COMMAND_TEXT_IN_1_MONTH: {
		"ru-RU": "Через месяц",
		"en-US": "In 1 month",
		"it-IT": "Fra un mese",
	},
	COMMAND_TEXT_SET_DATE: {
		"ru-RU": "Задать дату",
		"en-US": "Set date",
		"it-IT": "Imposta la data",
	},
	COMMAND_TEXT_MONDAY: {
		"ru-RU": "Понедельник",
		"en-US": "Monday",
		"it-IT": "Lunedi'",
	},
	COMMAND_TEXT_TUESDAY: {
		"ru-RU": "Вторник",
		"en-US": "Tuesday",
		"it-IT": "Martedi'",
	},
	COMMAND_TEXT_WEDNESDAY: {
		"ru-RU": "Среда",
		"en-US": "Wednesday",
		"it-IT": "Mercoledi'",
	},
	COMMAND_TEXT_THURSDAY: {
		"ru-RU": "Четверг",
		"en-US": "Thursday",
		"it-IT": "Giovedi'",
	},
	COMMAND_TEXT_FRIDAY: {
		"ru-RU": "Пятница",
		"en-US": "Friday",
		"it-IT": "Venerdi'",
	},
	COMMAND_TEXT_SATURDAY: {
		"ru-RU": "Суббота",
		"en-US": "Saturday",
		"it-IT": "Sabato",
	},
	COMMAND_TEXT_SUNDAY: {
		"ru-RU": "Воскресенье",
		"en-US": "Sunday",
		"it-IT": "Domenica",
	},
	COMMAND_TEXT_DO_NOT_SEND_RECEIPT: {
		"ru-RU": "Не отправлять квитанцию",
		"en-US": "Do not send the receipt",
		"it-IT": "Non inviare la ricevuta",
	},
	MESSAGE_TEXT_RECEIPT_WILL_NOT_BE_SENT: {
		"ru-RU": "Вы решили не отправлять квитанцию.",
		"en-US": "You've decided not to send the receipt.",
		"it-IT": "Hai scelto di non inviare la ricevuta",
	},
	COMMAND_TEXT_I_HAVE_CHANGED_MY_MIND: {
		"ru-RU": "Я передумал(а)",
		"en-US": "I've changed my mind",
		"it-IT": "Ho cambiato idea",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_TELEGRAM: {
		"ru-RU": "Отправить через Telelgram",
		"en-US": "Send by Telegram",
		"it-IT": "Invia tramite Telegram",
	},
	COMMAND_TEXT_COUNTERPARTY_HAS_NO_TELEGRAM: {
		"ru-RU": "Отправить через Viber, VK, FB, ...",
		"en-US": "Send by FB, WhatsApp, Viber, etc.",
		"it-IT": "Invia con FB, WhatsCrap, Viber, etc.",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_SMS: {
		"ru-RU": "Отправить через SMS",
		"en-US": "Send by SMS",
		"it-IT": "Invia tramite SMS",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_VK: {
		"ru-RU": "Отправить через ВКонтакте",
		"en-US": "Send throw VK.com",
		"it-IT": "Invia tramite VK.com (Facebook russo)",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_OK: {
		"ru-RU": "Отправить через Одноклассники",
		"en-US": "Send through OK",
		"it-IT": "Invia tramite OK",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_FB: {
		"ru-RU": "Отправить через Facebook",
		"en-US": "Send through Facebook",
		"it-IT": "Invia tramite Facebook",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_TWT: {
		"ru-RU": "Отправить через Twitter",
		"en-US": "Send through Twitter",
		"it-IT": "Invia tramite Twitter",
	},
	COMMAND_TEXT_CANCEL_SENDING_RECEIPT_BY_TELEGRAM: {
		"ru-RU": "Отменить отправку через Telegram",
		"en-US": "Cancel sending receipt by Telegram",
		"it-IT": "Annulla l'invio tramite Telegram",
	},
	COMMAND_TEXT_MAIN_MENU_TITLE: {
		"ru-RU": "Главное /меню",
		"en-US": "Main /menu",
		"it-IT": "Menu' /menu",
	},
	MESSAGE_TEXT_NOTHING_TO_CANCEL: {
		"ru-RU": "Нечего отменять.",
		"en-US": "Nothing to cancel.",
		"it-IT": "Nulla da annullare.",
	},
	MESSAGE_TEXT_TRANSFER_CREATION_CANCELED: {
		"ru-RU": "Создание записи о долге отменено.",
		"en-US": "Creation of debt record has been canceled.",
		"it-IT": "Creazione di un debito/credito e' stato cancellato",
	},
	COMMAND_TEXT_SHOW_ALL_CONTACTS: {
		"ru-RU": "Показать все...",
		"en-US": "Show all...",
		"it-IT": "Mostra tutto...",
	},
	COMMAND_TEXT_ADD_YOUR_OWN_OPTION: {
		"ru-RU": "Что-то другое",
		"en-US": "Something else",
		"it-IT": "Qualcos'altro",
	},
	MESSAGE_TEXT_REMINDER_ASK_IF_RETURNED: {
		"ru-RU": "Был ли этот долг возвращён?",
		"en-US": "Have this debt been returned?",
		"it-IT": "Questo debito e' stato saldato?",
	},
	MESSAGE_TEXT_ASK_WHEN_TO_REMIND_AGAIN: {
		"ru-RU": "Когда вам напомнить об этом долге ещё раз?",
		"en-US": "When should we remind you about this debt again?",
		"it-IT": "Quando devo ricordarti di questo debito?",
	},
	MESSAGE_TEXT_REPLIED_DEBT_RETURNED_FULLY: {
		"ru-RU": "Вы ответили что долг возвращён полностью.",
		"en-US": "You've replied back that debt has been returned fully.",
		"it-IT": "Hai confermato che il debito e' stato saldato.",
	},
	MESSAGE_TEXT_DEBT_IS_RETURNED: {
		"ru-RU": "Долг возвращён полностью.",
		"en-US": "The debt has been returned fully.",
		"it-IT": "Il debito e' stato saldato.",
	},
	MESSAGE_TEXT_DETAILS_ARE_HERE: {
		"ru-RU": "Подробности тут: %v",
		"en-US": "Details here: %v",
		"it-IT": "Dettagli qui: %v",
	},
	MESSAGE_TEXT_REMINDER: {
		"ru-RU": "Напоминание",
		"en-US": "Reminder",
		"it-IT": "Promemoria",
	},
	MESSAGE_TEXT_REMINDER_SET: {
		"ru-RU": "Напоминание установлено на: %v",
		"en-US": "Reminder set for: %v",
		"it-IT": "Imposta promemoria per: %v",
	},
	MESSAGE_TEXT_REMINDER_DISABLED: {
		"ru-RU": "Напоминания об этом долге отключены.",
		"en-US": "You've disabled reminders for this debt.",
		"it-IT": "Hai disabilitato il promemoria per questo debito.",
	},
	MESSAGE_TEXT_REMINDER_ALREADY_RESCHEDULED: {
		"ru-RU": "Напоминание об этом долге уже перенесено.",
		"en-US": "You've already rescheduled this reminder.",
		"it-IT": "Hai gia' impostato questo promemoria",
	},
	COMMAND_TEXT_REMINDER_RETURNED_IN_FULL: {
		"ru-RU": "Да, возвращено полностью",
		"en-US": "Yes, return in full",
		"it-IT": "Si, completamento saldato",
	},
	COMMAND_TEXT_REMINDER_RETURNED_PARTIALLY: {
		"ru-RU": "Возврашено частично",
		"en-US": "Returned partially",
		"it-IT": "Parzialmente saldato",
	},
	COMMAND_TEXT_REMINDER_NOT_RETURNED: {
		"ru-RU": "Не возвращено",
		"en-US": "Not returned",
		"it-IT": "Debito non saldato",
	},
	MESSAGE_TEXT_YOU_REPLIED: {
		"ru-RU": "Вы ответили: %v",
		"en-US": "You've replied: %v",
		"it-IT": "Hai risposto: %v",
	},
	"book": {
		"ru-RU": "книгу",
		"en-US": "book",
		"it-IT": "Libro",
	},
	"MESSAGE_TEXT_I_DID_NOT_UNDERSTAND_THE_COMMAND": {
		"ru-RU": "\xF0\x9F\x98\xB3 Извините, я не понял вашу команду. Возможно я немного туповат...\n\nЧтобы начать сначала нажмите /menu",
		"en-US": "\xF0\x9F\x98\xB3 Sorry, I did not understand your order. May be I'm a little bit dumb...\n\nYou can return to main /menu",
		"it-IT": "\xF0\x9F\x98\xB3 Scusami ma non ho capito cosa vuoi. Sono ancora un po' sciocco...\n\nPuoi ritornare al Menu con /menu",
	},
	"COMMAND_TEXT_LANGUAGE": {
		"ru-RU": "/Язык приложения",
		"en-US": "App /language",
		"it-IT": "Lingua /language",
	},
	"/start": {
		"ru-RU": "/старт",
		"en-US": "/start",
		"it-IT": "/start",

	},
	COMMAND_TEXT_DUE_RETURNS: {
		"ru-RU": "Предстоящие платежи",
		"en-US": "Due returns",
		"it-IT": "Debiti",
	},
	MESSAGE_TEXT_OVERDUE_RETURNS_HEADER: {
		"ru-RU": "<b>Просроченные долги:</b>",
		"en-US": "<b>Overdue debts:</b>",
		"it-IT": "<b>Debiti in ritardo:</b>",
	},
	MESSAGE_TEXT_DUE_RETURNS_HEADER: {
		"ru-RU": "<b>Ближайшие долги к возрату:</b>",
		"en-US": "<b>Closest debts to return:</b>",
		"it-IT": "<b>Debiti in scadenza:</b>",
	},
	MESSAGE_TEXT_DUE_RETURNS_ROW_BY_USER: {
		"ru-RU": "%v ожидает от вас возврата %v через %v",
		"en-US": "%v expects %v from you in %v",
		"it-IT": "%v aspetta %v da te entro il %v",
	},
	MESSAGE_TEXT_DUE_RETURNS_ROW_BY_COUNTERPARTY: {
		"ru-RU": "Вы ожидаете от %v возврата %v через %v",
		"en-US": "You expect %v to return %v to you in %v",
		"it-IT": "Stai aspettando %v che ti dia %v entro il %v",
	},

	MESSAGE_TEXT_DUE_RETURNS_EMPTY: {
		"ru-RU": "У вас нет долгов с назначеным сроком к возврату.",
		"en-US": "You have no debts with set due date.",
		"it-IT": "Non hai debiti con una data di scadenza.",
	},
	COMMAND_TEXT_GAVE: {
		"ru-RU": "/Дал",
		"en-US": "/Gave",
		"it-IT": "/Debito",
	},
	COMMAND_TEXT_GOT: {
		"ru-RU": "/Взял",
		"en-US": "/Got",
		"it-IT": "/Credito",
	},
	COMMAND_TEXT_RETURN: {
		"ru-RU": "/Вернул",
		"en-US": "/Returned",
		"it-IT": "/Rientra",
	},
	COMMAND_TEXT_BALANCE: {
		"ru-RU": "/Баланс",
		"en-US": "/Balance",
		"it-IT": "/Bilancio",
	},
	COMMAND_TEXT_SETTING: {
		"ru-RU": "/Настройки",
		"en-US": "/Settings",
		"it-IT": "/Impostazioni",
	},
	COMMAND_TEXT_HIGH_FIVE: {
		"ru-RU": "Дать пять!",
		"en-US": "High five!",
		"it-IT": "Batti 5 bro!",
	},
	COMMAND_TEXT_CHANGE_LANG: {
		"ru-RU": "/Язык",
		"en-US": "/Language",
		"it-IT": "/Lingua",
	},
	COMMAND_TEXT_HELP: {
		"ru-RU": "/Помощь",
		"en-US": "/Help",
		"it-IT": "/Aiuto",
	},
	COMMAND_TEXT_HISTORY: {
		"ru-RU": "/История",
		"en-US": "/History",
		"it-IT": "/Cronologia",
	},
	COMMAND_TEXT_CANCEL: {
		"ru-RU": "/Отменить",
		"en-US": "/Cancel",
		"it-IT": "/Annulla",
	},
	COMMAND_TEXT_SETTINGS_PRIMARY_CURRENCY: {
		"ru-RU": "Основная валюта",
		"en-US": "Primary currency",
		"it-IT": "Valuta principale",
	},
	COMMAND_TEXT_NEW_COUNTERPARTY: {
		"ru-RU": "Добавить",
		"en-US": "Add new",
		"it-IT": "Aggiungi nuovo",
	},
	MESSAGE_TEXT_LOGIN_CODE: {
		"ru-RU": "Ваш код для входа в приложение: <b>%v</b>",
		"en-US": "Your code for signing in to app: <b>%v</b>",
		"it-IT": "Il tuo codice per accedere all'app e': <b>%v</b>",
	},
	MESSAGE_TEXT_ASK_NEW_COUNTERPARTY_NAME: {
		"ru-RU": `<b>Имя для нового контакта</b>
Напишите сами или выберите из своей адресной книги (<i>через иконку "скрепка"</i>).

<i>Отправьте '.' для отмены</i>`,
		"en-US": `Please enter a name for the new contact:
You can type manually or choose from your address book (<i>through "clip" icon</i>).

<i>Send '.' to cancel</i>`,

		"it-IT": `Inserisci un nome per il nuovo contatto:
Puoi digitarlo o sceglierlo dalla tua rubrica (<i>attraverso l'icona "clip"</i>).

<i>Invia '.' per annullare</i>`,
	},
	MESSAGE_TEXT_TRANSFER_IS_CREATING: {
		"ru-RU": "Создаю запись...",
		"en-US": "Creating transfer...",
		"it-IT": "Sto creando la nuova voce...",
	},
	COMMAND_TEXT_PLEASE_WAIT: {
		"ru-RU": "Пожалуйста подождите",
		"en-US": "Please wait",
		"it-IT": "Aspetta per favore",
	},
	MESSAGE_TEXT_PLEASE_WAIT: {
		"ru-RU": "Пожалуйста подождите...",
		"en-US": "Please wait...",
		"it-IT": "Aspetta per favore...",
	},
	MESSAGE_TEXT_SELF_ACKNOWLEDGEMENT: {
		"ru-RU": "Подтверждение ожидается от %v",
		"en-US": "Confirm expected from %v",
		"it-IT": "Conferma in attesa da %v",
	},
	MESSAGE_TEXT_TRANSFER_ACCEPTED_BY_YOU: {
		"ru-RU": "Вы подтвердили эту транзакцию.",
		"en-US": "You've accepted this transaction.",
		"it-IT": "Hai accettato il debito/credito.",
	},
	MESSAGE_TEXT_TRANSFER_DECLINED_BY_YOU: {
		"ru-RU": "Вы НЕ подтвердили эту транзакцию.",
		"en-US": "You've declined this transaction.",
		"it-IT": "Hai rifiutato il debito/credito",
	},
	MESSAGE_TEXT_TRANSFER_ACCEPTED_BY_COUNTERPARTY: {
		"ru-RU": "%v подтвердил(a) вашу транзакцию:",
		"en-US": "%v accepted your transaction:",
		"it-IT": "%v ha accettato il tuo credito/debito:",
	},
	MESSAGE_TEXT_TRANSFER_DECLINED_BY_COUNTERPARTY: {
		"ru-RU": "%v НЕ подтвердил(a) вашу транзакцию.",
		"en-US": "%v declined your transaction.",
		"it-IT": "%v ha rifiutato il tuo credito/debito.",
	},
	COMMAND_TEXT_SUBSCRIBE_TO_APP: {
		"ru-RU": "Хочу приложение!",
		"en-US": "I want the app!",
		"it-IT": "Voglio l'aplicazione!",
	},
	COMMAND_TEXT_I_AM_FINE_WITH_BOT: {
		"ru-RU": "Меня вполне устраивает бот!",
		"en-US": "I'm fine with just the bot!",
		"it-IT": "Mi accontento del bot per ora!",
	},
	MESSAGE_TEXT_SUBSCRIBED_TO_APP: {
		"ru-RU": "Мы сообщим вам когда приложение будет доступно для загруки.",
		"en-US": "We'll let you know once the app is available for download.",
		"it-IT": "Ti faremo sapere non appena l'applicazione sara' disponibile al download.",
	},
	MESSAGE_TEXT_NOT_INTERESTED_IN_APP: {
		"ru-RU": "Чтож, мы рады что вас устраивает наш бот и нет необходимости загружать приложение.",
		"en-US": "Well, we are happy our bot is good enough and there is no need to download an app.",
		"it-IT": "Bene, siamo contenti che il nostro bot sia di tuo gradimento e non hai bisogno di scaricare l'applicazione.",
	},
	MESSAGE_TEXT_YOUR_AD_COULD_BE_HERE: {
		"ru-RU": "Здесь можно <a href>разместить рекламу</a>",
		"en-US": "You can <a href>advertise here</a>",
		"it-IT": "Puoi <a href>pubblicizzare qui</a>",
	},
	MESSAGE_TEXT_YOUR_ABOUT_ADS: {
		"ru-RU": emoji.ROBOT_ICON + `: Я конечно обоятельный робот, но пользоваться специализированным приложением бывает удобнее. Оно ещё не готово для общего доступа, но уже сейчас можно посмотреть как будет выглядеть: <a href="https://debtstracker.io/ru/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/ru/</a>

Хотите получить оповещение когда оно выйдет?`,
		"en-US": emoji.ROBOT_ICON + `: I'm a good robot, for sure. But sometimes it is more convinient to use a nice specialized app. It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/en/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/en/</a>

Do you want to get an invite when it gets released?`,

		"it-IT": emoji.ROBOT_ICON + `: Di sicuro son un bravo bot, ma alcune volte e' piu' conveniente usare un'applicazione specializzata. Non e' ancora pronta per la pubblicazione ma puoi controllare l'avanzamento a questo indirizzo: <a href="https://debtstracker.io/it/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/it/</a>

Vuoi essere invitato non appena viene rilasciata?`,

	},
	MESSAGE_TEXT_INVALID_FLOAT: {
		"ru-RU": "Извините, но вы можете использовать только числа в качестве суммы/количества (<i>до 2х знаком после запятой</i>).",
		"en-US": "Sorry, but you can use just numbers as amount/quantity (<i>with up to 2 digits after point</i>).",
		"it-IT": "Spiacente, puoi utilizzare solo numeri (<i>con un massimo di 2 numeri dopo il punto</i>).",
	},
	MESSAGE_TEXT_ASK_LENDING_TYPE: {
		"ru-RU": "<b>Что вы дали в долг?</b>",
		"en-US": "<b>What did you lend to someone?</b>",
		"it-IT": "<b>Che cos'hai prestato?</b>",
	},
	MESSAGE_TEXT_CHOOSE_CURRENCY: {
		"ru-RU": `Выберите из меню внизу экрана или <a>выберите валюту из списка</a>.

Если ни один из стандартных вариантов не подходит просто напишите текстом. Например: "<i>яблоко</i>".`,

		"en-US": `Please choose from the options below or <a>select a currency from the list</a>.

If standard options are not enough simply send a text. For example: "<i>apple</i>".`,

		"it-IT": `Scegli dalle opzioni qui sotto o <a>seleziona una moneta dalla lista</a>.

Se le opzioni standard non bastano semplicemente invia un testo. Per esempio: "<i>mele</i>".`,
	},
	MESSAGE_TEXT_ASK_LENDING_AMOUNT: {
		"ru-RU": "Сколько <b>%v</b> вы дали в долг?\n(<i>отправьте '.' чтобы отменить</i>)",
		"en-US": "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)",
		"it-IT": "Quanto hai prestato a <b>%v</b>?\n(<i>invia '.' per annullare</i>)",
	},
	MESSAGE_TEXT_ASK_LENDING_COUNTERPARTY: {
		"ru-RU": "Кому вы дали в долг <b>%v</b>?\n(<i>отправьте '.' чтобы отменить</i>)",
		"en-US": "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		"it-IT": "A chi hai prestato <b>%v</b>?\n(<i>invia '.' per annullare</i>)",
	},
	MESSAGE_TEXT_ASK_BORROWING_TYPE: {
		"ru-RU": "Что вы взяли в долг?",
		"en-US": "What did you lend?",
		"it-IT": "Cos'hai prestato?",
	},
	MESSAGE_TEXT_ASK_BORROWING_AMOUNT: {
		"ru-RU": "Сколько <b>%v</b> вы взяли в долг?\n(<i>отправьте '.' чтобы отменить</i>)",
		"en-US": "How much <b>%v</b> did you borrow?\n(<i>send '.' to cancel</i>)",
		"it-IT": "Quanto ti ha prestato <b>%v</b>?\n(<i>invia '.' per annullare</i>)",
	},
	MESSAGE_TEXT_ASK_BORROWING_COUNTERPARTY: {
		"ru-RU": "У кого вы взяли в долг <b>%v</b>?\n(<i>отправьте '.' чтобы отменить</i>)",
		"en-US": "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		"it-IT": "Chi ti ha prestato <b>%v</b>?\n(<i>invia '.' per annullare</i>)",
	},
	MESSAGE_TEXT_YOU_CAN_SEND_RECEIPT: {
		"ru-RU": "Отправить <a receipt>квитанцию</a> для <a counterparty>%v</a>?",
		"en-US": "Should we send a <a receipt>receipt</a> to <a counterparty>%v</a>?",
		"it-IT": "Devo inviare una <a receipt>notifica</a> a <a counterparty>%v</a>?",
	},
	MESSAGE_TEXT_YOU_CAN_SEND_RECEIPT_TO_YOURSELF_BY_SMS: {
		"ru-RU": "К сожалению отправка квитанцию себе по СМС в данный момент отключена. Но вы можете отправить её для %v.",
		"en-US": "Sorry, sending a receipt to yourself by SMS is not available at the moment. You can send it to %v though.",
		"it-IT": "Spiacente ma inviarsi da soli una notifica tramite SMS non e' al momento disponibile. Pero' puoi inviarla a %v.",
	},
	MESSAGE_TEXT_RECEIPT_IS_SENDING_BY_TELEGRAM: {
		"ru-RU": "Отправляем для %v извещение через Telegram...",
		"en-US": "We are sending receipt to %v by Telegram...",
		"it-IT": "Sto inviando la notifica a %v tramite Telegram...",
	},
	MESSAGE_TEXT_RECEIPT_NEW_DEBT_FROM_USER: {
		"ru-RU": "{{.Counterparty}} взял(а) в долг {{.Amount}}.",
		"en-US": "{{.Counterparty}} borrowed from you {{.Amount}}.",
		"it-IT": "{{.Counterparty}} ti deve dare {{.Amount}}.",
		//"it-IT": "{{.Counterparty}} ha preso in prestito da te {{.Amount}}.",
	},
	MESSAGE_TEXT_RECEIPT_NEW_DEBT_TO_USER: {
		"ru-RU": "{{.Counterparty}} дал(а) вам в долг {{.Amount}}.",
		"en-US": "{{.Counterparty}} lended to you {{.Amount}}.",
		"it-IT": "{{.Counterparty}} ti ha prestato {{.Amount}}.",
	},
	MESSAGE_TEXT_RECEIPT_RETURN_FROM_USER: {
		"ru-RU": "Вы вернули долг - {{.Counterparty}} получил(а) {{.Amount}}.",
		"en-US": "You returned {{.Amount}} to {{.Counterparty}}.",
		"it-IT": "Hai ridato {{.Amount}} a {{.Counterparty}}.",
	},
	MESSAGE_TEXT_RECEIPT_RETURN_TO_USER: {
		"ru-RU": "{{.Counterparty}} вернул вам {{.Amount}}.",
		"en-US": "{{.Counterparty}} returned to you {{.Amount}}.",
		"it-IT": "{{.Counterparty}} ti ha ridato {{.Amount}}.",
	},
	MESSAGE_TEXT_DUE_ON: {
		"ru-RU": "<b>Вернуть до</b>: %v",
		"en-US": "<b>Return till</b>: %v",
		"it-IT": "<b>Dare a</b>: %v",
	},
	MESSAGE_TEXT_NOTE: {
		"ru-RU": "Заметка",
		"en-US": "Note",
		"it-IT": "Nota",
	},
	MESSAGE_TEXT_COMMENT: {
		"ru-RU": "Комментарий",
		"en-US": "Comment",
		"it-IT": "Commento",
	},
	MESSAGE_TEXT_LOGIN_TO_WEB_APP: {
		"ru-RU": `Перейдите по <a>ссылке</a> чтобы запустить web-приложение.`,
		"en-US": `Click to <a>sign in</a> to web-app.`,
		"it-IT": `Clicca su <a>sign in</a> nella web-app.`,
	},
	MESSAGE_TEXT_DO_YOU_LIKE_OUR_BOT: {
		"ru-RU": "Вам нравится @{{bot}}?",
		"en-US": "Do you like @{{bot}}?",
		"it-IT": "Divertito con @{{bot}}?",
	},
	COMMAND_TEXT_YES_EXCLAMATION: {
		"ru-RU": "%v Да!",
		"en-US": "%v Yes!",
		"it-IT": "%v Si!",
	},
	COMMAND_TEXT_YES: {
		"ru-RU": "%v Да",
		"en-US": "%v Yes",
		"it-IT": "%v Si",
	},
	COMMAND_TEXT_NO: {
		"ru-RU": "%v Нет",
		"en-US": "%v No",
		"it-IT": "%v No",
	},
	COMMAND_TEXT_NOT_TOO_MUCH: {
		"ru-RU": "%v Не очень",
		"en-US": "%v Not too much",
		"it-IT": "%v Non troppo",
	},
	COMMAND_TEXT_FEEDBACK: {
		"ru-RU": "/Отзыв",
		"en-US": "/Feedback",
		"it-IT": "/Risposta",
	},
	COMMAND_TEXT_WRITE_FEEDBACK: {
		"ru-RU": "%v Написать отзыв",
		"en-US": "%v Write feedback",
		"it-IT": "%v Scrivi commenti",
	},
	MESSAGE_TEXT_THANKS: {
		"ru-RU": "🙏 Спасибо!",
		"en-US": "🙏 Thanks!",
		"it-IT": "🙏 Grazie!",
	},
	MESSAGE_TEXT_PLEASE_SEND_TEXT: {
		"ru-RU": "Пожалуйста отправьте текст.",
		"en-US": "Please send text.",
		"it-IT": "Si prega di inviare il testo.",
	},
	MESSAGE_TEXT_CAN_YOU_RATE_AT_STOREBOT: {
		"ru-RU": `🤖 Можете поставить ему высокую оценку и хороший отзыв в каталоге ботов Store Bot?

‎Это займет меньше минуты вашего времени! 😇`,
		"en-US": `🤖 Can you rate it high and write a good review in bots catalog Store Bot?

‎It will take less than a minute of your time! 😇`,
		"it-IT": `🤖 Puoi votarlo in alto e scrivere una buona revisione nel catalogo bot Bot Store?

Ci vorrà meno di un minuto del tuo tempo! 😇`,
	},
	MESSAGE_TEXT_ASK_TO_WRITE_FEEDBACK_WITHIN_MESSENGER: {
		"ru-RU": "Поделитесь вашими мыслями (на русском или английском) о том, что нужно сделать, чтобы бот стал лучше:",
		"en-US": "‎Share your thoughts (in English or Russian) about what could be done to make the bot better:",
		"it-IT": "Condividi i tuoi pensieri (in Inglese o Russo) su come sarebbe migliore secondo te il bot:",
	},
	MESSAGE_TEXT_HOW_TO_RATE_AT_STOREBOT: {
		"ru-RU": `‎<b>Как поставить оценку в три простых шага:</b>

1. Перейдите по ссылке, чтобы оставить оценку и отзыв:
‎https://t.me/storebot?start={{bot}}

‎2. Нажмите на кнопку "⭐️⭐️⭐️⭐️⭐️"

‎3. Напишите сообщение или нажмите кнопку "Пропустить этот шаг"

Спасибо вам большое! Благодаря этому о боте узнает больше людей — это служит дополнительной мотивацией для разработчиков! 😎`,

		"en-US": `<b>How to rate in 3 simple steps:</b>

1. Click on this link to rate and review:
https://t.me/storebot?start={{bot}}

‎2. Click on the "⭐️⭐️⭐️⭐️⭐️" button

‎3. Write your message or press "Skip this step" button

Thank you very much! As a result of your actions, even more people will learn about the bot. All this will serve as the additional motivation for the developers! 😎`,

		"it-IT": `<b>Come valutare in 3 semplici passaggi:</b>
‎1. Clicca su questo link per votare e lasciare una recensione:
‎https://t.me/storebot?start={{bot}}

‎2. Clicca sul "⭐️⭐️⭐️⭐️⭐️" bottone

‎3. Scrivi il tuo messaggio o premi "Salta questo step"

Grazie infinitamente! Come risultato delle tue azioni, altre persone guarderanno il bot. Dando anche un motivo in più per continuare ai developers! 😎`,
	},
	MESSAGE_TEXT_ASK_FOR_FEEDBAK: {
		"ru-RU": "Будем признетельны если вы оцените работу нашего приложения. Это займёт всего несколько секунд.",
		"en-US": "We would appreciate if tell us how we doing. It takes just few seconds.",
		"it-IT": "Ci farebbe piacere se lasciassi un voto per il nostro lavoto. Ti bastano solo alcuni secondi.",
	},
	COMMAND_TEXT_GIVE_FEEDBACK: {
		"ru-RU": "Оценить приложение",
		"en-US": "Rate this bot",
		"it-IT": "Vota questo bot",
	},
	COMMAND_TEXT_OPEN_STOREBOT_FOR_FEEDBACK: {
		"ru-RU": "Оценить на  @Storebot",
		"en-US": "Leave rating at @Storebot",
		"it-IT": "Lascia un voto a @Storebot",
	},
	MESSAGE_TEXT_ON_REFUSED_TO_RATE: {
		"ru-RU": `ОК, возможно вы смоежете поставить оценку в другой раз.

		{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

		Так же будем признательны если вы предложите любые улучшения.
		`,
		/*------------------------------------------------------------*/
		"en-US": `OK, maybe you can rate us another time.

		{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

		We also will appreciate if you suggest any improvements.
		`,
	},
	MESSAGE_TEXT_ON_FEEDBACK_POSITIVE: {
		"ru-RU": `Спасибо, мы очень старались!

		{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

		Так же будем признательны если вы <a suggest-idea>предложите улучшения</a>.
		`,
		/*------------------------------------------------------------*/
		"en-US": `Thanks, we worked hard!

		{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

		We also will appreciate if you <a suggest-idea>suggest improvements</a>.
		`,
		/*------------------------------------------------------------*/
		"it-IT": `GRAZIE MILLE, abbiamo lavoro duro!

		{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

		Sarebbe ancora piu' apprezzatto se ci <a suggest-idea>suggerisci qualche miglioramento</a>.
		`,
	},
	MESSAGE_TEXT_YOU_CAN_HELP_BY: {
		"ru-RU": `
		Вы нам очень поможете если:

		  * Оставите положительный <a storebot>отзыв в каталоге ботов</a>.

		  * Расскажите о боте своим друзьям.
		    Например во <a share-vk>ВКонтакте</a>, <a share-fb>Facebook</a> или <a share-twitter>Twitter</a>.

		  * Поддержите дальнейшую разработку - <a href="https://goo.gl/Qhh0yL">€2 через PayPal</a>
		`,
		/*------------------------------------------------------------*/
		"en-US": `
		You can help us a lot if you:

		  * Leave a positive feedback at <a storebot>directory of bots</a>.

		  * Tell about the app to your friends.
		    For example at <a share-fb>Facebook</a> or <a share-twitter>Twitter</a>.

		  * Support further development - <a href="https://goo.gl/Qhh0yL">€2 via PayPal</a> (<i>about $2.2</i>)
		`,
		/*------------------------------------------------------------*/
		"it-IT": `
		Ci aiuteresti moltissimo se:

		  * Lasci un feedback positivo alla <a storebot>pagina del bot</a>.

		  * Raccontare dell'app ai tuoi amici.
		    Per esempio su <a share-fb>Facebook</a> o su <a share-twitter>Twitter</a>.

		  * Supporta ulteriormente lo sviluppo del bot - <a href="https://goo.gl/Qhh0yL">2€ tramite PayPal</a> (<i>circa $2.2</i>)
		`,
	},
	MESSAGE_TEXT_YOU_CAN_HELP_BY_HTML: {
		"ru-RU": `
		<p>Вы нам очень поможете если:</>

		<ul>
		  <li>Оставите положительный <a storebot>отзыв в каталоге ботов</a>.</li>

		  <li>Расскажите о боте своим друзьям.
		    Например во <a share-vk>ВКонтакте</a>, <a share-fb>Facebook</a> или <a share-twitter>Twitter</a>.
		  </li>

		  <li>Поддержите дальнейшую разработку - <a href="https://goo.gl/Qhh0yL">€2 через PayPal</a></li>
		</ul>
		`,
		/*------------------------------------------------------------*/
		"en-US": `
		<p>You can help us a lot if you:</p>

		<ul>
		  <li>Leave a positive feedback at <a storebot>directory of bots</a>.</li>

		  <li>Tell about the app to your friends.
		    For example at <a share-fb>Facebook</a> or <a share-twitter>Twitter</a>.</li>

		  <li>Support further development - <a href="https://goo.gl/Qhh0yL">€2 via PayPal</a> (<i>about $2.2</i>)</li>
		</ul>
		`,
	},
	MESSAGE_TEXT_COUNTERPARTY_HAS_EMPTY_BALANCE: {
		"ru-RU": `Нулевой баланс для %v`,
		"en-US": `Balance is empty for %v`,
		"it-IT": `Il bilancio di %v e' vuoto al momento`,
	},
	MESSAGE_TEXT_ASK_TO_TRANSLATE: {
		"ru-RU": `Хотите чтобы наш бот разговаривал на другом языке? Вы можете <a>помочь с переводом</a>.`,
		"en-US": `Do you want our bot to talk in other language? You can <a>help with translation</a>.`,
		"it-IT": `Vuoi che il nostro bot parli altre lingue? Aiuta con la <a>traduzione</a>.`,
	},
	MESSAGE_TEXT_ON_FEEDBACK_NEUTRAL: {
		"ru-RU": `Чтож, мы очень старались. Ваша оценка будет передана разработчикам.

Может быть вы <a submit-bug>сообщите нам что не работает</a> или подскажите <a suggest-idea>как можно улучшить</a>?`,
		/*------------------------------------------------------------*/
		"en-US": `Well, we worked hard. You feedback will be passed to developers.

Maybe you can <a submit-bug>report your issue</a> or <a suggest-idea>suggest how we can improve</a>?`,
		/*------------------------------------------------------------*/
		"it-IT": `Bene, il nostro lavoro non e' andato in vano. Il tuo feedback sara' inoltrato agli sviluppatori.

Per caso vuoi anche <a submit-bug>segnalare un problema</a> oppure <a suggest-idea>suggerire come possiamo migliorare</a>?`,
	},
	MESSAGE_TEXT_ON_FEEDBACK_NEGATIVE: {
		"ru-RU": `Нам очень стыдно. Может быть вы <a submit-bug>подскажите что не так</a> или <a suggest-idea>предложите усовершенствования</a>?`,
		/*------------------------------------------------------------*/
		"en-US": `We are very sorry. Maybe you can <a submit-bug>let us know what is wrong</a> or <a suggest-idea>suggest how we can improve</a>?`,
		/*------------------------------------------------------------*/
		"it-IT": `Ci dispiace molto. Potresti farci sapere<a submit-bug>cosa non ti e' piaciuto</a> oppure <a suggest-idea>suggerirci come possiamo migliorare</a>?`,
	},
	COMMAND_TEXT_ASK_FOR_FEEDBACK: {
		"ru-RU": "Оцените наше приложение?",
		"en-US": "Please rate our app",
		"it-IT": "Per favore vota la nostra app",
	},
	COMMAND_TEXT_FEEDBACK_POSITIVE: {
		"ru-RU": "Да, отличное приложение!",
		"en-US": "Yes, it's a great app!",
		"it-IT": "Si, e' un app fantastica!",
	},
	COMMAND_TEXT_FEEDBACK_NEUTRAL: {
		"ru-RU": "Неплохо, но можно лучше.",
		"en-US": "Not bad but can be better.",
		"it-IT": "Non male ma potrebbe esser migliore.",
	},
	COMMAND_TEXT_FEEDBACK_NEGATIVE: {
		"ru-RU": "Не нравится",
		"en-US": "Don't like it",
		"it-IT": "Non mi piace",
	},
	COMMAND_TEXT_FEEDBACK_NOT_READY: {
		"ru-RU": "Пока не понятно",
		"en-US": "Not decided yet",
		"it-IT": "Sono indeciso",
	},
	MESSAGE_TEXT_SETTINGS: {
		"ru-RU": "Что будем настраивать?",
		"en-US": "What do you want to adjust?",
		"it-IT": "Cosa miglioreresti",
	},
	MESSAGE_TEXT_NOT_IMPLEMENTED_YET: {
		"ru-RU": "Извините, данный функционал ещё не реализован.",
		"en-US": "Sorry, this functionality is not implemented yet.",
		"it-IT": "Spiacenti ma questa funzionalita' non e' ancora attiva.",
	},
	MESSAGE_TEXT_ASK_INVITE_CHANNEL: {
		"ru-RU": "Как вы хотите получить код приглашения?",
		"en-US": "How do you want to get an invite?",
		"it-IT": "Come vuoi ricevere l'invito?",
	},
	MESSAGE_TEXT_PLEASE_ENTER_INVITE_CODE: {
		"ru-RU": "Пожалуйста введите код приглашения:",
		"en-US": "Please enter an invite code:",
		"it-IT": "Inserisci un codice invito:",
	},
	MESSAGE_TEXT_USER_EMAIL_FOR_INVITE_RECEIVED: {
		"ru-RU": "Мы отправили письмо на %v.\n\nПожалуйста откройте его и кликните на ссылку для подтверждения адреса.",
		"en-US": "We've sent a message to %v.\n\nPlease open the email and click a link to confirm your email address.",
		"it-IT": "Abbiamo inviato un messaggio a %v.\n\nPer favore apri l'email e clicca sul link per confermare il tuo indirizzo email",
	},
	MESSAGE_TEXT_USER_EMAIL_FOR_INVITE_SENT_TELEGRAM: {
		"ru-RU": "После того как откроется Telegram нажмите кнопку <b>Start</b>.",
		"en-US": "Once Telegram pop ups click the <b>Start</b> button.",
		"it-IT": "Una volta aperto il bot su telegram clicca su <b>Start</b>.",
	},
	MESSAGE_TEXT_USER_CONTACT_FOR_INVITE_RECEIVED: {
		"ru-RU": "Спасибо, вы записаны в очередь на получение приглашения.\n\nТекущее время ожидания 2-3 дня.\n\nВы можете получить приглашение сегодня если расскажите о нашем боте на Facebook.",
		"en-US": "Thanks, you've been queued for an invite.\n\nCurrent awaiting time is 2-3 days.\n\nYou can get an invite code today by sharing a link to the bot on Facebook.",
		"it-IT": "Grazie, ora sei in coda per un codice invito.\n\nTempo di attesa medio 2-3 giorni.\n\nPuoi ottenere un codice invito subito condividendo il link al bot su Facebook.",
	},
	MESSAGE_TEXT_PLEASE_PROVIDE_YOUR_EMAIL: {
		"ru-RU": "Пожалуйста напишите ваш email адрес:",
		"en-US": "Please provide your email address",
		"it-IT": "Inserisci il tuo indirizzo email:",
	},
	MESSAGE_TEXT_PLEASE_PROVIDE_YOUR_PHONE_NUMBER: {
		"ru-RU": "Пожалуйста напишите номер вашего телефона:",
		"en-US": "Please provide your phone number:",
		"it-IT": "Inserisci il tuo numero di telefono:",
	},
	MESSAGE_TEXT_WRONG_INVITE_CODE: {
		"ru-RU": "Неправильный код приглашения: %v",
		"en-US": "Wrong invite code: %v",
		"it-IT": "Codice invito: %v errato",
	},
	MESSAGE_TEXT_WRONG_EMAIL: {
		"ru-RU": "Неправильный email адрес.",
		"en-US": "Wrong email address.",
		"it-IT": "L'email inserita e' sbagliata.",
	},
	MESSAGE_TEXT_WRONG_PHONE_NUMBER: {
		"ru-RU": "Неправильный номер телефона.",
		"en-US": "Wrong phone number.",
		"it-IT": "Il numero inserito e' sbagliato.",
	},
	MESSAGE_TEXT_OK_PLEASE_TRY_AGAIN: {
		"ru-RU": "Хорошо, попробуйте ещё раз.",
		"en-US": "Ok, please try again.",
		"it-IT": "Ok, prova di nuovo.",
	},

	COMMAND_TEXT_MISTYPE_WILL_TRY_AGAIN: {
		"ru-RU": "Я опечатался, попробую ещё раз.",
		"en-US": "I've mistyped, will try again.",
		"it-IT": "Ho sbagliato, riprovo.",
	},
	COMMAND_TEXT_TELL_ME_MORE_ABOUT_INVITES: {
		"ru-RU": "Расскажите ка мне об этих кодах",
		"en-US": "Tell me more about the codes",
		"it-IT": "Ulteriori informazioni riguardo il codice invito.",
	},
	COMMAND_TEXT_INVITE_ME_BY_EMAIL: {
		"ru-RU": "Хочу код приглашения на email",
		"en-US": "Send me invite code by email",
		"it-IT": "Inviami il codice invito tramite email",
	},
	COMMAND_TEXT_INVITE_ME_BY_SMS: {
		"ru-RU": "Хочу код приглашения по SMS",
		"en-US": "Send me invite code by SMS",
		"it-IT": "Inviami il codice invito tramite SMS",
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_EMAIL: {
		"ru-RU": "Новый код приглашения на email",
		"en-US": "Send me new invite code by email",
		"it-IT": "Inviami il nuovo codice invito tramite email",
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_SMS: {
		"ru-RU": "Новый код приглашения по SMS",
		"en-US": "Send me new invite code by SMS",
		"it-IT": "Inviami il nuovo codice invito tramite SMS",
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_TELEGRAM: {
		"ru-RU": "Получить приграшение в Telegram",
		"en-US": "Send me new invite by Telegram",
		"it-IT": "Inviami il nuovo codice invito tramite Telegram",
	},
	MESSAGE_TEXT_UNKNOWN_LANGUAGE: {
		"ru-RU": "Незнакомый язык. Пожалуйста выберете один из предоставленных:",
		"en-US": "Unknown language. Please choose 1 from the options:",
		"it-IT": "Lingua socnosciuta. Per favore scegline una tra le opzioni:",
	},
	MESSAGE_TEXT_UNKNOWN_COUNTERPARTY: {
		"ru-RU": "Неизвестный контакт. Пожалуйста выберите <b>Добавить</b> если это новый контакт.",
		"en-US": "Unknown counterparty. Please choose <b>Add new</b> if it's a new contact.",
		"it-IT": "Destinatario sconosciuto. Scegli <b>Aggiugni nuovo</b> se e' un nuovo contatto.",
	},
	MESSAGE_TEXT_UNKNOWN_COUNTERPARTY_ON_RETURN: {
		"ru-RU": "Неизвестный контакт. Пожалуйста выберите из списка.",
		"en-US": "Unknown counterparty. Please choose from the list.",
		"it-IT": "Destinatario sconosciuto. Scegli dalla lista qui sotto.",
	},
	MESSAGE_TEXT_UNKNOWN_DEBT: {
		"ru-RU": "Неизвестный долг. Пожалуйста выберите из списка.",
		"en-US": "Unknown debt. Please choose from the list.",
		"it-IT": "Debito sconosciuto. Scegli dalla lista qui sotto.",
	},

	MESSAGE_TEXT_HI: { // This is the same for all languages.
		"ru-RU": `¡Hola! Hi! Привет! سلام!`,
		"en-US": `¡Hola! Hi! Привет! سلام!`,
		"it-IT": `¡Hola! Hi! Привет! سلام!`,
	},
	MESSAGE_TEXT_BACK_TO_MAIN_MENU: {
		"ru-RU": `Можно вернуться назад в главное /меню`,
		"en-US": `You can go back to main /menu`,
		"it-IT": `Puoi tornare al menu' principale tramite /menu`,
	},
	MESSAGE_TEXT_YOUR_SELECTED_PREFERRED_LANGUAGE: { // This is the same for all languages.
		"ru-RU": `Выбранный язык программы: %v`,
		"en-US": `Preferred app language: %v`,
		"it-IT": `Lingua del bot preferita: %v`,
	},
	MESSAGE_TEXT_ONBOARDING_ASK_TO_CHOOSE_LANGUAGE: {
		"ru-RU": `<b>%v</b>, на каком языке вы хотели бы общаться?

(What is your preferred language?)`,
		"en-US": `<b>%v</b>, what is your preferred language?`,
		"it-IT": `<b>%v</b> qual'e' la tua lingua madre?`,
	},
	MESSAGE_TEXT_CHOOSE_UI_LANGUAGE: {
		"ru-RU": "На каком языке вы хотели бы общаться со мной?",
		"en-US": "Which language you would like to talk to me?",
		"it-IT": "Con quale lingua vuoi chattare con me?",
	},
	MESSAGE_TEXT_LOCALE_CHANGED: {
		"ru-RU": "Вы поменяли язык на %v",
		"en-US": "You've switched language to %v",
		"it-IT": "Hai cambiato lingua in %v",
	},
	MESSAGE_TEXT_WHATS_NEXT: {
		"ru-RU": "Что будем делать дальше?",
		"en-US": "What's next?",
		"it-IT": "Ed ora? Che faccio?",
	},
	MESSAGE_TEXT_WHATS_NEXT_HINT: {
		"ru-RU": `
Если вы дали в долг воспользуйтесь командой /дал.
Если вы одолжили что-то - командой /взял.

Или воспользуйтесь меню внизу экрана.
`,
		"en-US": `
If you borrowed from someone to record it use /got.
If you lent to someone to record it use /gave.

Or use menu at the bottom.
`,

		"it-IT": `
Se qualcuno ti ha prestato qualcosa per memorizzarlo usa /got.
Se hai prestato qualcosa a qualcuno per memorizzarlo usa /gave.

O usa il menu' qui sotto.
`,
	},
	MESSAGE_TEXT_HISTORY_HEADER: {
		"ru-RU": "История",
		"en-US": "History",
		"it-IT": "Cronologia",
	},
	MESSAGE_TEXT_HISTORY_NO_RECORDS: {
		"ru-RU": "У вас пока нет ни одной записи.",
		"en-US": "You don't have any records yet.",
		"it-IT": "Non hai nulla memorizzato.",
	},
	MESSAGE_TEXT_HISTORY_LIST: {
		"ru-RU": `<b>%v</b> <i>(%d последних)</i>
─────────────
%v`,

		"en-US": `<b>%v</b> <i>(last %d):</i>
─────────────
%v`,

		"it-IT": `<b>%v</b> <i>(ultimo %d):</i>
─────────────
%v`,
	},
	MESSAGE_TEXT_BALANCE_IS_ZERO: {
		"ru-RU": "Нет записей о текущих долгах.",
		"en-US": "You have no records on current debts.",
		"it-IT": "Non hai nulla memorizzato nel debito corrente.",
	},
	MESSAGE_TEXT_BALANCE_CURRENCY_TOTAL_INTRO: {
		"ru-RU": "Всего",
		"en-US": "Total",
		"it-IT": "Totale",
	},
	MESSAGE_TEXT_PRIMARY_CURRENCY_IS_SET_TO: {
		"ru-RU": "OK, теперь я буду использовать '%v' как основную валюту.",
		"en-US": "OK, from now on I will use '%v' as a primary currency.",
		"it-IT": "OK, da ora in poi usero' '%v' come moneta principale.",
	},
	MESSAGE_TEXT_BALANCE_SINGLE_CURRENCY_COUNTERPARTY_DEBT_TO_USER: {
		"ru-RU": "<b>%v</b> - долг вам %v",
		"en-US": "<b>%v</b> - owes you %v",
		"it-IT": "<b>%v</b> - ti deve %v.",
	},
	MESSAGE_TEXT_BALANCE_CURRENCY_ROW_DEBT_TO_USER: {
		"ru-RU": "Вам должны %v",
		"en-US": "Owes to you %v",
		"it-IT": "%v e' in debito con te",
	},
	MESSAGE_TEXT_ON_RETURN_USER_DOES_NOT_OWE_ANYTHING_TO_COUNTERPARTY_ANYMORE: {
		"ru-RU": "Поздравляем! У вас не осталось долгов перед <b>%v</b>.",
		"en-US": "Congratulations! You don't owe anything more to <b>%v</b>.",
		"it-IT": "Bravo! Hai saldato il tuo debito con <b>%v</b>.",
	},
	MESSAGE_TEXT_ON_RETURN_COUNTERPARTY_DOES_NOT_OWE_ANYTHING_TO_USER_ANYMORE: {
		"ru-RU": "У <b>%v</b> больше не осталось долгов перед вами.",
		"en-US": "<b>%v</b> does not owe anything more to you.",
		"it-IT": "<b>%v</b> ha saldato il suo debito verso di te.",
	},
	MESSAGE_TEXT_BALANCE_CURRENCY_ROW_DEBT_BY_USER: {
		"ru-RU": "Вы должны %v",
		"en-US": "You owe %v",
		"it-IT": "Sei in debito con %v",
	},
	MESSAGE_TEXT_BALANCE_SINGLE_CURRENCY_COUNTERPARTY_DEBT_BY_USER: {
		"ru-RU": "<b>%v</b> - вы должны %v",
		"en-US": "<b>%v</b> - you owe %v",
		"it-IT": "<b>%v</b> - tu gli devi %v",
	},
	MESSAGE_TEXT_ASK_PRIMARY_CURRENCY: {
		"ru-RU": "Какая валюта для вас основная?",
		"en-US": "What is your primary currency?",
		"it-IT": "Qual'e' la tua valuta principale?",
	},
	MESSAGE_TEXT_FAILED_TO_DELETE_USER: {
		"ru-RU": "Не удалось удалить данные пользователя: %v",
		"en-US": "Failed to delete user: %v",
		"it-IT": "Errore durante la cancellazione dell'utente: %v",
	},
	MESSAGE_TEXT_USER_DELETED: {
		"ru-RU": "Данные пользователя удалены",
		"en-US": "User's data has been deleted",
		"it-IT": "I dati dell'utente sono stati cancellati",
	},
	MESSAGE_TEXT_PLEASE_WAIT_WHILE_WE_GENERATE_INVITE_CODE: {
		"ru-RU": "Пожалуйста подождите пока мы генерируем секретный код доступа...",
		"en-US": "Please wait a moment while we are generating a security access code...",
		"it-IT": "Aspetta un attimo mentre sto generando un codice di accesso sicuro...",
	},
	MESSAGE_TEXT_RETURN_ASK_TO_CHOOSE_COUNTERPARTY: {
		"ru-RU": "Выберете кому вы вернули долг или кто вернул его вам.",
		"en-US": "Please choose who returned the debt or to who you returned it.",
		"it-IT": "Scegli chi ha sanato il suo debito o con chi hai sanato tu il tuo debito",
	},
	MESSAGE_TEXT_CHOOSE_DEBT_THAT_HAS_BEEN_RETURNED: {
		"ru-RU": "Выберите долг который был возвращён целиком или частично.",
		"en-US": "Please choose a debt that has been returned fully or partially.",
		"it-IT": "Scegli un debito che e' stato restituito completamente o parzialmente.",
	},
	MESSAGE_TEXT_PLEASE_ACKNOWLEDGE_TRANSFER: {
		"ru-RU": "Пожалуйста подтвердите или отклоните эту транзакцию.",
		"en-US": "Please confirm or decline this transfer.",
		"it-IT": "Conferma o rifiuta questo debito/credito.",
	},
	MESSAGE_TEXT_ALREADY_ACCEPTED_TRANSFER: {
		"ru-RU": "Эта транзакция уже подтверждена.",
		"en-US": "This transfer has been accepted already.",
		"it-IT": "Questo debito/credito e' gia' stato accettato.",
	},
	MESSAGE_TEXT_ALREADY_DECLINED_TRANSFER: {
		"ru-RU": "Эта транзакция уже отклонена.",
		"en-US": "This transfer has been declined already.",
		"it-IT": "Questo debito/credito e' gia' stato rifiutato.",
	},
	MESSAGE_TEXT_RECEIPT_LINK: {
		"ru-RU": "Подробнее тут: %v",
		"en-US": "Details here: %v",
		"it-IT": "Maggiori dettagli qui: %v",
	},
	MESSAGE_TEXT_ASK_PHONE_NUMBER_OF_COUNTERPARTY: {
		"ru-RU": "Пожалуйста напишите номер телефона <b>%v</b>.",
		"en-US": "Plese provide phone number for <b>%v</b>",
		"it-IT": "Per favore fornisci il numero di telefono di <b>%v</b>",
	},
	MESSAGE_TEXT_USE_CONTACT_TO_SEND_PHONE_NUMBER: {
		"ru-RU": "Если номер телефона есть в записной книжке <b>воспользуйтесь кнопкой %v</b> (скрепка) чтобы отправить контакт.",
		"en-US": "If phone number is in your address book you can <b>use %v button</b> to send the contact.",
		"it-IT": "Se il numero e' nella tua rubrica, puoi <b> usare il pulsante %v</b> per inviare il contatto.",
	},
	MESSAGE_TEXT_ABOUT_PHONE_NUMBER_FORMAT: {
		"ru-RU": "Номер должен быть в международном формате:\n\t* Начинаться со знака '+' и кода страны\n\t* Состоять только из цифр\nПример: <pre>+</pre><b>7</b><code>999012345678</code>",
		"en-US": "The number should be in international standard:\n\t* Starts with '+' following by country code\n\t* Consist of numbers only\nExample: <pre>+</pre><b>1</b><code>999012345678</code>",
		"it-IT": "Il numero deve essere in formato internazionale:\n\t* Inizia con '+' seguito dal codice del paese (Italia +39)\n\t* \nEsempio: <pre>+</pre><b>39</b><code>34612345678</code>",
	},
	MESSAGE_TEXT_THIS_NUMBER_WILL_BE_USED_TO_SEND_RECEIPT: {
		"ru-RU": "На этот номер мы отправим SMS:",
		"en-US": "Will send an SMS to this number:",
		"it-IT": "Invieremo un SMS a questo numero:",
	},
	MESSAGE_TEXT_COUNTERPARTY_OWES_YOU_SINGLE_DEBT: {
		"ru-RU": `<b>%v</b> одалживал(а) у вас <b>%v</b>.`,
		"en-US": `<b>%v</b> owed to you <b>%v</b>.`,
		"it-IT": `<b>%v</b> ti deve <b>%v</b>.`,
	},
	MESSAGE_TEXT_YOU_OWE_TO_COUNTERPARTY_SINGLE_DEBT: {
		"ru-RU": `<b>%v</b> одалживал(а) вам <b>%v</b>.`,
		"en-US": "You owe to <b>%v</b> <b>%v</b>.",
		"it-IT": `Tu devi dare a <b>%v</b> <b>%v</b>.`,
	},
	MESSAGE_TEXT_IS_IT_RETURNED_IN_FULL: {
		"ru-RU": `Возвращено полностью?

		<i>Если частично можете сразу написать сумму.</i>`,

		"en-US": `Has this debt been returned in full?

		<i>If partially you can enter amount right away.</i>`,

		"it-IT": `Questo debito e' stato completamente saldato?

		<i>Se la risposta e' NO puoi inserire l'ammontare ora.</i>`,
	},
	MESSAGE_TEXT_PLEASE_HELP_MAKE_IT_BETTER: {
		"ru-RU": `Эта программа <b>бесплатна</b>. <a href="https://debtstracker.io/ru/help-us">Помогите</a> сделать её лучше!`,

		"en-US": `This program is <b>free to use</b>. Please <a href="https://debtstracker.io/en/help-us">help</a> to make it better!`,

		"it-IT": `Questo programma e' <b> completamente gratis</b>. Per favore <a href="https://debtstracker.io/">aiuta</a> a migliorarlo!`,
	},
	BUTTON_TEXT_YOU_OWE_AMOUNT_TO_SOMEONE: {
		"ru-RU": "%v | вы должны: %v",
		"en-US": "%v | you owe: %v",
		"it-IT": "%v | tu devi: %v",
	},
	BUTTON_TEXT_SOMEONE_OWES_TO_YOU_AMOUNT: {
		"ru-RU": "%v | долг вам: %v",
		"en-US": "%v | owes to you: %v",
		"it-IT": "%v | ti deve: %v",
	},
	BUTTON_TEXT_DEBT_RETURNED_FULLY: {
		"ru-RU": "Да, целиком",
		"en-US": "Yes, fully",
		"it-IT": "Si, completamente",
	},
	BUTTON_TEXT_DEBT_RETURNED_PARTIALLY: {
		"ru-RU": "Нет, только часть",
		"en-US": "No, just partially",
		"it-IT": "No, solo parzialmente",
	},
	MESSAGE_TEXT_ATTEMPT_TO_USE_OWN_INVITE: {
		"ru-RU": "Хорошая попытка пригласить самого себя ;)",
		"en-US": "You should not use your own invite ;)",
		"it-IT": "Non dovresti usare il tuo codice invito con te stesso :)",
	},
	MESSAGE_TEXT_WELCOME_ONBOARDING_INVITE_ACCEPTED: {
		"ru-RU": "Спасибо за то что воспользовались приглашением!",
		"en-US": "Welcome and thanks for accepting the invite!",
		"it-IT": "Benvenuto e grazie per aver accettato l'invito!",
	},
	MESSAGE_TEXT_FOR_COUNTERPARTY_ONLY: {
		"ru-RU": "Это действие доступно только для %v",
		"en-US": "This action for %v only.",
		"it-IT": "Questa azione e' disponibile solo per %v.",
	},
	BUTTON_TEXT_SEE_RECEIPT_DETAILS: {
		"ru-RU": "Показать детали",
		"en-US": "Show receipt details",
		"it-IT": "Mostra i dettagli del debito/credito",
	},
	MESSAGE_TEXT_YOU_SELECTED_INVITE_BY_EMAIL: {
		"ru-RU": "Вы решили пригласить друга через email.",
		"en-US": "You've selected to invite friend by email.",
		"it-IT": "Hai scelto di invitare l'amico tramite email.",
	},
	MESSAGE_TEXT_YOU_SELECTED_INVITE_BY_SMS: {
		"ru-RU": "Вы решили пригласить друга через SMS.",
		"en-US": "You've selected to invite friend by SMS.",
		"it-IT": "Hai scelto di invitare l'amico tramite SMS.",
	},
	MESSAGE_TEXT_ABOUT_INVITES: {
		"ru-RU": `На данный момент доступ к нашему боту ограничен, но вы можете пригласить друга.

Как вы хотите передать код приглашение?`,

		"en-US": `At the moment access to our bot is limited but you can invite your friend.

How do you want to pass the invite code?`,

		"it-IT": `AL momento l'accesso al nostro bot e' limitato ma puoi comunque invitare gli amici.

Come vuoi inviargli il codice invito?`,
	},
	MESSAGE_TEXT_USER_BLOCKED_TRANSFER_NOTIFICATIONS_BY: {
		"ru-RU": "%v заблокировал получение оповешений о транзакиях через: %v.",
		"en-US": "%v blocked reminders about transactions by: %v",
		"it-IT": "%v bloccato promemoria riguardo la transazione da: %v.",
	},
	COMMAND_TEXT_WAIT_A_SECOND: {
		"ru-RU": "Секундочку...",
		"en-US": "Wait a second...",
		"it-IT": "Solo un attimo...",
	},
	HTML_USING_TELEGRAM: {
		"ru-RU": "используя Telegram мессенджер",
		"en-US": "using Telegram messenger",
		"it-IT": "usa Telegram",
	},
	COMMAND_TEXT_ACCEPT: {
		"ru-RU": "Согласиться",
		"en-US": "Accept",
		"it-IT": "Accetta",
	},
	//BUTTON_TEXT_ACCEPT_TRANSFER_USING_TELEGRAM:{
	//	"ru-RU": "Подтвердить ",
	//	"en-US": "Accept ",
	//  "it-IT": "Accetta",
	//},
	//BUTTON_TEXT_DECLINE_TRANSFER_USING_TELEGRAM:{
	//	"ru-RU": "Отказаться (используя Telegram)",
	//	"en-US": "Decline (using Telegram messenger)",
	//  "it-IT": "Rifiuta (usando Telegram)",
	//},
	COMMAND_TEXT_DECLINE: {
		"ru-RU": "Отклонить",
		"en-US": "Decline",
		"it-IT": "Rifiuta",
	},
	COMMAND_TEXT_ACCEPT_INVITE: {
		"ru-RU": "Принять приглашение",
		"en-US": "Accept invite",
		"it-IT": "Accetta l'invito",
	},
	COMMAND_TEXT_VIEW_RECEIPT_DETAILS: {
		"ru-RU": "Посмотреть квитанцию",
		"en-US": "See receipt details",
		"it-IT": "Vedi dettagli",
	},
	COMMAND_TEXT_OTHER_WAYS_TO_SEND_INVITE: {
		"ru-RU": "Другие способы отправить приглашение",
		"en-US": "Other ways to send an invite",
		"it-IT": "Altri modi per inviare un invito",
	},
	COMMAND_TEXT_SEND_MY_PHONE_NUMBER: {
		"ru-RU": "Отправить мой номер",
		"en-US": "Send my phone number",
		"it-IT": "Invia il mio numero",
	},
	COMMAND_TEXT_SEND_BY_EMAIL: {
		"ru-RU": "Через Email",
		"en-US": "By Email",
		"it-IT": "Tramite email",
	},
	COMMAND_TEXT_SEND_BY_SMS: {
		"ru-RU": "Через SMS",
		"en-US": "By SMS",
		"it-IT": "Tramite SMS",
	},
	COMMAND_TEXT_INVITE_BY_TELEGRAM: {
		"ru-RU": "Пригласить через Telegram",
		"en-US": "Invite by Telegram",
		"it-IT": "Tramite Telegram",
	},
	MESSAGE_TEXT_INVITE_CREATED: {
		"ru-RU": `Мы отправили код приглашения на указынный вами адрес. (#%v)

Когда ваш друг потдвердит приглашение у вас будут общий баланс и транзакции (только между вами). Это поможет вам минимизировать усилия по ведению учёта.`,

		"en-US": `We've sent an invite code to your friend. (#%v)

Once your friend accepts invitation you'll share balance & transfers between you to make sure you both are on the same page with minimum efforts.`,

		"it-IT": `Abbiamo inviato il codice invito al tuo amico. (#%v)

Una volta che il tuo amico accetta l'invito potrete condividere i bilanci ed i trasferimenti con il minimo sforzo.`,
	},
	MESSAGE_TEXT_INVITE_BY_EMAIL: {
		"ru-RU": "Введите email вашего друга на который мы отправим код приглашения.",
		"en-US": "Please enter email address of your friend where we should send an invite code.",
		"it-IT": "Inserisci l'email dell'amico al quale inviare il codice invito.",
	},
	MESSAGE_TEXT_INVITE_ASK_EMAIL_FOR_RECEIPT: {
		"ru-RU": "Введите email вашего друга (%v) на который мы отправим квитанцию подтверждения.",
		"en-US": "Please enter emaill address of your friend (%v) where we should send the receipt.",
		"it-IT": "Inserisci l'email del tuo amico (%v) alla quale potremo inviare il debito/credito",
	},
	MESSAGE_TEXT_INVITE_BY_SMS: {
		"ru-RU": "Введите номер телефона вашего друга на который мы отправим код приглашения.",
		"en-US": "Please share a contact or enter phone number of your friend where we should send an invite code.",
		"it-IT": "COndividi il contatto o inserisci il numero di telefono del tuo amico al quale invieremo il codice invito.",
	},
	MESSAGE_TEXT_INVITE_BY_TELEGRAM: {
		"ru-RU": "Вставьте в чат контакт вашего друга которому вы хотите отправить приглашение.",
		"en-US": "Please share a contact of your friend you wish to send an invite code.",
		"it-IT": "Condividi il contatto di un amico al quale desideri inviare il codice invito.",
	},
	MESSAGE_TEXT_INVALID_EMAIL: {
		"ru-RU": "Неверный email. Проверьте и попробуйте ещё раз? /menu",
		"en-US": "Invalid email. Check and try it again? /menu",
		"it-IT": "Email scorretta. COntrolla e riprova. /menu",
	},
	MESSAGE_TEXT_INVALID_YEAR: {
		"ru-RU": "Неверный год. Год должен быть 2 или 4 цифры (<i>например 2016 или 16)</i>).",
		"en-US": "Invalid year. The year part should be 2 or 4 numbers (<i>e.g. 2016 or 16</i>).",
		"it-IT": "Anno scorretto. L'anno dev;essere composta da 2 o 4 numeri (<i>esempio 2017 oppure 17</i>)",
	},
	MESSAGE_TEXT_INVALID_MONTH: {
		"ru-RU": "Неверный месяц. Месяц должен быть задан целым числом от 1 до 12.",
		"en-US": "Invalid month. Month should be an integer from 1 to 12.",
		"it-IT": "Mese scorretto. Il mese dovrebbe essere un numero da 1 a 12.",
	},
	MESSAGE_TEXT_INVALID_DAY: {
		"ru-RU": "Неверный день. День должен быть задан целым числом от 1 до 31.",
		"en-US": "Invalid day. The day should be an integer from 1 to 31.",
		"it-IT": "Giorno scorretto. Il giorno dovrebbe essere un numero da 1 a 31.",
	},
	MESSAGE_TEXT_INVALID_DATE: {
		"ru-RU": "Неверный формат даты. Например для 20 февраля 2019 года надо ввести: 20.02.2019 или 20.02.19",
		"en-US": "Invalid date format. For exampel for 20 February 2019 please submit: 20.02.2019 or 20.02.19",
		"it-IT": "Formato data sbagliato. Esempio: per il 20 Febbraio 2019 inserisci: 20.02.2019 oppure 20.02.19",
	},
	MESSAGE_TEXT_INVALID_PHONE_NUMBER: {
		"ru-RU": "Неверный номер. Проверьте и попробуйте ещё раз? /menu",
		"en-US": "Invalid phone number. Check and try it again? /menu",
		"it-IT": "Numero di telefono invalido. Controlla e riprova. /menu",
	},
	MESSAGE_TEXT_PHONE_NUMBER_IS_NOT_SMS_CAPABLE: {
		"ru-RU": "Данный номер не принимает SMS. Попробуйте другой номер? /menu",
		"en-US": "This phone number not able to receive SMS. Try another number? /menu",
		"it-IT": "Questo numero di telefono non e' abilitato a ricevere SMS. Vuoi provare un altro numero? /menu",
	},
	MESSAGE_TEXT_NO_CONTACT_RECEIVED: {
		"ru-RU": "Мы не получили контакта. Тут инструкция как это сделать. /menu",
		"en-US": "We have not received any contacts. INSTRUCTION HOW TO DO IT. /menu",
		"it-IT": "Non abbiamo ricevuto nesusn contatto. ISTRUZIONI SU COME FARE. /menu",
	},
	MESSAGE_TEXT_CONTACT_NAME_IS_NUMBER: {
		"ru-RU": "Вы ввели только цифры в качестве имени контакта. Пожалуйста используйте текстовые символы.",
		"en-US": "You've entered just digits for a contact name. Please use some text characters.",
		"it-IT": "Hai inserito solamente numeri per un nome contatto. Usa anche alcune lettere.",
	},
	MESSAGE_TEXT_CURRENCY_NAME_IS_NUMBER: {
		"ru-RU": "Вы ввели только цифры в качестве номинала. Пожалуйста используйте текстовые символы.",
		"en-US": "You've entered just digits for currency. Please use some text characters.",
		"it-IT": "Hai inserito solamente numeri per la valuta. Usa anche alcune lettere.",
	},
	MESSAGE_TEXT_HISTORY_ROW_TO_USER_WITH_NAME: {
		"ru-RU": "%v - %s ⇒ Вам : %s",
		"en-US": "%v - %s ⇒ to you: %s",
		"it-IT": "%v - %s ⇒ a te: %s",
	},
	MESSAGE_TEXT_HISTORY_ROW_FROM_USER_WITH_NAME: {
		"ru-RU": "%v - Вы ⇒ %s : %s",
		"en-US": "%v - You ⇒ %s : %s",
		"it-IT": "%v - Tu ⇒ %s : %s",
	},
	MESSAGE_TEXT_LETS_SEND_SMS: {
		"ru-RU": "Давайте отправим SMS",
		"en-US": "Let's send SMS",
		"it-IT": "Inviamo un SMS",
	},
	MESSAGE_TEXT_SMS_QUEUING_FOR_SENDING: {
		"ru-RU": "SMS ставится в очередь на отправку на номер %v...",
		"en-US": "Queuing SMS for sending to number %v...",
		"it-IT": "SMS in coda per l'invio al numero %v...",
	},
	MESSAGE_TEXT_SMS_QUEUED_FOR_SENDING: {
		"ru-RU": "SMS поставлена в очередь на отправку на номер %v",
		"en-US": "SMS is queued for sending to number %v",
		"it-IT": "SMS inserito in coda per l'invio al numero %v",
	},
	MESSAGE_TEXT_BALANCE_HEADER: {
		"ru-RU": "Баланс",
		"en-US": "Balance",
		"it-IT": "Bilancio",
	},
	MESSAGE_TEXT_RECEIPT_AVAILABLE_CHANNELS: {
		"ru-RU": "Извините, в данный момент доступны только эти каналы для отправки квитанции:",
		"en-US": "Sorry, at the moment just this channels are available for sending a receipt:",
		"it-IT": "Spiacenti ma al momento solo questi canali sono disponibili per inviare debiti/crediti:",
	},
	MESSAGE_TEXT_RECEIPT_SENT_THROW_TELEGRAM: {
		"ru-RU": "Квитанция отправлена через телеграм.",
		"en-US": "Receipt sent through Telegram.",
		"it-IT": "Credito/debito inviato tramite Telegram",
	},
	MESSAGE_TEXT_RECEIPT_NOT_SENT_AS_COUNTERPARTY_HAS_DISABLED_TG_BOT: {
		"ru-RU": "Квитанция НЕ отправлена через телеграм так как %v удалил чат с ботом.",
		"en-US": "Receipt NOT sent through Telegram as %v has deleted chat with the bot.",
		"it-IT": "Credito/debito NON inviato tramite Telegram a %v perche' ha cancellato la chat con il bot",
	},
	MESSAGE_TEXT_RECEIPT_SENT_THROW_EMAIL: {
		"ru-RU": "Квитанция отправлена через email. (id: %v)",
		"en-US": "Receipt sent through email. (id: %v)",
		"it-IT": "Credito/debito inviato tramite email (id: %v)",
	},
	MESSAGE_TEXT_RECEIPT_SENT_THROW_SMS: {
		"ru-RU": "Квитанция отправлена через SMS.",
		"en-US": "Receipt sent through SMS.",
		"it-IT": "Credito/debito inviato trmaite SMS",
	},
	MESSAGE_TEXT_SWITCH_TO_PM_TO_VIEW_RECEIPT: {
		"ru-RU": "Переключитьсь на чат с ботом чтобы посмотреть квитанцию",
		"en-US": "Switch to private mode to see receipt details.",
		"it-IT": "Passa alla modalita' privata per vedere i dettagli dei tuoi crediti/debiti.",
	},
	MESSAGE_TEXT_RECEIPT_VIEWED_BY_COUNTERPARTY: {
		"ru-RU": "Квитанция просмотрена",
		"en-US": "Receipt viewed",
		//"it-IT": "Debiti visti",
	},
	MESSAGE_TEXT_VIEW_MY_NUMBER_IN_INTERNATIONAL_FORMAT: {
		"ru-RU": "Вы можете посмотреть свой номер телефона в ожидаемоем нами формате.",
		"en-US": "You can view your own phone number in the format we are expecting.",
		"it-IT": "Puoi visionare il tuo numero di telefono nel formato previsto.",
	},
	COMMAND_TEXT_VIEW_MY_NUMBER_IN_INTERNATIONAL_FORMAT: {
		"ru-RU": "Посмотреть мой номер в ожидаемоем формате",
		"en-US": "View my number in the expectd format",
		"it-IT": "Mostra il mio numero nel formato previsto",
	},
	INLINE_BUTTON_SHOW_FULL_HISTORY: {
		"ru-RU": "Показать всю историю",
		"en-US": "Show full history",
		"it-IT": "Mostra cronologia completa",
	},
	MESSAGE_TEXT_INCORRECT_VALUE_NOT_A_NUMBER: {
		"ru-RU": "Это не цифра",
		"en-US": "it is not a number",
		"it-IT": "Non e' un numero",
	},
	MESSAGE_TEXT_INCORRECT_VALUE_IS_NEGATIVE: {
		"ru-RU": "Цифра должна быть положительной (<i>больше нуля</i>)",
		"en-US": "The number should be positive (<i>greater than 0</i>)",
		"it-IT": "Il numero dovrebbe essere positivo (<i>maggiore di 0<i/>)",
	},
	MESSAGE_TEXT_ASK_HOW_MUCH_HAS_BEEN_RETURNED: {
		"ru-RU": "Сколько было возвращено?",
		"en-US": "How much have been returned?",
		"it-IT": "Quanti soldi ti son stati restituiti?",
	},
	MESSAGE_TEXT_HELP: {
		"ru-RU": "Вы можете сообщить о проблеме или сделать предложения по улучшению программы на нашем сайте.",
		"en-US": "Please report any issue or submit a feature request at our website.",
		"it-IT": "Segnala un problema o proponi una miglioria al nostro sito web.",
	},
	COMMAND_TEXT_OPEN_USER_REPORT: {
		"ru-RU": "Cтраница поддержки ",
		"en-US": "Support page",
		"it-IT": "Pagina d'aiuto",
	},
	COMMAND_TEXT_REPORT_A_BUG: {
		"ru-RU": "Сообщить об ошибке",
		"en-US": "Report a bug",
		"it-IT": "Segnala un bug",
	},
	COMMAND_TEXT_SUBMIT_AN_IDEA: {
		"ru-RU": "Предложить идею",
		"en-US": "Add an idea",
		"it-IT": "Proponi un idea",
	},
	MESSAGE_TEXT_WELCOME: {
		"ru-RU": `Привет, я Коллектиус - Ваш персональный счетовод и коллектор.

Могу записать кто кому чего должен и, и при необходимости, напомнить когда должок пора возвращать.

Чего изволит новый хозяин?`,
		"en-US": `Hi, I'm Collectius - your personal accountant & collector.

I can record who is owing to whom and remind when the return is due.

What would you like to do?`,

		"it-IT": `Ciao, sono Collectius - il tuo contabile ed esattore.

Posso annotare chi deve soldi a chi e ricordarti la data di scadenza.

Cosa vorresti fare ora?`,
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE: {
		"ru-RU": "Хочу получить приглашение",
		"en-US": "I want to get an invite",
		"it-IT": "Voglio ottenere un invito",
	},
	COMMAND_TEXT_I_HAVE_INVITE: {
		"ru-RU": "У меня есть код приглашения",
		"en-US": "I have the invitation code",
		"it-IT": "Ho il codice invito",
	},
	COMMAND_TEXT_I_HAVE_NOT_GOT_EMAIL: {
		"ru-RU": "Я не получил письма на email",
		"en-US": "I have not got any emails",
		"it-IT": "Non ho ricevuto nessun email",
	},
	MESSAGE_TEXT_ONBOARDING_TELL_ABOUT_INVITES: {
		"ru-RU": `<b>%v</b>,

На данный момент наш бот доступен только тем кто получил приглашение от друзей.

Если у вас нет кода можете оставить свои контакты и мы вышлем вам приглашение как только наступит ваша очередь.

Мы высылаем по 10 кодов в день первоочередникам + 1 случайным образом.`,

		"en-US": `<b>%v</b>,

At the moment our bot is available just by invitation from friends.

If you have no code you can leave your contact details and we'll send you an invite as soon as your queue is due.

We send 10 invites per day to those in the head of the queue and 1 randomly.`,

		"it-IT": `<b>%v</b>,

Al momento il nostro bot e' disponibile solo tramite invito da amici.

Se non hai un codice puoi lasciarci il tuo contatto e ti manderemo un invito non appena sara' il tuo turno.

Inviamo 10 inviti al giorno ai primi 10 della lista d'attesa ed 1 in modo casuale.`,
	},
	EMAIL_INVITE_SUBJ: {
		"ru-RU": "Приглашение от {{.FromName}} - код: {{.InviteCode}}",
		"en-US": "An invite from {{.FromName}} - code: {{.InviteCode}}",
		"it-IT": "Hai ricevuto un codice invito da {{.FromName}} - codice: {{.InviteCode}}",
	},
	SMS_INVITE_TEXT: {
		"ru-RU": `Привет {{.ToName}}, {{.FromName}} рекомендует приложение для учёта долгов - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Код приглашения: {{.InviteCode}}`,

		"en-US": `Hi {{.ToName}}, {{.FromName}} is inviting you to try debts tracking app - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Your personal invitation code is: {{.InviteCode}}`,

		"it-IT": `Ciao {{.ToName}}, {{.FromName}} ti ha invitato a provare 'debts tracking app' - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Il tuo codice di invito personale e': {{.InviteCode}}`,
	},
	EMAIL_INVITE_TEXT: {
		"ru-RU": `Привет {{.ToName}},

{{.FromName}} приглашает тебя попробовать приложение для учёта долгов - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Ваш код приглашения: {{.InviteCode}}`,

		"en-US": `Hi {{.ToName}},

{{.FromName}} is inviting you to use debts tracking app - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

You invitation code is: {{.InviteCode}}`,

		"it-IT": `Ciao {{.ToName}},

{{.FromName}} ti ha invitato a provare 'debts tracking app' - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Il tuo codice di invito personale e': {{.InviteCode}}`,
	},
	EMAIL_INVITE_HTML: {
		"ru-RU": `<p>Привет {{.ToName}},</P

<p>{{.FromName}} приглашает тебя <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">попробовать приложение для учёта долгов</a>.</p>

<p>Ваш код приглашения: <b>{{.InviteCode}}</b></p>`,

		"en-US": `<p>Hi {{.ToName}},</p>

<p>{{.FromName}} is inviting you to <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">try debts tracking app</a>.</p>

<p>Your invitation code is: <b>{{.InviteCode}}</b></p>`,

		"it-IT": `<p>Ciao {{.ToName}},</p>

<p>{{.FromName}} ti ha invitato a provare <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">debts tracking app</a>.</p>

<p>Il tuo codice di invito personale e': <b>{{.InviteCode}}</b></p>`,
	},
	EMAIL_RECEIPT_SUBJ: {
		"ru-RU": "Запись о долге - {{.FromName}}",
		"en-US": "Debt record - {{.FromName}}",
		"it-IT": "Debito - {{.FromName}}",
	},
	EMAIL_RECEIPT_BODY_TEXT: {
		"ru-RU": "{{.FromName}} создал(а) запись о долге: {{.ReceiptURL}}",
		"en-US": "{{.FromName}} created a debt record: {{.ReceiptURL}}",
		"it-IT": "{{.FromName}} ha creato un debito: {{.ReceiptURL}} ",
	},
	MESSENGER_RECEIPT_TEXT: {
		"ru-RU": "Я создал(а) запись о долге, подробности тут - {{.ReceiptURL}}",
		"en-US": "I've created a debt record regards you, see details at {{.ReceiptURL}}",
		"it-IT": "Ho creato un debito a tuo nome, controlla i dettagli qui - {{.ReceiptURL}}",
	},
	EMAIL_RECEIPT_BODY_HTML: {
		"ru-RU": "{{.FromName}} создал(а) запись о долге: {{.ReceiptURL}}",
		"en-US": "{{.FromName}} created a debt record: {{.ReceiptURL}}",
		"it-IT": "{{.FromName}} ha creato un debito: {{.ReceiptURL}}",
	},
	INLINE_RECEIPT_TITLE: {
		"ru-RU": "Квитанция: %v",
		"en-US": "Receipt: %v",
		"it-IT": "Debito/credito: %v",
	},
	INLINE_RECEIPT_DESCRIPTION: {
		"ru-RU": "Нажмите здесь чтобы отправить квитанцию",
		"en-US": "Click here to send the receipt",
		"it-IT": "Clicca qui per inviare un debito/credito",
	},
	INLINE_RECEIPT_CHOOSE_LANGUAGE: {
		"ru-RU": "<b>Выберите язык чтобы посмотреть подробности записи о долге</b> которую создал(а) {{.Creator}}.",
		"en-US": "<b>Please choose language to see details of the debt</b> that has been recorded by {{.Creator}}.",
		"it-IT": "<b>Scegli la lingua per vedere i dettagli del debito</b> registrato da {{.Creator}}.",
	},
	INLINE_RECEIPT_MESSAGE: {
		"ru-RU": `<b>{{.Creator}} создал(а) запись о долге</b> касающегося Вас.

{{.SiteLink}} — программа для учёта долгов поможет:

  - Всегда знать кто кому сколько должен

  - Незабыть вовремя отдать или востребовать долг
    <i>(напоминания вам и вашим должникам)</i>`,
		//-------------------------------------------------------
		"en-US": `<b>{{.Creator}} recorded a debt</b> associated with you.

{{.SiteLink}} — an app for debts tracking will help you to:

  - Always know your bottom line

  - Return debts on time
    <i>(reminders to you & your debtors)</i>`,
		//-------------------------------------------------------
		"it-IT": `<b>{{.Creator}} ha registrato un debito</b> associato a te.

{{.SiteLink}} — un app per i debiti che ti consento di:

  - Sapere sempre chi deve soldi a chi

  - Restituire i soldi in tempo
    <i>(lo ricorda a te ed al tuo debitore)</i>`,
	},
	INLINE_INVITE_TITLE: {
		"ru-RU": "Приглашение в %v",
		"en-US": "Invitation to %v",
		"it-IT": "Invito per %v",
	},
	INLINE_INVITE_DESCRIPTION: {
		"ru-RU": "Нажмите здесь для отправки приглашения",
		"en-US": "Click here to send an invite",
		"it-IT": "Clicca qui per spedire un invito",
	},
	INLINE_INVITE_MESSAGE: {
		"ru-RU": "%v пригласил вас попробовать %v",
		"en-US": "%v invited you to try %v",
		"it-IT": "%v ti ha invitato a provare %v",
	},
	SMS_RECEIPT_YOU_GOT: {
		"ru-RU": "Вы получили %v от %v.",
		"en-US": "You've got %v from %v.",
		"it-IT": "Hai ricevuto %v da %v",
	},
	SMS_RECEIPT_YOU_GAVE: {
		"ru-RU": "Вы дали %v - взял(а) %v.",
		"en-US": "You've given %v to %v.",
		"it-IT": "Hai dato %v a %v",
	},
	SMS_CLICK_TO_CONFIRM_OR_DECLINE: {
		"ru-RU": "Нажмите %v чтобы подтвердить или отвергнуть.",
		"en-US": "Click %v to confirm or decline.",
		"it-IT": "Clicca %v per confermare o rifiutare",
	},
	HTML_DATE: {
		"ru-RU": "Дата",
		"en-US": "Date",
		"it-IT": "Data",
	},
	HTML_RECEIPT: {
		"ru-RU": "Квитанция",
		"en-US": "Receipt",
		"it-IT": "Scontrino", //To upgrade, not the best translation from Russian
	},
	HTML_AMOUNT: {
		"ru-RU": "Сумма",
		"en-US": "Amount",
		"it-IT": "Totale",
	},
	HTML_FROM: {
		"ru-RU": "Дал",
		"en-US": "From",
		"it-IT": "Da",
	},
	HTML_TO: {
		"ru-RU": "Получил",
		"en-US": "To",
		"it-IT": "Per",
	},
	TELEGRAM_RECEIPT: {
		"ru-RU": "{{.FromName}} создал запись о долге ({{.TransferCurrency}})",
		"en-US": "{{.FromName}} created a debtrecord ({{.TransferCurrency}})",
		"it-IT": "{{.FromName}} ha registrato un debito ({{.TransferCurrency}})",
	},
	MESSAGE_TEXT_PLEASE_CHOOSE_FROM_OPTIONS_PROVIDED: {
		"ru-RU": "Пожалуйста выберете из предоставленных опций.",
		"en-US": "Please choose from provided options.",
		"it-IT": "Scegli tra le opzioni fornite.",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_NOTE_OR_COMMENT: {
		"ru-RU": "<b>Хотите добавить заметку или комментарий?</b>\n%v Заметки хранятся для вашего личго пользования.\n%v Комментарий виден всем кому разрешён просмотр этой транзакции.",
		"en-US": "<b>Do you want to add a note or comment?</b>\n%v Memos are private records for yoru own reference.\n%v Comments are available to everyone who has permission to view this transaction.",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_NOTE: {
		"ru-RU": "Напишите заметку:",
		"en-US": "Please write a note:",
		"it-IT": "Per favore scrivi un appunto:",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_COMMENT_ONLY: {
		"ru-RU": `Если хотите добавить комментарий просто отправьте текст.`,
		"en-US": `If you want to add a comment just send a text now.`,
		"it-IT": `Se vuoi aggiungere un commento invia del testo ora.`,
	},
	MESSAGE_TEXT_VISIBLE_TO_YOU_AND_COUNTERPARTY: {
		"ru-RU": "виден вам и %v",
		"en-US": "visible to you & %v",
		"it-IT": "visibile a te e %v",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_COMMENT: {
		"ru-RU": "Напишите комментарий:",
		"en-US": "Please write the comment:",
		"it-IT": "Per favore scrivi un commento:",
	},
	MESSAGE_TEXT_TRANSFER_NOTE_ADDED_ASK_FOR_COMMENT: {
		"ru-RU": "Заметка добавлена. Хотите написать комментарий?",
		"en-US": "Memo have been added. Do you want to write a comment?",
		"it-IT": "Promemoria aggiunto. Vuoi scrivere un commento?",
	},
	MESSAGE_TEXT_TRANSFER_COMMENT_ADDED_ASK_FOR_NOTE: {
		"ru-RU": "Комментарий добавлен. Хотите написать заметку?",
		"en-US": "Comment have been added. Do you want to write a note?",
		"it-IT": "Commento aggiunto. Vuoi scrivere un appunto?",
	},
	COMMAND_TEXT_NO_COMMENT_OR_NOTE_FOR_TRANSFER: {
		"ru-RU": "Без заметок и комментариев",
		"en-US": "Without notes or comments",
		"it-IT": "Senza appunti o commenti",
	},
	COMMAND_TEXT_NO_COMMENT_FOR_TRANSFER: {
		"ru-RU": "Без комментариев",
		"en-US": "No comments",
		"it-IT": "No commenti",
	},
	COMMAND_TEXT_NO_NOTE_FOR_TRANSFER: {
		"ru-RU": "Без заметок",
		"en-US": "Without notes",
		"it-IT": "Senza appunti/note",
	},
	COMMAND_TEXT_ADD_NOTE_TO_TRANSFER: {
		"ru-RU": "Добавить заметку",
		"en-US": "Add a note (private)",
		"it-IT": "Aggiungi una nota (privata)",
	},
	COMMAND_TEXT_ADD_COMMENT_TO_TRANSFER: {
		"ru-RU": "Добавить комментарий",
		"en-US": "Add a comment (public)",
		"it-IT": "Aggiungi un commento (pubblico)",
	},
	DUE_IN_NOW: {
		"ru-RU": "прямо сейчас",
		"en-US": "now",
		"it-IT": "adesso",
	},
	DUE_IN_A_MINUTE: {
		"ru-RU": "через минуту",
		"en-US": "in a minute",
		"it-IT": "in un minuto",
	},
	DUE_IN_X_MINUTES: {
		"ru-RU": "через %v минут(ы)",
		"en-US": "in %v minutes",
		"it-IT": "in %v minuti/o",
	},
	DUE_IN_AN_HOUR: {
		"ru-RU": "через час",
		"en-US": "in an hour",
		"it-IT": "in un ora",
	},
	DUE_IN_X_HOURS: {
		"ru-RU": "через %v часа/часов",
		"en-US": "in %v hours",
		"it-IT": "in %v ore/a",
	},
	DUE_IN_X_DAYS: {
		"ru-RU": "через %v дня/дней",
		"en-US": "in %v days",
		"it-IT": "in %v giorni/o",
	},
	//-------------------------------------------------------------------------------------------------------------------
	WS_ALEX_T: {
		"ru-RU": "Александр Трахимёнок",
		"en-US": "Alexander Trakhimenok",
		"it-IT": "Alexander Trakhimenok",
	},

	WS_INDEX_TITLE: {
		"ru-RU": "DebtsTracker.io - программа для учёта личных долгов и активов",
		"en-US": "DebtsTracker.io - an IOU app to track your personal debts & assets",
		"es-ES": "DebtsTracker.io - una aplicación para el seguimiento de sus deudas personales",
		"fa-IR": "DebtsTracker.io - برنامه برای پیگیری بدهی های شخصی خود را",
		"pl-PL": "DebtsTracker.io - aplikacja do śledzenia osobistych długów",
		"pt-PT": "DebtsTracker.io - um aplicativo para controlar suas dívidas pessoais",
		"de-DE": "DebtsTracker.io - eine App, um Ihre persönlichen Schulden zu verfolgen",
		"fr-FR": "DebtsTracker.io - une application pour suivre vos dettes personnelles",
		"it-IT": "DebtsTracker.io - un app per monitorare i tuoi debiti personali",
		"ko-KO": "DebtsTracker.io 은 - 앱이 사용자의 개인 채무를 추적",
		"ja-JP": "DebtsTracker.io は - アプリはあなたの個人的な借金を追跡します",
		"zh-CN": "DebtsTracker.io - 一个应用程序来跟踪你的个人债务",
	},
	WS_LIVE_DEMO: {
		"ru-RU": "Демо версия online",
		"en-US": "Live demo",
		"es-ES": "Demo en vivo",
		"fa-IR": "نسخه ی نمایشی زنده",
		"pl-PL": "Demo na żywo",
		"pt-PT": "Demonstração ao vivo",
		"de-DE": "Echtzeit Vorführung",
		"fr-FR": "Démo en direct",
		"it-IT": "Demo online",
		"ko-KO": "실시간 데모",
		"ja-JP": "ライブデモ",
		"zh-CN": "现场演示",
	},
	WS_INDEX_TG_BOT_H2: {
		"ru-RU": "Бот для Telegram",
		"en-US": "Chat bot for Telegram messenger",
		"es-ES": "Chat bot para Telegrama mensajero",
		"fa-IR": "ربات چت برای رسول تلگرام",
		"pl-PL": "Chat bot do telegramu posłańca",
		"pt-PT": "bot de bate-papo para Telegram messenger",
		"de-DE": "Chat-Bot für Telegrammbote",
		"fr-FR": "bot Chat for Telegram messenger",
		"it-IT": "Bot Chat per Telegram",
		"ko-KO": "전보 메신저 채팅 봇",
		"ja-JP": "電報メッセンジャーのためのチャットボット",
		"zh-CN": "聊天机器人的电报使者",
	},
	WS_INDEX_TG_BOT_OPEN: {
		"ru-RU": "Открыть в Телеграмме &#x1F680;",
		"en-US": "Open in Telegram &#x1F680;",
		"es-ES": "Abrir en Telegrama &#x1F680;",
		"fa-IR": "باز کردن در تلگرام &#x1F680;",
		"pl-PL": "Otwórz w telegramu &#x1F680;",
		"pt-PT": "Open in Telegram &#x1F680;",
		"de-DE": "Open in Telegramm &#x1F680;",
		"fr-FR": "Open in Telegram &#x1F680;",
		"it-IT": "Apri su Telegram &#x1F680;",
		"ko-KO": "전보 &#x1F680; 에서 열기;",
		"ja-JP": "電報 &#x1F680; で開きます。",
		"zh-CN": "打开在电报 &#x1F680;",
	},

	WS_INDEX_TG_BOT_P: {
		"ru-RU": "В настоящий момент наша программа доступна в мессенджере <a href='https://telegram.org/'>Телеграм</a>.",
		"en-US": "At the moment our program is available just on <a href='https://telegram.org/'><b>Telegram</b> messenger</a>",
		"es-ES": "Por el momento nuestro programa está disponible sólo en <a href='https://telegram.org/'> <b> Telegrama </b> mensajero </a>",
		"fa-IR": "در حال حاضر برنامه های ما فقط در دسترس است <a href='https://telegram.org/'>Telegram</a>",
		"pl-PL": "W tej chwili nasz program jest dostępny tylko na <a href='https://telegram.org/'> <b> Telegram </b> messenger </a>",
		"pt-PT": "No momento em que o nosso programa está disponível apenas na <a href='https://telegram.org/'> <b> Telegram </b> messenger </a>",
		"de-DE": "Im Moment ist unser Programm nur auf verfügbar <a href='https://telegram.org/'> <b> Telegramm </b> Bote </a>",
		"fr-FR": "Au moment de notre programme est disponible seulement sur <a href='https://telegram.org/'> <b> Telegram </b> messager </a>",
		"it-IT": "Al momento il nostro programma è disponibile solo su <a href='https://telegram.org/'> <b> Telegram </b></a>",
		"ko-KO": "지금이 순간 우리의 프로그램은 단지에 <a href='https://telegram.org/'>의 <b> 전보 </b>을 메신저 </a>를 볼 수 있습니다",
		"ja-JP": "現時点では私たちのプログラムは、ちょうど上の<a href='https://telegram.org/'><B>電報</b>のメッセンジャー</a>で提供されています",
		"zh-CN": "目前我们的计划是只提供在<a href='https://telegram.org/'><B>电报</b>的使者</A>",
	},
	WS_MOTTO: {
		"ru-RU": "Платежи по долгам целиком и вовремя!",
		"en-US": "Know your bottom line & never miss a debt payment!",
		"es-ES": "Conozca a su línea de fondo y nunca se pierda un pago de la deuda!",
		"fa-IR": "از موجودی حساب خود و هرگز پرداخت بدهی از دست ندهید!",
		"pl-PL": "Znaj swoją równowagę i nigdy nie przegapisz zapłatę długu!",
		"pt-PT": "Conheça o seu equilíbrio e nunca perca um pagamento da dívida!",
		"de-DE": "Ihr Kontostand wissen und nie eine Schuld Zahlung verpassen!",
		"fr-FR": "Apprenez à connaître votre solde et ne jamais manquer un paiement de la dette!",
		"it-IT": "Tieni sott'occhio il tuo bilancio e non dimenticarti mai di un debito!",
		"ko-KO": "균형을 알고 및 채무 지불을 놓칠 수 없어!",
		"ja-JP": "あなたのバランスを知っている＆債務の支払いを見逃すことはありません！",
		"zh-CN": "了解天平＆不会错过任何一个债务付款！",
	},
	WS_SHORT_DESC: {
		"ru-RU": "DebtsTracker.io - мобильное приложение и сервис напоминаний для учёта и своевременной выплаты долгов. Отсылает автоматические уведомления вашим должникам по email и SMS.",
		"en-US": "DebtsTracker.io is a mobile IOU app & a reminder service that helps to track your debts, credits & assets. Sends automated email & SMS reminders to your debtors.",
		"es-ES": "DebtsTracker.io es un servicio de aplicaciones móviles y recordatorio de que ayuda a realizar un seguimiento de sus deudas y créditos. Envía notificaciones por correo electrónico y SMS automatizados a sus deudores.",
		"fa-IR": "DebtsTracker.io یک سرویس برنامه های تلفن همراه و یادآوری است که کمک می کند تا برای پیگیری بدهی و اعتبار خود را است. ارسال ایمیل و اس ام اس اطلاعیه های خودکار به بدهکاران خود را.",
		"pl-PL": "DebtsTracker.io to aplikacje mobilne i przypomnienia usługa, która pozwala na śledzenie swoich długów i kredytów. Wysyła automatycznych powiadomień e-mail i SMS do swoich dłużników.",
		"pt-PT": "DebtsTracker.io é um serviço de aplicativos móveis e lembrete de que ajuda a controlar seus débitos e créditos. Envia e-mail e SMS notificações automáticas aos seus devedores.",
		"de-DE": "DebtsTracker.io ist eine mobile Apps und Erinnerungs-Service, die Ihre Schulden und Kredite zu verfolgen hilft. Sendet automatisierte E-Mail und SMS-Benachrichtigungen an Ihre Schuldner.",
		"fr-FR": "DebtsTracker.io est une des applications mobiles et rappel service qui permet de suivre vos dettes et crédits. Envoie automatisés email & SMS reminders à vos débiteurs.",
		"it-IT": "DebtsTracker.io è un servizio di applicazioni mobili che ricordare e aiuta a monitorare i debiti e crediti. Invia notifiche e-mail e SMS automatici ai i vostri debitori.",
		"ko-KO": "DebtsTracker.io 은 채무 및 크레딧을 추적하는 데 도움이 모바일 앱 및 알림 서비스입니다. 당신의 채무자에 자동화 된 이메일 및 SMS 알림을 보냅니다.",
		"ja-JP": "DebtsTracker.io は、あなたの借金＆クレジットを追跡するのに役立ちますモバイルアプリ＆リマインダーサービスです。あなたの債務者に自動メール＆SMS通知を送信します。",
		"zh-CN": "DebtsTracker.io 是一个移动应用和提醒服务，帮助跟踪你的债务和信用。发送自动电子邮件和短信通知到您的债务人。",
	},

	WS_HELP_US_TITLE: {
		"en-US": "How you can help to DebtsTracker.io project",
		"ru-RU": "Как вы можете помочь проекту DebtsTracker.io",
	},
	WS_ADS_TITLE: {
		"en-US": "Ads @ DebtsTracker.IO",
		"ru-RU": "Реклама на DebtsTracker.IO",
		"it-IT": "Pubblicita' @ DebtsTracker.IO",
	},
	WS_ADS_CONTENT: {
		"en-US": `To place ads in our app please send us an email to <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
		"ru-RU": `Чтобы разместить рекламу в нашем приложении пишите нам на <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
		"it-IT": `Per inserire della pubblicita nella nostra app inviaci un email a <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
	},
	FB_MAKE_RECORD_HEADER: {
		"en-US": "Record debt",
		"ru-RU": "Записать долг",
		"it-IT": "Registra il debito",
	},
	FB_MAKE_RECORD_HEADLINE: {
		"en-US": "Scroll left to see other options.",
		"ru-RU": "Пролистайте карточки влево чтобы увидеть другие опции.",
		"it-IT": "Scrolla a sinistra per vedere altre opzioni",
	},
	FB_HOW_ARE_THINGS_HEADER: {
		"en-US": "How are you doing?",
		"ru-RU": "Как идут дела?",
		"it-IT": "Come te la passi?",
	},
}
