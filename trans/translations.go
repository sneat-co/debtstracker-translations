package trans

import (
	"github.com/DebtsTracker/translations/emoji"
)

const adsCommandTitle = "\xE2\xAD\x90\xE2\xAD\x90\xE2\xAD\x90"

// TRANS - translation string
var TRANS = map[string]map[string]string{
	"EXAMPLE": {
		"ru-RU": "ПРИМЕР",
		"en-US": "SAMPLE",
	},
	"Jan": {
		"en-US": "Jan",
		"ru-RU": "Янв.",
	},
	"Feb": {
		"en-US": "Feb",
		"ru-RU": "Фев.",
	},
	"Mar": {
		"en-US": "Mar",
		"ru-RU": "Мрт.",
	},
	"Apr": {
		"en-US": "Apr",
		"ru-RU": "Апр.",
	},
	"May": {
		"en-US": "May",
		"ru-RU": "Май ",
	},
	"Jun": {
		"en-US": "Jun",
		"ru-RU": "Июнь",
	},
	"Jul": {
		"en-US": "Jul",
		"ru-RU": "Июль",
	},
	"Aug": {
		"en-US": "Aug",
		"ru-RU": "Авг.",
	},
	"Sep": {
		"en-US": "Sep",
		"ru-RU": "Сен.",
	},
	"Oct": {
		"en-US": "Oct",
		"ru-RU": "Окт.",
	},
	"Nov": {
		"en-US": "Nov",
		"ru-RU": "Нбр.",
	},
	"Dec": {
		"en-US": "Dec",
		"ru-RU": "Дек.",
	},

	adsCommandTitle: {
		"ru-RU": adsCommandTitle,
		"en-US": adsCommandTitle,
	},
	" and ": {
		"ru-RU": " и ",
		"en-US": " and ",
	},
	"MESSAGE_TEXT_OOPS_SOMETHING_WENT_WRONG": {
		"ru-RU": "Упс, что-то пошло не так... \xF0\x9F\x98\xB3",
		"en-US": "Oops, something went wrong... \xF0\x9F\x98\xB3",
	},
	MESSAGE_TEXT_ASK_DUE: {
		"ru-RU": "Когда дата возврата?",
		"en-US": "When is the due date?",
	},
	MESSAGE_TEXT_ASK_DATE_TO_REMIND: {
		"ru-RU": `Чтобы задать дату напопинания напишите её в формате <i>ДД.MM.ГОД</i>.
<b>Например</b> для 20 января 2017 г. отправьте:
    <i>20.01.2017</i>`,

		"en-US": `To set date for next reminder please send it as a text in format of <i>DD.MM.YEAR</i>.
<b>For example</b> for 20th of January 2017 submit:
    <i>20.01.2017</i>`,
	},
	MESSAGE_TEXT_ASK_DUE_DATE: {
		"ru-RU": `Чтобы задать дату возврата напишите её в формате <i>ДД.MM.ГОД</i>.
<b>Например</b> для 20 января 2017 г. отправьте:
    <i>20.01.2017</i>`,

		"en-US": `To set due date please send it as a text in format of <i>DD.MM.YEAR</i>.
<b>For example</b> for 20th of January 2017 submit:
    <i>20.01.2017</i>`,
	},
	MESSAGE_TEXT_WRONG_DATE: {
		"ru-RU": "Извините, что-то не так с датой которую вы отправили.",
		"en-US": "Sorry, there is something wrong with the date you've provided.",
	},
	COMMAND_TEXT_DISABLE_REMINDER: {
		"ru-RU": "Не напоминать",
		"en-US": "No reminder",
	},
	COMMAND_TEXT_TOMORROW: {
		"ru-RU": "Завтра",
		"en-US": "Tomorrow",
	},
	COMMAND_TEXT_DAY_AFTER_TOMORROW: {
		"ru-RU": "Послезавтра",
		"en-US": "Day after tomorrow",
	},
	COMMAND_TEXT_THIS_WEEK: {
		"ru-RU": "На этой неделе",
		"en-US": "This week",
	},
	COMMAND_TEXT_YES_IT_HAS_RETURN_DEADLINE: {
		"ru-RU": "Да, есть срок возврата!",
		"en-US": "Yes, it has a deadline!",
	},
	COMMAND_TEXT_NO_IT_CAN_BE_RETURNED_ANYTIME: {
		"ru-RU": "Нет, срок возврата не важен.",
		"en-US": "No, whenever is fine.",
	},
	COMMAND_TEXT_IT_CAN_BE_RETURNED_ANYTIME: {
		"ru-RU": "Когда-нибудь",
		"en-US": "Whenever is fine",
	},
	COMMAND_TEXT_IN_FEW_MINUTES: {
		"ru-RU": "Через минуту",
		"en-US": "In few minutes",
	},
	COMMAND_TEXT_IN_1_WEEK: {
		"ru-RU": "Через неделю",
		"en-US": "In 1 week",
	},
	COMMAND_TEXT_IN_1_MONTH: {
		"ru-RU": "Через месяц",
		"en-US": "In 1 month",
	},
	COMMAND_TEXT_SET_DATE: {
		"ru-RU": "Задать дату",
		"en-US": "Set date",
	},
	COMMAND_TEXT_MONDAY: {
		"ru-RU": "Понедельник",
		"en-US": "Monday",
	},
	COMMAND_TEXT_TUESDAY: {
		"ru-RU": "Вторник",
		"en-US": "Tuesday",
	},
	COMMAND_TEXT_WEDNESDAY: {
		"ru-RU": "Среда",
		"en-US": "Wednesday",
	},
	COMMAND_TEXT_THURSDAY: {
		"ru-RU": "Четверг",
		"en-US": "Thursday",
	},
	COMMAND_TEXT_FRIDAY: {
		"ru-RU": "Пятница",
		"en-US": "Friday",
	},
	COMMAND_TEXT_SATURDAY: {
		"ru-RU": "Суббота",
		"en-US": "Saturday",
	},
	COMMAND_TEXT_SUNDAY: {
		"ru-RU": "Воскресенье",
		"en-US": "Sunday",
	},
	COMMAND_TEXT_DO_NOT_SEND_RECEIPT: {
		"ru-RU": "Не отправлять квитанцию",
		"en-US": "Do not send the receipt",
	},
	MESSAGE_TEXT_RECEIPT_WILL_NOT_BE_SENT: {
		"ru-RU": "Вы решили не отправлять квитанцию.",
		"en-US": "You've decided not to send the receipt.",
	},
	COMMAND_TEXT_I_HAVE_CHANGED_MY_MIND: {
		"ru-RU": "Я передумал(а)",
		"en-US": "I've changed my mind",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_TELEGRAM: {
		"ru-RU": "Отправить через Telelgram",
		"en-US": "Send by Telegram",
	},
	COMMAND_TEXT_COUNTERPARTY_HAS_NO_TELEGRAM: {
		"ru-RU": "Отправить через Viber, VK, FB, ...",
		"en-US": "Send by FB, WhatsApp, Viber, etc.",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_SMS: {
		"ru-RU": "Отправить через SMS",
		"en-US": "Send by SMS",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_VK: {
		"ru-RU": "Отправить через ВКонтакте",
		"en-US": "Send throw VK.com",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_OK: {
		"ru-RU": "Отправить через Одноклассники",
		"en-US": "Send throw OK",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_FB: {
		"ru-RU": "Отправить через Facebook",
		"en-US": "Send throw Facebook",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_TWT: {
		"ru-RU": "Отправить через Twitter",
		"en-US": "Send throw Twiter",
	},
	COMMAND_TEXT_CANCEL_SENDING_RECEIPT_BY_TELEGRAM: {
		"ru-RU": "Отменить отправку через Telelgram",
		"en-US": "Cancel sending receipt by Telegram",
	},
	COMMAND_TEXT_MAIN_MENU_TITLE: {
		"ru-RU": "Главное /меню",
		"en-US": "Main /menu",
	},
	MESSAGE_TEXT_NOTHING_TO_CANCEL: {
		"ru-RU": "Нечего отменять.",
		"en-US": "Nothing to cancel.",
	},
	MESSAGE_TEXT_TRANSFER_CREATION_CANCELED: {
		"ru-RU": "Создание записи о долге отменено.",
		"en-US": "Creation of debt record has been canceled.",
	},
	COMMAND_TEXT_SHOW_ALL_CONTACTS: {
		"ru-RU": "Показать все...",
		"en-US": "Show all...",
	},
	COMMAND_TEXT_ADD_YOUR_OWN_OPTION: {
		"ru-RU": "Что-то другое",
		"en-US": "Something else",
	},
	MESSAGE_TEXT_REMINDER_ASK_IF_RETURNED: {
		"ru-RU": "Был ли этот долг возвращён?",
		"en-US": "Have this debt been returned?",
	},
	MESSAGE_TEXT_ASK_WHEN_TO_REMIND_AGAIN: {
		"ru-RU": "Когда вам напомнить об этом долге ещё раз?",
		"en-US": "When should we remind you about this debt again?",
	},
	MESSAGE_TEXT_REPLIED_DEBT_RETURNED_FULLY: {
		"ru-RU": "Вы ответили что долг возвращён полностью.",
		"en-US": "You've replied that debt has been returned fully.",
	},
	MESSAGE_TEXT_DEBT_IS_RETURNED: {
		"ru-RU": "Долг возвращён полностью.",
		"en-US": "The debt has been returned fully.",
	},
	MESSAGE_TEXT_DETAILS_ARE_HERE: {
		"ru-RU": "Подробности тут: %v",
		"en-US": "Details here: %v",
	},
	MESSAGE_TEXT_REMINDER: {
		"ru-RU": "Напоминание",
		"en-US": "Reminder",
	},
	MESSAGE_TEXT_REMINDER_SET: {
		"ru-RU": "Напоминание установлено на: %v",
		"en-US": "Reminder set for: %v",
	},
	MESSAGE_TEXT_REMINDER_DISABLED: {
		"ru-RU": "Напоминания об этом долге отключены.",
		"en-US": "You've disabled reminders for this debt.",
	},
	MESSAGE_TEXT_REMINDER_ALREADY_RESCHEDULED: {
		"ru-RU": "Напоминание об этом долге уже перенесено.",
		"en-US": "You've already rescheduled this reminder.",
	},
	COMMAND_TEXT_REMINDER_RETURNED_IN_FULL: {
		"ru-RU": "Да, возвращено полностью",
		"en-US": "Yes, returne in full",
	},
	COMMAND_TEXT_REMINDER_RETURNED_PARTIALLY: {
		"ru-RU": "Возврашено частично",
		"en-US": "Returned partially",
	},
	COMMAND_TEXT_REMINDER_NOT_RETURNED: {
		"ru-RU": "Не возвращено",
		"en-US": "Not returned",
	},
	MESSAGE_TEXT_YOU_REPLIED: {
		"ru-RU": "Вы ответили: %v",
		"en-US": "You've replied: %v",
	},
	"book": {
		"ru-RU": "книгу",
		"en-US": "book",
	},
	"MESSAGE_TEXT_I_DID_NOT_UNDERSTAND_THE_COMMAND": {
		"ru-RU": "\xF0\x9F\x98\xB3 Извините, я не понял вашу команду. Возможно я немного туповат...\n\nЧтобы начать сначала нажмите /menu",
		"en-US": "\xF0\x9F\x98\xB3 Sorry, I did not understand your order. May be I'm a little bit dumb...\n\nYou can return to main /menu",
	},
	"COMMAND_TEXT_LANGUAGE": {
		"ru-RU": "/Язык приложения",
		"en-US": "App /language",
	},
	"/start": {
		"ru-RU": "/старт",
		"en-US": "/start",
	},
	COMMAND_TEXT_DUE_RETURNS: {
		"ru-RU": "Предстоящие платежи",
		"en-US": "Due returns",
	},
	MESSAGE_TEXT_OVERDUE_RETURNS_HEADER: {
		"ru-RU": "<b>Просроченные долги:</b>",
		"en-US": "<b>Overdue debts:</b>",
	},
	MESSAGE_TEXT_DUE_RETURNS_HEADER: {
		"ru-RU": "<b>Ближайшие долги к возрату:</b>",
		"en-US": "<b>Closest debts to return:</b>",
	},
	MESSAGE_TEXT_DUE_RETURNS_ROW_BY_USER: {
		"ru-RU": "%v ожидает от вас возврата %v через %v",
		"en-US": "%v expects %v from you in %v",
	},
	MESSAGE_TEXT_DUE_RETURNS_ROW_BY_COUNTERPARTY: {
		"ru-RU": "Вы ожидаете от %v возврата %v через %v",
		"en-US": "You expect %v to retun %v to you in %v",
	},

	MESSAGE_TEXT_DUE_RETURNS_EMPTY: {
		"ru-RU": "У вас нет долгов с назначеным сроком к возврату.",
		"en-US": "You have no debts with set due date.",
	},
	COMMAND_TEXT_GAVE: {
		"ru-RU": "/Дал",
		"en-US": "/Gave",
	},
	COMMAND_TEXT_GOT: {
		"ru-RU": "/Взял",
		"en-US": "/Got",
	},
	COMMAND_TEXT_RETURN: {
		"ru-RU": "/Вернул",
		"en-US": "/Returned",
	},
	COMMAND_TEXT_BALANCE: {
		"ru-RU": "/Баланс",
		"en-US": "/Balance",
	},
	COMMAND_TEXT_SETTING: {
		"ru-RU": "/Настройки",
		"en-US": "/Settings",
	},
	COMMAND_TEXT_HIGH_FIVE: {
		"ru-RU": "Дать пять!",
		"en-US": "High five!",
	},
	COMMAND_TEXT_CHANGE_LANG: {
		"ru-RU": "/Язык",
		"en-US": "/Language",
	},
	COMMAND_TEXT_HELP: {
		"ru-RU": "/Помощь",
		"en-US": "/Help",
	},
	COMMAND_TEXT_HISTORY: {
		"ru-RU": "/История",
		"en-US": "/History",
	},
	COMMAND_TEXT_CANCEL: {
		"ru-RU": "/Отменить",
		"en-US": "/Cancel",
	},
	COMMAND_TEXT_SETTINGS_PRIMARY_CURRENCY: {
		"ru-RU": "Основная валюта",
		"en-US": "Primary currency",
	},
	COMMAND_TEXT_NEW_COUNTERPARTY: {
		"ru-RU": "Добавить",
		"en-US": "Add new",
	},
	MESSAGE_TEXT_LOGIN_CODE: {
		"ru-RU": "Ваш код для входа в приложение: <b>%v</b>",
		"en-US": "Your code for signing in to app: <b>%v</b>",
	},
	MESSAGE_TEXT_ASK_NEW_COUNTERPARTY_NAME: {
		"ru-RU": `<b>Имя для нового контакта</b>
Напишите сами или выберите из своей адресной книги (<i>через иконку "скрепка"</i>).

<i>Отправьте '.' для отмены</i>`,
		"en-US": `Please enter a name for the new contact:
You can type manually or choose from your address book (<i>throw "clip" icon</i>).

<i>Send '.' to cancel</i>`,
	},
	MESSAGE_TEXT_TRANSFER_IS_CREATING: {
		"ru-RU": "Создаю запись...",
		"en-US": "Creating transfer...",
	},
	COMMAND_TEXT_PLEASE_WAIT: {
		"ru-RU": "Пожалуйста подождите",
		"en-US": "Please wait",
	},
	MESSAGE_TEXT_PLEASE_WAIT: {
		"ru-RU": "Пожалуйста подождите...",
		"en-US": "Please wait...",
	},
	MESSAGE_TEXT_SELF_ACKNOWLEDGEMENT: {
		"ru-RU": "Подтверждение ожидается от %v",
		"en-US": "Acknowledgement is expected from %v",
	},
	MESSAGE_TEXT_TRANSFER_ACCEPTED_BY_YOU: {
		"ru-RU": "Вы подтвердили эту транзакцию.",
		"en-US": "You've acepted this transaction.",
	},
	MESSAGE_TEXT_TRANSFER_DECLINED_BY_YOU: {
		"ru-RU": "Вы НЕ подтвердили эту транзакцию.",
		"en-US": "You've declined this transaction.",
	},
	MESSAGE_TEXT_TRANSFER_ACCEPTED_BY_COUNTERPARTY: {
		"ru-RU": "%v подтвердил(a) вашу транзакцию:",
		"en-US": "%v accepted your transaction:",
	},
	MESSAGE_TEXT_TRANSFER_DECLINED_BY_COUNTERPARTY: {
		"ru-RU": "%v НЕ подтвердил(a) вашу транзакцию.",
		"en-US": "%v declined your transaction.",
	},
	COMMAND_TEXT_SUBSCRIBE_TO_APP: {
		"ru-RU": "Хочу приложение!",
		"en-US": "I want the app!",
	},
	COMMAND_TEXT_I_AM_FINE_WITH_BOT: {
		"ru-RU": "Меня вполне устраивает бот!",
		"en-US": "I'm fine with just the bot!",
	},
	MESSAGE_TEXT_SUBSCRIBED_TO_APP: {
		"ru-RU": "Мы сообщим вам когда приложение будет доступно для загруки.",
		"en-US": "We'll let you know once the app is available for download.",
	},
	MESSAGE_TEXT_NOT_INTERESTED_IN_APP: {
		"ru-RU": "Чтож, мы рады что вас устраивает наш бот и нет необходимости загружать приложение.",
		"en-US": "Well, we are happy our bot is good enough and there is no need to download an app.",
	},
	MESSAGE_TEXT_YOUR_AD_COULD_BE_HERE: {
		"ru-RU": "Здесь можно <a href>разместить рекламу</a>",
		"en-US": "You can <a href>advertise here</a>",
	},
	MESSAGE_TEXT_YOUR_ABOUT_ADS: {
		"ru-RU": emoji.ROBOT_ICON + `: Я конечно обоятельный робот, но пользоваться специализированным приложением бывает удобнее. Оно ещё не готово для общего доступа, но уже сейчас можно посмотреть как будет выглядеть: <a href="https://debtstracker.io/ru/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/ru/</a>

Хотите получить оповещение когда оно выйдет?`,
		"en-US": emoji.ROBOT_ICON + `: I'm a good robot, for sure. But sometimes it is more convinient to use a nice specialized app. It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/en/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/en/</a>

Do you want to get an invite when it gets released?
`,
	},
	MESSAGE_TEXT_INVALID_FLOAT: {
		"ru-RU": "Извините, но вы можете использовать только числа в качестве суммы/количества (<i>до 2х знаком после запятой</i>).",
		"en-US": "Sorry, but you can use just numbers as amount/quantity (<i>with up to 2 digits after point</i>).",
	},
	MESSAGE_TEXT_ASK_LENDING_TYPE: {
		"ru-RU": "<b>Что вы дали в долг?</b>",
		"en-US": "<b>What did you lend to someone?</b>",
	},
	MESSAGE_TEXT_CHOOSE_CURRENCY: {
		"ru-RU": `Выберите из меню внизу экрана или <a>выберите валюту из списка</a>.

Если ни один из стандартных вариантов не подходит просто напишите текстом. Например: "<i>яблоко</i>".`,

		"en-US": `Please choose from the options below or <a>select a currency from the list</a>.

If standard options are not enough simply send a text. For example: "<i>apple</i>".`,
	},
	MESSAGE_TEXT_ASK_LENDING_AMOUNT: {
		"ru-RU": "Сколько <b>%v</b> вы дали в долг?\n(<i>отправьте '.' чтобы отменить</i>)",
		"en-US": "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)",
	},
	MESSAGE_TEXT_ASK_LENDING_COUNTERPARTY: {
		"ru-RU": "Кому вы дали в долг <b>%v</b>?\n(<i>отправьте '.' чтобы отменить</i>)",
		"en-US": "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
	},
	MESSAGE_TEXT_ASK_BORROWING_TYPE: {
		"ru-RU": "Что вы взяли в долг?",
		"en-US": "What did you lend?",
	},
	MESSAGE_TEXT_ASK_BORROWING_AMOUNT: {
		"ru-RU": "Сколько <b>%v</b> вы взяли в долг?\n(<i>отправьте '.' чтобы отменить</i>)",
		"en-US": "How much <b>%v</b> did you borrow?\n(<i>send '.' to cancel</i>)",
	},
	MESSAGE_TEXT_ASK_BORROWING_COUNTERPARTY: {
		"ru-RU": "У кого вы взяли в долг <b>%v</b>?\n(<i>отправьте '.' чтобы отменить</i>)",
		"en-US": "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
	},
	MESSAGE_TEXT_YOU_CAN_SEND_RECEIPT: {
		"ru-RU": "Отправить <a receipt>квитанцию</a> для <a counterparty>%v</a>?",
		"en-US": "Should we send a <a receipt>receipt</a> to <a counterparty>%v</a>?",
	},
	MESSAGE_TEXT_YOU_CAN_SEND_RECEIPT_TO_YOURSELF_BY_SMS: {
		"ru-RU": "К сожалению отправка квитанцию себе по СМС в данный момент отключена. Но вы можете отправить её для %v.",
		"en-US": "Sorry, sending a receipt to yourself by SMS is not available at the moment. You can send it to %v though.",
	},
	MESSAGE_TEXT_RECEIPT_IS_SENDING_BY_TELEGRAM: {
		"ru-RU": "Отправляем для %v извещение через Telegram...",
		"en-US": "We are sending receipt to %v by Telegram...",
	},
	MESSAGE_TEXT_RECEIPT_NEW_DEBT_FROM_USER: {
		"ru-RU": "{{.Counterparty}} взял(а) в долг {{.Amount}}.",
		"en-US": "{{.Counterparty}} borrowed from you {{.Amount}}.",
	},
	MESSAGE_TEXT_RECEIPT_NEW_DEBT_TO_USER: {
		"ru-RU": "{{.Counterparty}} дал(а) вам в долг {{.Amount}}.",
		"en-US": "{{.Counterparty}} lended to you {{.Amount}}.",
	},
	MESSAGE_TEXT_RECEIPT_RETURN_FROM_USER: {
		"ru-RU": "Вы вернули долг - {{.Counterparty}} получил(а) {{.Amount}}.",
		"en-US": "You returned {{.Amount}} to {{.Counterparty}}.",
	},
	MESSAGE_TEXT_RECEIPT_RETURN_TO_USER: {
		"ru-RU": "{{.Counterparty}} вернул вам {{.Amount}}.",
		"en-US": "{{.Counterparty}} returned to you {{.Amount}}.",
	},
	MESSAGE_TEXT_DUE_ON: {
		"ru-RU": "<b>Вернуть до</b>: %v",
		"en-US": "<b>Return till</b>: %v",
	},
	MESSAGE_TEXT_NOTE: {
		"ru-RU": "Заметка",
		"en-US": "Note",
	},
	MESSAGE_TEXT_COMMENT: {
		"ru-RU": "Комментарий",
		"en-US": "Comment",
	},
	MESSAGE_TEXT_LOGIN_TO_WEB_APP: {
		"ru-RU": `Перейдите по <a>ссылке</a> чтобы запустить web-приложение.`,
		"en-US": `Click to <a>sign in</a> to web-app.`,
	},
	MESSAGE_TEXT_DO_YOU_LIKE_OUR_BOT: {
		"ru-RU": "Вам нравится наш бот?",
		"en-US": "Do you like our bot?",
	},
	MESSAGE_TEXT_ASK_FOR_FEEDBAK: {
		"ru-RU": "Будем признетельны если вы оцените работу нашего приложения. Это займёт всего несколько секунд.",
		"en-US": "We would appreciate if tell us how we doing. It takes just few seconds.",
	},
	COMMAND_TEXT_GIVE_FEEDBACK: {
		"ru-RU": "Оценить приложение",
		"en-US": "Rate this bot",
	},
	COMMAND_TEXT_OPEN_STOREBOT_FOR_FEEDBACK: {
		"ru-RU": "Оценить на  @Storebot",
		"en-US": "Leave rating at @Storebot",
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
	},
	MESSAGE_TEXT_ASK_TO_TRANSLATE: {
		"ru-RU": `Хотите чтобы наш бот разговаривал на другом языке? Вы можете <a>помочь с переводом</a>.`,
		"en-US": `Do you want our bot to talk in other language? You can <a>help with translation</a>.`,
	},
	MESSAGE_TEXT_ON_FEEDBACK_NEUTRAL: {
		"ru-RU": `Чтож, мы очень старались. Ваша оценка будет передана разработчикам.

Может быть вы <a submit-bug>сообщите нам что не работает</a> или подскажите <a suggest-idea>как можно улучшить</a>?`,
		/*------------------------------------------------------------*/
		"en-US": `Well, we worked hard. You feedback will be passed to developers.

Maybe you can <a submit-bug>report your issue</a> or <a suggest-idea>suggest how we can improve</a>?`,
	},
	MESSAGE_TEXT_ON_FEEDBACK_NEGATIVE: {
		"ru-RU": `Нам очень стыдно. Может быть вы <a submit-bug>подскажите что не так</a> или <a suggest-idea>предложите усовершенствования</a>?`,
		/*------------------------------------------------------------*/
		"en-US": `We are very sorry. Maybe you can <a submit-bug>let us know what is wrong</a> or <a suggest-idea>suggest how we can improve</a>?`,
	},
	COMMAND_TEXT_ASK_FOR_FEEDBACK: {
		"ru-RU": "Оцените наше приложение?",
		"en-US": "Please rate our app",
	},
	COMMAND_TEXT_FEEDBACK_POSITIVE: {
		"ru-RU": "Да, отличное приложение!",
		"en-US": "Yes, it's a great app!",
	},
	COMMAND_TEXT_FEEDBACK_NEUTRAL: {
		"ru-RU": "Неплохо, но можно лучше.",
		"en-US": "Not bad but can be better.",
	},
	COMMAND_TEXT_FEEDBACK_NEGATIVE: {
		"ru-RU": "Не нравится",
		"en-US": "Don't like it",
	},
	COMMAND_TEXT_FEEDBACK_NOT_READY: {
		"ru-RU": "Пока не понятно",
		"en-US": "Not decided yet",
	},
	MESSAGE_TEXT_SETTINGS: {
		"ru-RU": "Что будем настраивать?",
		"en-US": "What do you want to adjust?",
	},
	MESSAGE_TEXT_NOT_IMPLEMENTED_YET: {
		"ru-RU": "Извините, данный функционал ещё не реализован.",
		"en-US": "Sorry, this functionality is not implemented yet.",
	},
	MESSAGE_TEXT_ASK_INVITE_CHANNEL: {
		"ru-RU": "Как вы хотите получить код приглашения?",
		"en-US": "How do you want to get an invite?",
	},
	MESSAGE_TEXT_PLEASE_ENTER_INVITE_CODE: {
		"ru-RU": "Пожалуйста введите код приглашения:",
		"en-US": "Please enter an invite code:",
	},
	MESSAGE_TEXT_USER_EMAIL_FOR_INVITE_RECEIVED: {
		"ru-RU": "Мы отправили письмо на %v.\n\nПожалуйста откройте его и кликните на ссылку для подтверждения адреса.",
		"en-US": "We've sent a message to %v.\n\nPlease open the email and click a link to confirm your email address.",
	},
	MESSAGE_TEXT_USER_EMAIL_FOR_INVITE_SENT_TELEGRAM: {
		"ru-RU": "После того как откроется Telegram нажмите кнопку <b>Start</b>.",
		"en-US": "Once Telegram pop ups click the <b>Start</b> button.",
	},
	MESSAGE_TEXT_USER_CONTACT_FOR_INVITE_RECEIVED: {
		"ru-RU": "Спасибо, вы записаны в очередь на получение приглашения.\n\nТекущее время ожидания 2-3 дня.\n\nВы можете получить приглашение сегодня если расскажите о нашем боте на Facebook.",
		"en-US": "Thanks, you've been queued for an invite.\n\nCurrent awaiting time is 2-3 days.\n\nYou can get an invite code today by sharing a link to the bot on Facebook.",
	},
	MESSAGE_TEXT_PLEASE_PROVIDE_YOUR_EMAIL: {
		"ru-RU": "Пожалуйста напишите ваш email адрес:",
		"en-US": "Please provide your email address",
	},
	MESSAGE_TEXT_PLEASE_PROVIDE_YOUR_PHONE_NUMBER: {
		"ru-RU": "Пожалуйста напишите номер вашего телефона:",
		"en-US": "Please provide your email address",
	},
	MESSAGE_TEXT_WRONG_INVITE_CODE: {
		"ru-RU": "Неправильный код приглашения: %v",
		"en-US": "Wrong invite code: %v",
	},
	MESSAGE_TEXT_WRONG_EMAIL: {
		"ru-RU": "Неправильный email адрес.",
		"en-US": "Wrong email address.",
	},
	MESSAGE_TEXT_WRONG_PHONE_NUMBER: {
		"ru-RU": "Неправильный номер телефона.",
		"en-US": "Wrong phone number.",
	},
	MESSAGE_TEXT_OK_PLEASE_TRY_AGAIN: {
		"ru-RU": "Хорошо, попробуйте ещё раз.",
		"en-US": "Ok, please try again.",
	},

	COMMAND_TEXT_MISTYPE_WILL_TRY_AGAIN: {
		"ru-RU": "Я опечатался, попробую ещё раз.",
		"en-US": "I've mistyped, will try again.",
	},
	COMMAND_TEXT_TELL_ME_MORE_ABOUT_INVITES: {
		"ru-RU": "Расскажите ка мне об этих кодах",
		"en-US": "Tell me more about the codes",
	},
	COMMAND_TEXT_INVITE_ME_BY_EMAIL: {
		"ru-RU": "Хочу код приглашения на email",
		"en-US": "Send me invite code by email",
	},
	COMMAND_TEXT_INVITE_ME_BY_SMS: {
		"ru-RU": "Хочу код приглашения по SMS",
		"en-US": "Send me invite code by SMS",
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_EMAIL: {
		"ru-RU": "Новый код приглашения на email",
		"en-US": "Send me new invite code by email",
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_SMS: {
		"ru-RU": "Новый код приглашения по SMS",
		"en-US": "Send me new invite code by SMS",
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_TELEGRAM: {
		"ru-RU": "Получить приграшение в Telegram",
		"en-US": "Send me new invite by Telegram",
	},
	MESSAGE_TEXT_UNKNOWN_LANGUAGE: {
		"ru-RU": "Незнакомый язык. Пожалуйста выберете один из предоставленных:",
		"en-US": "Unknown language. Please choose 1 from the options:",
	},
	MESSAGE_TEXT_UNKNOWN_COUNTERPARTY: {
		"ru-RU": "Неизвестный контакт. Пожалуйста выберите <b>Добавить</b> если это новый контакт.",
		"en-US": "Unknown counterparty. Please choose <b>Add new</b> if it's a new contact.",
	},
	MESSAGE_TEXT_UNKNOWN_COUNTERPARTY_ON_RETURN: {
		"ru-RU": "Неизвестный контакт. Пожалуйста выберите из списка.",
		"en-US": "Unknown counterparty. Please choose from the list.",
	},
	MESSAGE_TEXT_UNKNOWN_DEBT: {
		"ru-RU": "Неизвестный долг. Пожалуйста выберите из списка.",
		"en-US": "Unknown debt. Please choose from the list.",
	},

	MESSAGE_TEXT_HI: { // This is the same for all languages.
		"ru-RU": `¡Hola! Hi! Привет! سلام!`,
		"en-US": `¡Hola! Hi! Привет! سلام!`,
	},
	MESSAGE_TEXT_BACK_TO_MAIN_MENU: {
		"ru-RU": `Можно вернуться назад в главное /меню`,
		"en-US": `You can go back to main /menu`,
	},
	MESSAGE_TEXT_YOUR_SELECTED_PREFERRED_LANGUAGE: { // This is the same for all languages.
		"ru-RU": `Выбранный язык программы: %v`,
		"en-US": `Preferred app language: %v`,
	},
	MESSAGE_TEXT_ONBOARDING_ASK_TO_CHOOSE_LANGUAGE: {
		"ru-RU": `<b>%v</b>, на каком языке вы хотели бы общаться?

(What is your preferred language?)`,
		"en-US": `<b>%v</b>, what is your preferred language?`,
	},
	MESSAGE_TEXT_CHOOSE_UI_LANGUAGE: {
		"ru-RU": "На каком языке вы хотели бы общаться со мной?",
		"en-US": "Which language you would like to talk to me?",
	},
	MESSAGE_TEXT_LOCALE_CHANGED: {
		"ru-RU": "Вы поменяли язык на %v",
		"en-US": "You've switched language to %v",
	},
	MESSAGE_TEXT_WHATS_NEXT: {
		"ru-RU": "Что будем делать дальше?",
		"en-US": "What's next?",
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
	},
	MESSAGE_TEXT_HISTORY_HEADER: {
		"ru-RU": "История",
		"en-US": "History",
	},
	MESSAGE_TEXT_HISTORY_NO_RECORDS: {
		"ru-RU": "У вас пока нет ни одной записи.",
		"en-US": "You don't have any records yet.",
	},
	MESSAGE_TEXT_HISTORY_LIST: {
		"ru-RU": `<b>%v</b> <i>(%d последних)</i>
─────────────
%v`,
		"en-US": `<b>%v</b> <i>(last %d):</i>
─────────────
%v`,
	},
	MESSAGE_TEXT_BALANCE_IS_ZERO: {
		"ru-RU": "Нет записей о текущих долгах.",
		"en-US": "You have no records on current debts.",
	},
	MESSAGE_TEXT_BALANCE_CURRENCY_TOTAL_INTRO: {
		"ru-RU": "Всего",
		"en-US": "Total",
	},
	MESSAGE_TEXT_PRIMARY_CURRENCY_IS_SET_TO: {
		"ru-RU": "OK, теперь я буду использовать '%v' как основную валюту.",
		"en-US": "OK, from now on I will use '%v' as a primary currency.",
	},
	MESSAGE_TEXT_BALANCE_SINGLE_CURRENCY_COUNTERPARTY_DEBT_TO_USER: {
		"ru-RU": "<b>%v</b> - долг вам %v",
		"en-US": "<b>%v</b> - owes you %v",
	},
	MESSAGE_TEXT_BALANCE_CURRENCY_ROW_DEBT_TO_USER: {
		"ru-RU": "Вам должны %v",
		"en-US": "Owes to you %v",
	},
	MESSAGE_TEXT_ON_RETURN_USER_DOES_NOT_OWE_ANYTHING_TO_COUNTERPARTY_ANYMORE: {
		"ru-RU": "Поздравляем! У вас не осталось долгов перед <b>%v</b>.",
		"en-US": "Congratulations! You don't owe anything more to <b>%v</b>.",
	},
	MESSAGE_TEXT_ON_RETURN_COUNTERPARTY_DOES_NOT_OWE_ANYTHING_TO_USER_ANYMORE: {
		"ru-RU": "У <b>%v</b> больше не осталось долгов перед вами.",
		"en-US": "<b>%v</b> does not owe anything more to you.",
	},
	MESSAGE_TEXT_BALANCE_CURRENCY_ROW_DEBT_BY_USER: {
		"ru-RU": "Вы должны %v",
		"en-US": "You owe %v",
	},
	MESSAGE_TEXT_BALANCE_SINGLE_CURRENCY_COUNTERPARTY_DEBT_BY_USER: {
		"ru-RU": "<b>%v</b> - вы должны %v",
		"en-US": "<b>%v</b> - you owe %v",
	},
	MESSAGE_TEXT_ASK_PRIMARY_CURRENCY: {
		"ru-RU": "Какая валюта для вас основная?",
		"en-US": "What is your primary currency?",
	},
	MESSAGE_TEXT_FAILED_TO_DELETE_USER: {
		"ru-RU": "Не удалось удалить данные пользователя: %v",
		"en-US": "Failed to delete user: %v",
	},
	MESSAGE_TEXT_USER_DELETED: {
		"ru-RU": "Данные пользователя удалены",
		"en-US": "User's data has been deleted",
	},
	MESSAGE_TEXT_PLEASE_WAIT_WHILE_WE_GENERATE_INVITE_CODE: {
		"ru-RU": "Пожалуйста подождите пока мы генерируем секретный код доступа...",
		"en-US": "Please wait a moment while we are generating a security access code...",
	},
	MESSAGE_TEXT_RETURN_ASK_TO_CHOOSE_COUNTERPARTY: {
		"ru-RU": "Выберете кому вы вернули долг или кто вернул его вам.",
		"en-US": "Please choose who returned the debt or to who you returned it.",
	},
	MESSAGE_TEXT_CHOOSE_DEBT_THAT_HAS_BEEN_RETURNED: {
		"ru-RU": "Выберите долг который был возвращён целиком или частично.",
		"en-US": "Please choose a debt that has been returned fully or partially.",
	},
	MESSAGE_TEXT_PLEASE_ACKNOWLEDGE_TRANSFER: {
		"ru-RU": "Пожалуйста подтвердите или отклоните эту транзакцию.",
		"en-US": "Please confirm or decline this transfer.",
	},
	MESSAGE_TEXT_ALREADY_ACCEPTED_TRANSFER: {
		"ru-RU": "Эта транзакция уже подтверждена.",
		"en-US": "This transfer has been accepted already.",
	},
	MESSAGE_TEXT_ALREADY_DECLINED_TRANSFER: {
		"ru-RU": "Эта транзакция уже отклонена.",
		"en-US": "This transfer has been declined already.",
	},
	MESSAGE_TEXT_RECEIPT_LINK: {
		"ru-RU": "Подробнее тут: %v",
		"en-US": "Details here: %v",
	},
	MESSAGE_TEXT_ASK_PHONE_NUMBER_OF_COUNTERPARTY: {
		"ru-RU": "Пожалуйста напишите номер телефона <b>%v</b>.",
		"en-US": "Plese provide phone number for <b>%v</b>",
	},
	MESSAGE_TEXT_USE_CONTACT_TO_SEND_PHONE_NUMBER: {
		"ru-RU": "Если номер телефона есть в записной книжке <b>воспользуйтесь кнопкой %v</b> (скрепка) чтобы отправить контакт.",
		"en-US": "If phone number is in your address book you can <b>use %v button</b> to send the contact.",
	},
	MESSAGE_TEXT_ABOUT_PHONE_NUMBER_FORMAT: {
		"ru-RU": "Номер должен быть в международном формате:\n\t* Начинаться со знака '+' и кода страны\n\t* Состоять только из цифр\nПример: <pre>+</pre><b>7</b><code>999012345678</code>",
		"en-US": "The number should be in international standard:\n\t* Starts with '+' following by country code\n\t* Consist of numbers only\nExample: <pre>+</pre><b>1</b><code>999012345678</code>",
	},
	MESSAGE_TEXT_THIS_NUMBER_WILL_BE_USED_TO_SEND_RECEIPT: {
		"ru-RU": "На этот номер мы отправим SMS:",
		"en-US": "Will send an SMS to this number:",
	},
	MESSAGE_TEXT_COUNTERPARTY_OWES_YOU_SINGLE_DEBT: {
		"ru-RU": `<b>%v</b> одалживал(а) у вас <b>%v</b>.`,
		"en-US": `<b>%v</b> owed to you <b>%v</b>.`,
	},
	MESSAGE_TEXT_YOU_OWE_TO_COUNTERPARTY_SINGLE_DEBT: {
		"ru-RU": `<b>%v</b> одалживал(а) вам <b>%v</b>.`,
		"en-US": "You owe to <b>%v</b> <b>%v</b>.",
	},
	MESSAGE_TEXT_IS_IT_RETURNED_IN_FULL: {
		"ru-RU": `Возвращено полностью?

		<i>Если частично можете сразу написать сумму.</i>`,

		"en-US": `Has this debt been returned in full?

		<i>If partially you can enter amount right away.</i>`,
	},
	MESSAGE_TEXT_PLEASE_HELP_MAKE_IT_BETTER: {
		"ru-RU": `Эта программа <b>бесплатна</b>. <a href="https://debtstracker.io/ru/help-us">Помогите</a> сделать её лучше!`,

		"en-US": `This program is <b>free to use</b>. Please <a href="https://debtstracker.io/en/help-us">help</a> to make it better!`,
	},
	BUTTON_TEXT_YOU_OWE_AMOUNT_TO_SOMEONE: {
		"ru-RU": "%v | вы должны: %v",
		"en-US": "%v | you owe: %v",
	},
	BUTTON_TEXT_SOMEONE_OWES_TO_YOU_AMOUNT: {
		"ru-RU": "%v | долг вам: %v",
		"en-US": "%v | owes to you: %v",
	},
	BUTTON_TEXT_DEBT_RETURNED_FULLY: {
		"ru-RU": "Да, целиком",
		"en-US": "Yes, fully",
	},
	BUTTON_TEXT_DEBT_RETURNED_PARTIALLY: {
		"ru-RU": "Нет, только часть",
		"en-US": "No, just partially",
	},
	MESSAGE_TEXT_ATTEMPT_TO_USE_OWN_INVITE: {
		"ru-RU": "Хорошая попытка пригласить самого себя ;)",
		"en-US": "You should not use your own invite ;)",
	},
	MESSAGE_TEXT_WELCOME_ONBOARDING_INVITE_ACCEPTED: {
		"ru-RU": "Спасибо за то что воспользовались приглашением!",
		"en-US": "Welcome and thanks for accepting the invite!",
	},
	MESSAGE_TEXT_FOR_COUNTERPARTY_ONLY: {
		"ru-RU": "Это действие доступно только для %v",
		"en-US": "This action for %v only.",
	},
	BUTTON_TEXT_SEE_RECEIPT_DETAILS: {
		"ru-RU": "Показать детали",
		"en-US": "Show receipt details",
	},
	MESSAGE_TEXT_YOU_SELECTED_INVITE_BY_EMAIL: {
		"ru-RU": "Вы решили пригласить друга через email.",
		"en-US": "You've selected to invite friend by email.",
	},
	MESSAGE_TEXT_YOU_SELECTED_INVITE_BY_SMS: {
		"ru-RU": "Вы решили пригласить друга через SMS.",
		"en-US": "You've selected to invite friend by SMS.",
	},
	MESSAGE_TEXT_ABOUT_INVITES: {
		"ru-RU": `На данный момент доступ к нашему боту ограничен, но вы можете пригласить друга.

Как вы хотите передать код приглашение?`,
		"en-US": `At the moment access to our bot is limited but you can invite your friend.

How do you want to pass the invite code?`,
	},
	MESSAGE_TEXT_USER_BLOCKED_TRANSFER_NOTIFICATIONS_BY: {
		"ru-RU": "%v заблокировал получение оповешений о транзакиях через: %v.",
		"en-US": "%v blocked reminders about transactions by: %v",
	},
	COMMAND_TEXT_WAIT_A_SECOND: {
		"ru-RU": "Секундочку...",
		"en-US": "Wait a second...",
	},
	HTML_USING_TELEGRAM: {
		"ru-RU": "используя Telegram мессенджер",
		"en-US": "using Telegram messenger",
	},
	COMMAND_TEXT_ACCEPT: {
		"ru-RU": "Согласиться",
		"en-US": "Accept",
	},
	//BUTTON_TEXT_ACCEPT_TRANSFER_USING_TELEGRAM:{
	//	"ru-RU": "Подтвердить ",
	//	"en-US": "Accept ",
	//},
	//BUTTON_TEXT_DECLINE_TRANSFER_USING_TELEGRAM:{
	//	"ru-RU": "Отказаться (используя Telegram)",
	//	"en-US": "Decline (using Telegram messenger)",
	//},
	COMMAND_TEXT_DECLINE: {
		"ru-RU": "Отклонить",
		"en-US": "Decline",
	},
	COMMAND_TEXT_ACCEPT_INVITE: {
		"ru-RU": "Принять приглашение",
		"en-US": "Accept invite",
	},
	COMMAND_TEXT_VIEW_RECEIPT_DETAILS: {
		"ru-RU": "Посмотреть квитанцию",
		"en-US": "See receipt details",
	},
	COMMAND_TEXT_OTHER_WAYS_TO_SEND_INVITE: {
		"ru-RU": "Другие способы отправить приглашение",
		"en-US": "Other ways to send an invite",
	},
	COMMAND_TEXT_SEND_MY_PHONE_NUMBER: {
		"ru-RU": "Отправить мой номер",
		"en-US": "Send my phone number",
	},
	COMMAND_TEXT_SEND_BY_EMAIL: {
		"ru-RU": "Через Email",
		"en-US": "By Email",
	},
	COMMAND_TEXT_SEND_BY_SMS: {
		"ru-RU": "Через SMS",
		"en-US": "By SMS",
	},
	COMMAND_TEXT_INVITE_BY_TELEGRAM: {
		"ru-RU": "Пригласить через Telegram",
		"en-US": "InviteBy Telegram",
	},
	MESSAGE_TEXT_INVITE_CREATED: {
		"ru-RU": `Мы отправили код приглашения на указынный вами адрес. (#%v)

Когда ваш друг потдвердит приглашение у вас будут общий баланс и транзакции (только между вами). Это поможет вам минимизировать усилия по ведению учёта.`,

		"en-US": `We've sent an invite code to your friend. (#%v)

Once your friend accepts invitation you'll share balance & transfers between you to make sure you both are on the same page with minimum efforts.`,
	},
	MESSAGE_TEXT_INVITE_BY_EMAIL: {
		"ru-RU": "Введите email вашего друга на который мы отправим код приглашения.",
		"en-US": "Please enter emaill address of your friend where we should send an invite code.",
	},
	MESSAGE_TEXT_INVITE_ASK_EMAIL_FOR_RECEIPT: {
		"ru-RU": "Введите email вашего друга (%v) на который мы отправим квитанцию подтверждения.",
		"en-US": "Please enter emaill address of your friend (%v) where we should send the receipt.",
	},
	MESSAGE_TEXT_INVITE_BY_SMS: {
		"ru-RU": "Введите номер телефона вашего друга на который мы отправим код приглашения.",
		"en-US": "Please share a contact or enter phone number of your friend where we should send an invite code.",
	},
	MESSAGE_TEXT_INVITE_BY_TELEGRAM: {
		"ru-RU": "Вставьте в чат контакт вашего друга которому вы хотите отправить приглашение.",
		"en-US": "Please share a contact of your friend you wish to send an invite code:",
	},
	MESSAGE_TEXT_INVALID_EMAIL: {
		"ru-RU": "Неверный email. Проверьте и попробуйте ещё раз? /menu",
		"en-US": "Invalid email. Check and try it again? /menu",
	},
	MESSAGE_TEXT_INVALID_YEAR: {
		"ru-RU": "Неверный год. Год должен быть 2 или 4 цифры (<i>например 2016 или 16)</i>).",
		"en-US": "Invalid year. The year part should be 2 or 4 numbers (<i>e.g. 2016 or 16</i>).",
	},
	MESSAGE_TEXT_INVALID_MONTH: {
		"ru-RU": "Неверный месяц. Месяц должен быть задан целым числом от 1 до 12.",
		"en-US": "Invalid month. Month should be an integer from 1 to 12.",
	},
	MESSAGE_TEXT_INVALID_DAY: {
		"ru-RU": "Неверный день. День должен быть задан целым числом от 1 до 31.",
		"en-US": "Invalid day. The day should be an integer from 1 to 31.",
	},
	MESSAGE_TEXT_INVALID_DATE: {
		"ru-RU": "Неверный формат даты. Например для 20 февраля 2019 года надо ввести: 20.02.2019 или 20.02.19",
		"en-US": "Invalid date format. For exampel for 20 February 2019 please submit: 20.02.2019 or 20.02.19",
	},
	MESSAGE_TEXT_INVALID_PHONE_NUMBER: {
		"ru-RU": "Неверный номер. Проверьте и попробуйте ещё раз? /menu",
		"en-US": "Invalid phone number. Check and try it again? /menu",
	},
	MESSAGE_TEXT_PHONE_NUMBER_IS_NOT_SMS_CAPABLE: {
		"ru-RU": "Данный номер не принимает SMS. Попробуйте другой номер? /menu",
		"en-US": "This phone number not able to receive SMS. Try another number? /menu",
	},
	MESSAGE_TEXT_NO_CONTACT_RECEIVED: {
		"ru-RU": "Мы не получили контакта. Тут инструкция как это сделать. /menu",
		"en-US": "We have not received any contacts. INSTRUCTION HOW TO DO IT. /menu",
	},
	MESSAGE_TEXT_CONTACT_NAME_IS_NUMBER: {
		"ru-RU": "Вы ввели только цифры в качестве имени контакта. Пожалуйста используйте текстовые символы.",
		"en-US": "You've entered just digits for a contact name. Please use some text characters.",
	},
	MESSAGE_TEXT_CURRENCY_NAME_IS_NUMBER: {
		"ru-RU": "Вы ввели только цифры в качестве номинала. Пожалуйста используйте текстовые символы.",
		"en-US": "You've entered just digits for currency. Please use some text characters.",
	},
	MESSAGE_TEXT_HISTORY_ROW_TO_USER_WITH_NAME: {
		"ru-RU": "%v - %s ⇒ Вам : %s",
		"en-US": "%v - %s ⇒ to you: %s",
	},
	MESSAGE_TEXT_HISTORY_ROW_FROM_USER_WITH_NAME: {
		"ru-RU": "%v - Вы ⇒ %s : %s",
		"en-US": "%v - You ⇒ %s : %s",
	},
	MESSAGE_TEXT_LETS_SEND_SMS: {
		"ru-RU": "Давайте отправим SMS",
		"en-US": "Let's send SMS",
	},
	MESSAGE_TEXT_SMS_QUEUING_FOR_SENDING: {
		"ru-RU": "SMS ставится в очередь на отправку на номер %v...",
		"en-US": "Queuing SMS for sending to number %v...",
	},
	MESSAGE_TEXT_SMS_QUEUED_FOR_SENDING: {
		"ru-RU": "SMS поставлена в очередь на отправку на номер %v",
		"en-US": "SMS is queued for sending to number %v",
	},
	MESSAGE_TEXT_BALANCE_HEADER: {
		"ru-RU": "Баланс",
		"en-US": "Balance",
	},
	MESSAGE_TEXT_RECEIPT_AVAILABLE_CHANNELS: {
		"ru-RU": "Извините, в данный момент доступны только эти каналы для отправки квитанции:",
		"en-US": "Sorry, at the moment just this channels are available for sending a receipt:",
	},
	MESSAGE_TEXT_RECEIPT_SENT_THROW_TELEGRAM: {
		"ru-RU": "Квитанция отправлена через телеграм.",
		"en-US": "Receipt sent through Telegram.",
	},
	MESSAGE_TEXT_RECEIPT_NOT_SENT_AS_COUNTERPARTY_HAS_DISABLED_TG_BOT: {
		"ru-RU": "Квитанция НЕ отправлена через телеграм так как %v удалил чат с ботом.",
		"en-US": "Receipt NOT sent through Telegram as %v has deleted chat with the bot.",
	},
	MESSAGE_TEXT_RECEIPT_SENT_THROW_EMAIL: {
		"ru-RU": "Квитанция отправлена через email. (id: %v)",
		"en-US": "Receipt sent through email. (id: %v)",
	},
	MESSAGE_TEXT_RECEIPT_SENT_THROW_SMS: {
		"ru-RU": "Квитанция отправлена через SMS.",
		"en-US": "Receipt sent through SMS.",
	},
	MESSAGE_TEXT_SWITCH_TO_PM_TO_VIEW_RECEIPT: {
		"ru-RU": "Переключитьсь на чат с ботом чтобы посмотреть квитанцию",
		"en-US": "Switch to private mode to see receipt details.",
	},
	MESSAGE_TEXT_RECEIPT_VIEWED_BY_COUNTERPARTY: {
		"ru-RU": "Квитанция просмотрена",
		"en-US": "Receipt viewed",
	},
	MESSAGE_TEXT_VIEW_MY_NUMBER_IN_INTERNATIONAL_FORMAT: {
		"ru-RU": "Вы можете посмотреть свой номер телефона в ожидаемоем нами формате.",
		"en-US": "You can view your own phone number in the format we are expecting.",
	},
	COMMAND_TEXT_VIEW_MY_NUMBER_IN_INTERNATIONAL_FORMAT: {
		"ru-RU": "Посмотреть мой номер в ожидаемоем формате",
		"en-US": "View my number in the expectd format",
	},
	INLINE_BUTTON_SHOW_FULL_HISTORY: {
		"ru-RU": "Показать всю историю",
		"en-US": "Show full history",
	},
	MESSAGE_TEXT_INCORRECT_VALUE_NOT_A_NUMBER: {
		"ru-RU": "Это не цифра",
		"en-US": "it is not a number",
	},
	MESSAGE_TEXT_INCORRECT_VALUE_IS_NEGATIVE: {
		"ru-RU": "Цифра должна быть положительной (<i>больше нуля</i>)",
		"en-US": "The number should be positive (<i>greater than 0</i>)",
	},
	MESSAGE_TEXT_ASK_HOW_MUCH_HAS_BEEN_RETURNED: {
		"ru-RU": "Сколько было возвращено?",
		"en-US": "How much have been returned?",
	},
	MESSAGE_TEXT_HELP: {
		"ru-RU": "Вы можете сообщить о проблеме или сделать предложения по улучшению программы на нашем сайте.",
		"en-US": "Please report any issue or submit a feature request at our website.",
	},
	COMMAND_TEXT_OPEN_USER_REPORT: {
		"ru-RU": "Cтраница поддержки ",
		"en-US": "Support page",
	},
	COMMAND_TEXT_REPORT_A_BUG: {
		"ru-RU": "Сообщить об ошибке",
		"en-US": "Report a bug",
	},
	COMMAND_TEXT_SUBMIT_AN_IDEA: {
		"ru-RU": "Предложить идею",
		"en-US": "Add an idea",
	},
	MESSAGE_TEXT_WELCOME: {
		"ru-RU": `Привет, я Коллектиус - Ваш персональный счетовод и коллектор.

Могу записать кто кому чего должен и, и при необходимости, напомнить когда должок пора возвращать.

Чего изволит новый хозяин?`,
		"en-US": `Hi, I'm Collectius - your personal accountant & collector.

I can record who is owing to whom and remind when the return is due.

What would you like to do?`,
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE: {
		"ru-RU": "Хочу получить приглашение",
		"en-US": "I want to get an invite",
	},
	COMMAND_TEXT_I_HAVE_INVITE: {
		"ru-RU": "У меня есть код приглашения",
		"en-US": "I have the invitation code",
	},
	COMMAND_TEXT_I_HAVE_NOT_GOT_EMAIL: {
		"ru-RU": "Я не получил письма на email",
		"en-US": "I have not got any emails",
	},
	MESSAGE_TEXT_ONBOARDING_TELL_ABOUT_INVITES: {
		"ru-RU": `<b>%v</b>,

На данный момент наш бот доступен только тем кто получил приглашение от друзей.

Если у вас нет кода можете оставить свои контакты и мы вышлем вам приглашение как только наступит ваша очередь.

Мы высылаем по 10 кодов в день первоочередникам + 1 случайным образом.`,
		"en-US": `<b>%v</b>,

At the momnet our bot is available just by invitation from friends.

If you have no code you can leave you contact details and we'll send you an invite as soon as your queue is due.

We send 10 invites per day to those in the head of the queue and 1 randomly.`,
	},
	EMAIL_INVITE_SUBJ: {
		"ru-RU": "Приглашение от {{.FromName}} - код: {{.InviteCode}}",
		"en-US": "An invite from {{.FromName}} - code: {{.InviteCode}}",
	},
	SMS_INVITE_TEXT: {
		"ru-RU": `Привет {{.ToName}}, {{.FromName}} рекомендует приложение для учёта долгов - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Код приглашения: {{.InviteCode}}`,
		"en-US": `Hi{{.ToName}}, {{.FromName}} is inviting you to try debts tracking app - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

You invitation code is: {{.InviteCode}}`,
	},
	EMAIL_INVITE_TEXT: {
		"ru-RU": `Привет {{.ToName}},

{{.FromName}} приглашает тебя попробовать приложение для учёта долгов - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Ваш код приглашения: {{.InviteCode}}`,
		"en-US": `Hi{{.ToName}},

{{.FromName}} is inviting you to use debts tracking app - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

You invitation code is: {{.InviteCode}}`,
	},
	EMAIL_INVITE_HTML: {
		"ru-RU": `<p>Привет {{.ToName}},</P

<p>{{.FromName}} приглашает тебя <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">попробовать приложение для учёта долгов</a>.</p>

<p>Ваш код приглашения: <b>{{.InviteCode}}</b></p>`,
		"en-US": `<p>Hi{{.ToName}},</p>

<p>{{.FromName}} п is inviting you to <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">try debts tracking app</a>.</p>

<p>You invitation code is: <b>{{.InviteCode}}</b></p>`,
	},
	EMAIL_RECEIPT_SUBJ: {
		"ru-RU": "Запись о долге - {{.FromName}}",
		"en-US": "Debt record - {{.FromName}}",
	},
	EMAIL_RECEIPT_BODY_TEXT: {
		"ru-RU": "{{.FromName}} создал(а) запись о долге: {{.ReceiptURL}}",
		"en-US": "{{.FromName}} created a debt record: {{.ReceiptURL}}",
	},
	MESSENGER_RECEIPT_TEXT: {
		"ru-RU": "Я создал(а) запись о долге, подробности тут - {{.ReceiptURL}}",
		"en-US": "I've created a debt record regards you, see details at {{.ReceiptURL}}",
	},
	EMAIL_RECEIPT_BODY_HTML: {
		"ru-RU": "{{.FromName}} создал(а) запись о долге: {{.ReceiptURL}}",
		"en-US": "{{.FromName}} created a debt record: {{.ReceiptURL}}",
	},
	INLINE_RECEIPT_TITLE: {
		"ru-RU": "Квитанция: %v",
		"en-US": "Receipt: %v",
	},
	INLINE_RECEIPT_DESCRIPTION: {
		"ru-RU": "Нажмите здесь чтобы отправить квитанцию",
		"en-US": "Click here to send the receipt",
	},
	INLINE_RECEIPT_CHOOSE_LANGUAGE: {
		"ru-RU": "<b>Выберите язык чтобы посмотреть подробности записи о долге</b> которую создал(а) {{.Creator}}.",
		"en-US": "<b>Please choose language to see details of the debt</b> that has been recorded by {{.Creator}}.",
	},
	INLINE_RECEIPT_MESSAGE: {
		"ru-RU": `<b>{{.Creator}} создал(а) запись о долге</b> касающегося Вас.

{{.SiteLink}} — программа для учёта долгов поможет:

  - Всегда знать кто кому сколько должен

  - Незабыть вовремя отдать или востребовать долг
    <i>(напоминания вам и вашим должникам)</i>`,
		//-------------------------------------------------------
		"en-US": `<b>{{.Creator}} recorded a debt</b> associated with you.

{{.SiteLink}} — an app for debts tacking will help you to:

  - Always know your bottom line

  - Return debts on time
    <i>(reminders to you & your debtors)</i>`,
	},
	INLINE_INVITE_TITLE: {
		"ru-RU": "Приглашение в %v",
		"en-US": "Invitation to %v",
	},
	INLINE_INVITE_DESCRIPTION: {
		"ru-RU": "Нажмите здесь для отправки приглашения",
		"en-US": "Click here to send an invite",
	},
	INLINE_INVITE_MESSAGE: {
		"ru-RU": "%v пригласил вас попробовать %v",
		"en-US": "%v invited you to try %v",
	},
	SMS_RECEIPT_YOU_GOT: {
		"ru-RU": "Вы получили %v от %v.",
		"en-US": "You've got %v from %v.",
	},
	SMS_RECEIPT_YOU_GAVE: {
		"ru-RU": "Вы дали %v - взял(а) %v.",
		"en-US": "You've given %v to %v.",
	},
	SMS_CLICK_TO_CONFIRM_OR_DECLINE: {
		"ru-RU": "Нажмите %v чтобы подтвердить или отвергнуть.",
		"en-US": "Click %v to confirm or decline.",
	},
	HTML_DATE: {
		"ru-RU": "Дата",
		"en-US": "Date",
	},
	HTML_RECEIPT: {
		"ru-RU": "Квитанция",
		"en-US": "Receipt",
	},
	HTML_AMOUNT: {
		"ru-RU": "Сумма",
		"en-US": "Amount",
	},
	HTML_FROM: {
		"ru-RU": "Дал",
		"en-US": "From",
	},
	HTML_TO: {
		"ru-RU": "Получил",
		"en-US": "To",
	},
	TELEGRAM_RECEIPT: {
		"ru-RU": "{{.FromName}} создал запись о долге ({{.TransferCurrency}})",
		"en-US": "{{.FromName}} created a debtrecord ({{.TransferCurrency}})",
	},
	MESSAGE_TEXT_PLEASE_CHOOSE_FROM_OPTIONS_PROVIDED: {
		"ru-RU": "Пожалуйста выберете из предоставленных опций.",
		"en-US": "Please choose from provided options.",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_NOTE_OR_COMMENT: {
		"ru-RU": "<b>Хотите добавить заметку или комментарий?</b>\n%v Заметки хранятся для вашего личго пользования.\n%v Комментарий виден всем кому разрешён просмотр этой транзакции.",
		"en-US": "<b>Do you want to add a note or comment?</b>\n%v Memos are private records for yoru own reference.\n%v Comments are available to everyone who has permission to view this transaction.",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_NOTE: {
		"ru-RU": "Напишите заметку:",
		"en-US": "Please write a note:",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_COMMENT_ONLY: {
		"ru-RU": `Если хотите добавить комментарий просто отправьте текст.`,
		"en-US": `If you want to add a comment just send a text now.`,
	},
	MESSAGE_TEXT_VISIBLE_TO_YOU_AND_COUNTERPARTY: {
		"ru-RU": "виден вам и %v",
		"en-US": "visible to you & %v",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_COMMENT: {
		"ru-RU": "Напишите комментарий:",
		"en-US": "Please write the comment:",
	},
	MESSAGE_TEXT_TRANSFER_NOTE_ADDED_ASK_FOR_COMMENT: {
		"ru-RU": "Заметка добавлена. Хотите написать комментарий?",
		"en-US": "Memo have been added. Do you want to write a comment?",
	},
	MESSAGE_TEXT_TRANSFER_COMMENT_ADDED_ASK_FOR_NOTE: {
		"ru-RU": "Комментарий добавлен. Хотите написать заметку?",
		"en-US": "Comment have been added. Do you want to write a note?",
	},
	COMMAND_TEXT_NO_COMMENT_OR_NOTE_FOR_TRANSFER: {
		"ru-RU": "Без заметок и комментариев",
		"en-US": "Without notes or comments",
	},
	COMMAND_TEXT_NO_COMMENT_FOR_TRANSFER: {
		"ru-RU": "Без комментариев",
		"en-US": "No comments",
	},
	COMMAND_TEXT_NO_NOTE_FOR_TRANSFER: {
		"ru-RU": "Без заметок",
		"en-US": "Without notes",
	},
	COMMAND_TEXT_ADD_NOTE_TO_TRANSFER: {
		"ru-RU": "Добавить заметку",
		"en-US": "Add a note (private)",
	},
	COMMAND_TEXT_ADD_COMMENT_TO_TRANSFER: {
		"ru-RU": "Добавить комментарий",
		"en-US": "Add a comment (public)",
	},
	DUE_IN_NOW: {
		"ru-RU": "прямо сейчас",
		"en-US": "now",
	},
	DUE_IN_A_MINUTE: {
		"ru-RU": "через минуту",
		"en-US": "in a minute",
	},
	DUE_IN_X_MINUTES: {
		"ru-RU": "через %v минут(ы)",
		"en-US": "in %v minutes",
	},
	DUE_IN_AN_HOUR: {
		"ru-RU": "через час",
		"en-US": "in an hour",
	},
	DUE_IN_X_HOURS: {
		"ru-RU": "через %v часа/часов",
		"en-US": "in %v hours",
	},
	DUE_IN_X_DAYS: {
		"ru-RU": "через %v дня/дней",
		"en-US": "in %v days",
	},
	//-------------------------------------------------------------------------------------------------------------------
	WS_ALEX_T: {
		"ru-RU": "Александр Трахимёнок",
		"en-US": "Alexander Trakhimenok",
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
		"it-IT": "Bot Chat per Telegram messenger",
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
		"it-IT": "Open in Telegram &#x1F680;",
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
		"it-IT": "Al momento il nostro programma è disponibile solo su <a href='https://telegram.org/'> <b> Telegramma </b> messenger </a>",
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
		"it-IT": "Conosci il tuo equilibrio e non perdere mai un pagamento del debito!",
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
		"it-IT": "DebtsTracker.io è un servizio di applicazioni mobili e per ricordare che aiuta a monitorare i debiti e crediti. Invia notifiche e-mail e SMS automatici per i vostri debitori.",
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
	},
	WS_ADS_CONTENT: {
		"en-US": `To place ads in our app please send us an email to <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
		"ru-RU": `Чтобы разместить рекламу в нашем приложении пишите нам на <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
	},
	FB_MAKE_RECORD_HEADER: {
		"en-US": "Record debt",
		"ru-RU": "Записать долг",
	},
	FB_MAKE_RECORD_HEADLINE: {
		"en-US": "Scroll left to see other options.",
		"ru-RU": "Пролистайте карточки влево чтобы увидеть другие опции.",
	},
	FB_HOW_ARE_THINGS_HEADER: {
		"en-US": "How are you doing?",
		"ru-RU": "Как идут дела?",
	},
}
