package trans

import "github.com/strongo/bots-framework/core"

var TRANS = map[string]map[string]string{
	"EXAMPLE": map[string]string{
		"ru-RU": "ПРИМЕР",
		"en-US": "SAMPLE",
	},
	" and ": map[string]string{
		"ru-RU": " и ",
		"en-US": " and ",
	},
	bots.MESSAGE_TEXT_OOPS_SOMETHING_WENT_WRONG: map[string]string{
		"ru-RU": "Упс, что-то пошло не так... \xF0\x9F\x98\xB3",
		"en-US": "Oops, something went wrong... \xF0\x9F\x98\xB3",
	},
	MESSAGE_TEXT_ASK_DUE_DATE: map[string]string{
		"ru-RU": "Когда дата возврата?",
		"en-US": "When is the due date?",
	},
	COMMAND_TEXT_ON_SPECIFIC_DATE: map[string]string{
		"ru-RU": "Задать дату",
		"en-US": "On a specific date",
	},
	COMMAND_TEXT_TOMORROW: map[string]string{
		"ru-RU": "Завтра",
		"en-US": "Tomorrow",
	},
	COMMAND_TEXT_DAY_AFTER_TOMORROW: map[string]string{
		"ru-RU": "Послезавтра",
		"en-US": "Day after tomorrow",
	},
	COMMAND_TEXT_THIS_WEEK: map[string]string{
		"ru-RU": "На этой неделе",
		"en-US": "This week",
	},
	COMMAND_TEXT_YES_IT_HAS_RETURN_DEADLINE: map[string]string{
		"ru-RU": "Да, есть срок возврата!",
		"en-US": "Yes, it has a deadline!",
	},
	COMMAND_TEXT_NO_IT_CAN_BE_RETURNED_ANYTIME: map[string]string{
		"ru-RU": "Нет, срок возврата не важен.",
		"en-US": "No, whenever is fine.",
	},
	COMMAND_TEXT_IT_CAN_BE_RETURNED_ANYTIME: map[string]string{
		"ru-RU": "Срок возврата не важен.",
		"en-US": "Whenever is fine.",
	},
	COMMAND_TEXT_IN_FEW_MINUTES: map[string]string{
		"ru-RU": "Через несколько минут",
		"en-US": "In few minutes",
	},
	COMMAND_TEXT_IN_1_WEEK: map[string]string{
		"ru-RU": "Через неделю",
		"en-US": "In 1 week",
	},
	COMMAND_TEXT_IN_1_MONTH: map[string]string{
		"ru-RU": "Через месяц",
		"en-US": "In 1 month",
	},
	COMMAND_TEXT_SET_DATE: map[string]string{
		"ru-RU": "Задать дату",
		"en-US": "Set dte",
	},
	COMMAND_TEXT_MONDAY: map[string]string{
		"ru-RU": "Понедельник",
		"en-US": "Monday",
	},
	COMMAND_TEXT_TUESDAY: map[string]string{
		"ru-RU": "Вторник",
		"en-US": "Tuesday",
	},
	COMMAND_TEXT_WEDNESDAY: map[string]string{
		"ru-RU": "Среда",
		"en-US": "Wednesday",
	},
	COMMAND_TEXT_THURSDAY: map[string]string{
		"ru-RU": "Четверг",
		"en-US": "Thursday",
	},
	COMMAND_TEXT_FRIDAY: map[string]string{
		"ru-RU": "Пятница",
		"en-US": "Friday",
	},
	COMMAND_TEXT_SATURDAY: map[string]string{
		"ru-RU": "Суббота",
		"en-US": "Saturday",
	},
	COMMAND_TEXT_SUNDAY: map[string]string{
		"ru-RU": "Воскресенье",
		"en-US": "Sunday",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_TELEGRAM: map[string]string{
		"ru-RU": "Отправить через Telelgram",
		"en-US": "Send by Telegram",
	},
	COMMAND_TEXT_CANCEL_SENDING_RECEIPT_BY_TELEGRAM: map[string]string{
		"ru-RU": "Отменить отправку через Telelgram",
		"en-US": "Cancel sending receipt by Telegram",
	},
	COMMAND_TEXT_MAIN_MENU_TITLE: map[string]string{
		"ru-RU": "Главное /меню",
		"en-US": "Main /menu",
	},
	MESSAGE_TEXT_NOTHING_TO_CANCEL: map[string]string{
		"ru-RU": "Нечего отменять.",
		"en-US": "Nothing to cancel.",
	},
	MESSAGE_TEXT_TRANSFER_CREATION_CANCELED: map[string]string{
		"ru-RU": "Создание записи о долге отменено.",
		"en-US": "Creation of debt record has been canceled.",
	},
	COMMAND_TEXT_SHOW_ALL_CONTACTS: map[string]string{
		"ru-RU": "Показать все...",
		"en-US": "Show all...",
	},
	COMMAND_TEXT_ADD_YOUR_OWN_OPTION: map[string]string{
		"ru-RU": "Что-то другое",
		"en-US": "Something else",
	},
	"book": map[string]string{
		"ru-RU": "книгу",
		"en-US": "book",
	},
	bots.MESSAGE_TEXT_I_DID_NOT_UNDERSTAND_THE_COMMAND: map[string]string{
		"ru-RU": "\xF0\x9F\x98\xB3 Извините, я не понял вашу команду. Возможно я немного туповат...\n\nЧтобы начать сначала нажмите /меню",
		"en-US": "\xF0\x9F\x98\xB3 Sorry, I did not understand yoru order. May be I'm a little bit dumb...\n\nYou can return to main /menu",
	},
	"COMMAND_TEXT_LANGUAGE": map[string]string{
		"ru-RU": "/Язык приложения",
		"en-US": "App /language",
	},
	"/start": map[string]string{
		"ru-RU": "/старт",
		"en-US": "/start",
	},
	COMMAND_TEXT_NOTIFICATIONS: map[string]string{
		"ru-RU": "/Извещения",
		"en-US": "/Notifications",
	},
	COMMAND_TEXT_GAVE: map[string]string{
		"ru-RU": "/Дал",
		"en-US": "/Gave",
	},
	COMMAND_TEXT_GOT: map[string]string{
		"ru-RU": "/Взял",
		"en-US": "/Got",
	},
	COMMAND_TEXT_RETURN: map[string]string{
		"ru-RU": "/Возврат",
		"en-US": "/Return",
	},
	COMMAND_TEXT_BALANCE: map[string]string{
		"ru-RU": "/Баланс",
		"en-US": "/Balance",
	},
	COMMAND_TEXT_SETTING: map[string]string{
		"ru-RU": "/Настройки",
		"en-US": "/Settings",
	},
	COMMAND_TEXT_HELP: map[string]string{
		"ru-RU": "/Помощь",
		"en-US": "/Help",
	},
	COMMAND_TEXT_HISTORY: map[string]string{
		"ru-RU": "/История",
		"en-US": "/History",
	},
	COMMAND_TEXT_CANCEL: map[string]string{
		"ru-RU": "/Отменить",
		"en-US": "/Cancel",
	},
	COMMAND_TEXT_SETTINGS_PRIMARY_CURRENCY: map[string]string{
		"ru-RU": "Основная валюта",
		"en-US": "Primary currency",
	},
	COMMAND_TEXT_NEW_COUNTERPARTY: map[string]string{
		"ru-RU": "Добавить",
		"en-US": "Add new",
	},
	MESSAGE_TEXT_ASK_NEW_COUNTERPARTY_NAME: map[string]string{
		"ru-RU": "Имя для нового контакта:\n<i>(отправьте '.' для отмены)</i>",
		"en-US": "Please enter a name for the new contact:\n<i>(send '.' to cancel)</i>",
	},
	MESSAGE_TEXT_TRANSFER_IS_CREATING: map[string]string{
		"ru-RU": "Создаю запись...",
		"en-US": "Creating transfer...",
	},
	COMMAND_TEXT_SUBSCRIBE_TO_APP: map[string]string{
		"ru-RU": "Хочу приложение!",
		"en-US": "I want the app!",
	},
	COMMAND_TEXT_I_AM_FINE_WITH_BOT: map[string]string{
		"ru-RU": "Меня вполне устраивает бот!",
		"en-US": "I'm fine with just the bot!",
	},
	MESSAGE_TEXT_YOUR_AD_COULD_BE_HERE: map[string]string{
		"ru-RU": "Здесь можно <a href>разместить рекламу</a>",
		"en-US": "You can <a href>advertise here</a>",
	},
	MESSAGE_TEXT_YOUR_ABOUT_ADS: map[string]string{
		"ru-RU": `Здесь может быть ваша реклама если напишите нам на ads@debtstracker.io.
А пока что я порекламирую наше приложения для смартфонов.

🤖: Я конечно клёвый и обоятельный робот, но пользоваться специализированным приложением может быть даже удобнее. Иногда.

Оно ещё не готово для общего доступа, но уже сейчас можно посмотреть как оно будет выглядеть: https://debtstracker.io/#utm_source=telegram&utm_campaign=ads_screen

Хотите получить оповещение когда оно выйдет?`,
		"en-US": `Here could be your ad if you write us to ads@debtstracker.io. But for now we'll promote our own app for smartphones.

🤖: I'm a good robot, for sure, but sometimes it just more convinient to use a nice specialized app.

It's not ready for public preview just yet but you arleady can have check how it's going to looks: https://debtstracker.io/#utm_source=telegram&utm_campaign=ads_screen

Do you want to get an invite when it gets released?
`,
	},
	MESSAGE_TEXT_INVALID_INT: map[string]string{
		"ru-RU": "Извините, на данный момент я могу сохранять <b>только</b> целые числа (<i>1, 2, 3, ...</>) в качестве суммы/количества.",
		"en-US": "Sorry, at the moment I can <b>store</b> just integer numbers (<i>1, 2, 3, ...</>) в качестве суммы/количества.",
	},
	MESSAGE_TEXT_ASK_LENDING_TYPE: map[string]string{
		"ru-RU": "Что вы дали в долг?",
		"en-US": "What did you lend to someone?",
	},
	MESSAGE_TEXT_ASK_LENDING_AMOUNT: map[string]string{
		"ru-RU": "Сколько <b>%v</b> вы дали в долг?\n(<i>отправьте '.' чтобы отменить</i>)",
		"en-US": "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)",
	},
	MESSAGE_TEXT_ASK_LENDING_COUNTERPARTY: map[string]string{
		"ru-RU": "Кому вы дали в долг <b>%v</b>?\n(<i>отправьте '.' чтобы отменить</i>)",
		"en-US": "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
	},
	MESSAGE_TEXT_ASK_BORROWING_TYPE: map[string]string{
		"ru-RU": "Что вы взяли в долг?",
		"en-US": "What did you lend?",
	},
	MESSAGE_TEXT_ASK_BORROWING_AMOUNT: map[string]string{
		"ru-RU": "Сколько <b>%v</b> вы взяли в долг?\n(<i>отправьте '.' чтобы отменить</i>)",
		"en-US": "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)",
	},
	MESSAGE_TEXT_ASK_BORROWING_COUNTERPARTY: map[string]string{
		"ru-RU": "У кого вы взяли в долг <b>%v</b>?\n(<i>отправьте '.' чтобы отменить</i>)",
		"en-US": "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
	},
	MESSAGE_TEXT_YOU_CAN_SEND_RECEIPT: map[string]string{
		"ru-RU": "Отправить квитанцию для %v?",
		"en-US": "Should we send a receipt to %v?",
	},
	MESSAGE_TEXT_RECEIPT_IS_SENDING_BY_TELEGRAM: map[string]string{
		"ru-RU": "Отправляем для %v извещение через Telegram...",
		"en-US": "We are sending receipt to %v by Telegram...",
	},
	MESSAGE_TEXT_TRANSFER_FROM_USER_COMPLETED: map[string]string{
		"ru-RU": "%v взял(а) в долг %v.",
		"en-US": "%v borrowed from you %v.",
	},
	MESSAGE_TEXT_TRANSFER_TO_USER_COMPLETED: map[string]string{
		"ru-RU": "%v дал(а) вам в долг %v.",
		"en-US": "%v lended to you %v.",
	},
	MESSAGE_TEXT_DUE_ON: map[string]string{
		"ru-RU": "<b>Вернуть до</b>: %v",
		"en-US": "<b>Return till</b>: %v",
	},
	MESSAGE_TEXT_NOTE: map[string]string{
		"ru-RU": "Заметка",
		"en-US": "Note",
	},
	MESSAGE_TEXT_COMMENT: map[string]string{
		"ru-RU": "Комментарий",
		"en-US": "Comment",
	},
	MESSAGE_TEXT_SETTINGS: map[string]string{
		"ru-RU": "Что будем настраивать?",
		"en-US": "What do you want to adjust?",
	},
	MESSAGE_TEXT_NOT_IMPLEMENTED_YET: map[string]string{
		"ru-RU": "Извините, данный функционал ещё не реализован.",
		"en-US": "Sorry, this functionality is not implemented yet.",
	},
	MESSAGE_TEXT_ASK_INVITE_CHANNEL: map[string]string{
		"ru-RU": "Как вы хотите получить код приглашения?",
		"en-US": "How do you want to get an invite?",
	},
	MESSAGE_TEXT_PLEASE_ENTER_INVITE_CODE: map[string]string{
		"ru-RU": "Пожалуйста введите код приглашения:",
		"en-US": "Please enter an invite code:",
	},
	MESSAGE_TEXT_USER_EMAIL_FOR_INVITE_RECEIVED:  map[string]string{
		"ru-RU": "Мы отправили письмо на %v.\n\nПожалуйста откройте его и кликните на ссылку для подтверждения адреса.",
		"en-US": "We've sent a message to %v.\n\nPlease open the email and click a link to confirm your email address.",
	},
	MESSAGE_TEXT_USER_CONTACT_FOR_INVITE_RECEIVED: map[string]string{
		"ru-RU": "Спасибо, вы записаны в очередь на получение приглашения.\n\nТекущее время ожидания 2-3 дня.\n\nВы можете получить приглашение сегодня если расскажите о нашем боте на Facebook.",
		"en-US": "Thanks, you've been queued for an invite.\n\nCurrent awaiting time is 2-3 days.\n\nYou can get an invite code today by sharing a link to the bot on Facebook.",
	},
	MESSAGE_TEXT_PLEASE_PROVIDE_YOUR_EMAIL: map[string]string{
		"ru-RU": "Пожалуйста напишите ваш email адрес:",
		"en-US": "Please provide your email adress",
	},
	MESSAGE_TEXT_PLEASE_PROVIDE_YOUR_PHONE_NUMBER: map[string]string{
		"ru-RU": "Пожалуйста напишите номер вашего телефона:",
		"en-US": "Please provide your email adress",
	},
	MESSAGE_TEXT_WRONG_INVITE_CODE: map[string]string{
		"ru-RU": "Неправильный код приглашения: %v",
		"en-US": "Wrong invite code: %v",
	},
	MESSAGE_TEXT_WRONG_EMAIL: map[string]string{
		"ru-RU": "Неправильный email адрес.",
		"en-US": "Wrong email address.",
	},
	MESSAGE_TEXT_WRONG_PHONE_NUMBER: map[string]string{
		"ru-RU": "Неправильный номер телефона.",
		"en-US": "Wrong phone number.",
	},
	MESSAGE_TEXT_OK_PLEASE_TRY_AGAIN: map[string]string{
		"ru-RU": "Хорошо, попробуйте ещё раз.",
		"en-US": "Ok, please try again.",
	},

	COMMAND_TEXT_MISTYPE_WILL_TRY_AGAIN: map[string]string{
		"ru-RU": "Я опечатался, попробую ещё раз.",
		"en-US": "I've mistyped, will try again.",
	},
	COMMAND_TEXT_TELL_ME_MORE_ABOUT_INVITES: map[string]string{
		"ru-RU": "Расскажите ка мне об этих кодах",
		"en-US": "Tell me more about the codes",
	},
	COMMAND_TEXT_INVITE_ME_BY_EMAIL: map[string]string{
		"ru-RU": "Хочу код приглашения на email",
		"en-US": "Send me invite code by email",
	},
	COMMAND_TEXT_INVITE_ME_BY_SMS: map[string]string{
		"ru-RU": "Хочу код приглашения по SMS",
		"en-US": "Send me invite code by SMS",
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_EMAIL: map[string]string{
		"ru-RU": "Новый код приглашения на email",
		"en-US": "Send me new invite code by email",
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_SMS: map[string]string{
		"ru-RU": "Новый код приглашения по SMS",
		"en-US": "Send me new invite code by SMS",
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_TELEGRAM: map[string]string{
		"ru-RU": "Получить приграшение в Telegram",
		"en-US": "Send me new invite by Telegram",
	},
	MESSAGE_TEXT_UNKNOWN_LANGUAGE: map[string]string{
		"ru-RU": "Незнакомый язык. Пожалуйста выберете один из предоставленных:",
		"en-US": "Unknown language. Please choose 1 from the options:",
	},
	MESSAGE_TEXT_ONBOARDING_ASK_TO_CHOOSE_LANGUAGE: map[string]string{
		"ru-RU": `¡Hola! Hi! Привет! سلام!

<b>%v</b>, на каком языке вы хотели бы общаться?

(What is your prefered language?)`,
		"en-US": `¡Hola! Hi! Привет! سلام!

<b>%v</b>, what is your prefered language?`,
	},
	MESSAGE_TEXT_CHOOSE_UI_LANGUAGE: map[string]string{
		"ru-RU": "На каком языке вы хотели бы общаться со мной?",
		"en-US": "Which language you would like to talk to me?",
	},
	MESSAGE_TEXT_WHATS_NEXT: map[string]string{
		"ru-RU": "Что будем делать дальше?",
		"en-US": "What's next?",
	},
	"History": map[string]string{
		"ru-RU": "История",
		"en-US": "History",
	},
	MESSAGE_TEXT_HISTORY: map[string]string{
		"ru-RU": `<b>История</b> <i>(%d последних)</i>
─────────────
%v`,
		"en-US": `<b>History</b> <i>(last %d):</i>
─────────────
%v`,
	},
	MESSAGE_TEXT_BALANCE_IS_ZERO: map[string]string{
		"ru-RU": "Нет записей о текущих долгах.",
		"en-US": "You have no records on current debts.",
	},
	MESSAGE_TEXT_BALANCE_CURRENCY_TOTAL_INTRO: map[string]string{
		"ru-RU": "Всего",
		"en-US": "Total",
	},
	MESSAGE_TEXT_PRIMARY_CURRENCY_IS_SET_TO: map[string]string{
		"ru-RU": "OK, теперь я буду использовать '%v' как основную валюту.",
		"en-US": "OK, from now on I will use '%v' as a primary currency.",
	},
	MESSAGE_TEXT_BALANCE_SINGLE_CURRENCY_COUNTERPARTY_DEBT_TO_USER: map[string]string{
		"ru-RU": "<b>%v</b> - долг вам %v",
		"en-US": "<b>%v</b> - owes you %v",
	},
	MESSAGE_TEXT_BALANCE_CURRENCY_ROW_DEBT_TO_USER: map[string]string{
		"ru-RU": "Вам должны %v",
		"en-US": "Owes to you %v",
	},
	MESSAGE_TEXT_ON_RETURN_USER_DOES_NOT_OWE_ANYTHING_TO_COUNTERPARTY_ANYMORE: map[string]string{
		"ru-RU": "Поздравляем! У вас не осталось долгов перед <b>%v</b>.",
		"en-US": "Congratulations! You don't owe anything more to <b>%v</b>.",
	},
	MESSAGE_TEXT_ON_RETURN_COUNTERPARTY_DOES_NOT_OWE_ANYTHING_TO_USER_ANYMORE: map[string]string{
		"ru-RU": "У <b>%v</b> больше не осталось долгов перед вами.",
		"en-US": "<b>%v</b> does not owe anything more to you.",
	},
	MESSAGE_TEXT_BALANCE_CURRENCY_ROW_DEBT_BY_USER:map[string]string{
		"ru-RU": "Вы должны %v",
		"en-US": "You owe %v",
	},
	MESSAGE_TEXT_BALANCE_SINGLE_CURRENCY_COUNTERPARTY_DEBT_BY_USER: map[string]string{
		"ru-RU": "<b>%v</b> - вы должны %v",
		"en-US": "<b>%v</b> - you owe %v",
	},
	MESSAGE_TEXT_ASK_PRIMARY_CURRENCY: map[string]string{
		"ru-RU": "Какая валюта для вас основная?",
		"en-US": "What is your primary currency?",
	},
	MESSAGE_TEXT_FAILED_TO_DELETE_USER: map[string]string{
		"ru-RU": "Не удалось удалить данные пользователя: %v",
		"en-US": "Failed to delete user: %v",
	},
	MESSAGE_TEXT_USER_DELETED: map[string]string{
		"ru-RU": "Данные пользователя удалены",
		"en-US": "User's data has been deleted",
	},
	MESSAGE_TEXT_PLEASE_WAIT_WHILE_WE_GENERATE_INVITE_CODE: map[string]string{
		"ru-RU": "Пожалуйста подождите пока мы генерируем секретный код доступа...",
		"en-US": "Please wait a moment while we are generating a security access code...",
	},
	MESSAGE_TEXT_RETURN_ASK_TO_CHOOSE_COUNTERPARTY: map[string]string{
		"ru-RU": "Выберете кому вы вернули долг или кто вернул его вам.",
		"en-US": "Please choose who returned the debt or to who you returned it.",
	},
	MESSAGE_TEXT_CHOOSE_DEBT_THAT_HAS_BEEN_RETURNED: map[string]string{
		"ru-RU": "Выберите долг который был возвращён целиком или частично.",
		"en-US": "Please choose a debt that has been returned fully or partially.",
	},
	MESSAGE_TEXT_YOU_HAVE_RETURNED: map[string]string{
		"ru-RU": "Вы вернули %v",
		"en-US": "You've returned %v",
	},
	MESSAGE_TEXT_YOU_HAVE_GOT_BACK: map[string]string{
		"ru-RU": "Вам вернули %v",
		"en-US": "You've got back %v",
	},
	MESSAGE_TEXT_PLEASE_ACKNOWLEDGE_TRANSFER: map[string]string{
		"ru-RU": "Пожалуйста подтвердите или отклоните эту транзакцию.",
		"en-US": "Please confirm or decline this transfer.",
	},
	MESSAGE_TEXT_RECEIPT_LINK: map[string]string{
		"ru-RU": "Подробнее тут: %v",
		"en-US": "Details here: %v",
	},
	MESSAGE_TEXT_ASK_PHONE_NUMBER_OF_COUNTERPARTY: map[string]string{
		"ru-RU": "Пожалуйста напишите номер для <b>%v</b>",
		"en-US": "Plese provide phone number for <b>%v</b>",
	},
	MESSAGE_TEXT_ABOUT_PHONE_NUMBER_FORMAT: map[string]string{
		"ru-RU": "Номер должен быть в международном формате:\n\t* Начинаться со знака '+' и кода страны\n\t* Состоять только из цифр\nПример: <pre>+</pre><b>7</b><code>999012345678</code>",
		"en-US": "The number should be in internationl standard:\n\t* Starts with '+' following by country code\n\t* Consist of numbers only\nExample: <pre>+</pre><b>1</b><code>999012345678</code>",
	},
	MESSAGE_TEXT_THIS_NUMBER_WILL_BE_USED_TO_SEND_RECEIPT: map[string]string{
		"ru-RU": "На этот номер мы отправим SMS:",
		"en-US": "Will send an SMS to this number:",
	},
	MESSAGE_TEXT_COUNTERPARTY_OWES_YOU_SINGLE_DEBT_IS_IT_RETURNED_IN_FULL: map[string]string{
		"ru-RU": "<b>%v</b> одалживал(а) у вас <b>%v</b>. Возвращено полностью?",
		"en-US": "<b>%v</b> owed to you <b>%v</b>. Has this debt been returned in full?",
	},
	MESSAGE_TEXT_YOU_OWE_TO_COUNTERPARTY_SINGLE_DEBT_IS_IT_RETURNED_IN_FULL: map[string]string{
		"ru-RU": "<b>%v</b> одалживал(а) вам <b>%v</b>. Возвращено полностью?",
		"en-US": "You owe to <b>%v</b> <b>%v</b>. Has this debt been returned in full?",
	},
	BUTTON_TEXT_YOU_OWE_AMOUNT_TO_SOMEONE: map[string]string{
		"ru-RU": "%v | вы должны: %v",
		"en-US": "%v | you owe: %v",
	},
	BUTTON_TEXT_SOMEONE_OWES_TO_YOU_AMOUNT: map[string]string{
		"ru-RU": "%v | долг вам: %v",
		"en-US": "%v | owes to you: %v",
	},
	BUTTON_TEXT_DEBT_RETURNED_FULLY: map[string]string{
		"ru-RU": "Да, целиком",
		"en-US": "Yes, fully",
	},
	BUTTON_TEXT_DEBT_RETURNED_PARTIALLY: map[string]string{
		"ru-RU": "Нет, только часть",
		"en-US": "No, just partially",
	},
	MESSAGE_TEXT_ATTEMPT_TO_USE_OWN_INVITE: map[string]string{
		"ru-RU": "Хорошая попытка пригласить самого себя ;)",
		"en-US": "You should not use your own invite ;)",
	},
	BUTTON_TEXT_SEE_RECEIPT_DETAILS:  map[string]string{
		"ru-RU": "Показать детали",
		"en-US": "Show receipt details ",
	},
	MESSAGE_TEXT_ABOUT_INVITES: map[string]string{
		"ru-RU": `На данный момент доступ к нашему боту ограничен, но вы можете пригласить друга.

Как вы хотите передать код приглашение?`,
		"en-US": `At the moment access to our bot is limited but you can invite your friend.

How do you want to pass the invite code?`,
	},
	MESSAGE_TEXT_USER_BLOCKED_TRANSFER_NOTIFICATIONS_BY: map[string]string{
		"ru-RU": "%v заблокировал получение оповешений о транзакиях через: %v.",
		"en-US": "%v blocked notifications about transactions by: %v",
	},
	COMMAND_TEXT_ACCEPT: map[string]string{
		"ru-RU": "Согласиться",
		"en-US": "Accept",
	},
	BUTTON_TEXT_ACCEPT_TRANSFER_USING_TELEGRAM: map[string]string{
		"ru-RU": "Подтвердить (используя Telegram)",
		"en-US": "Accept (using Telegram messenger)",
	},
	BUTTON_TEXT_DECLINE_TRANSFER_USING_TELEGRAM: map[string]string{
		"ru-RU": "Отказаться (используя Telegram)",
		"en-US": "Decline (using Telegram messenger)",
	},
	COMMAND_TEXT_DECLINE: map[string]string{
		"ru-RU": "Отклонить",
		"en-US": "Decline",
	},
	COMMAND_TEXT_ACCEPT_INVITE: map[string]string{
		"ru-RU": "Принять приглашение",
		"en-US": "Accept invite",
	},
	COMMAND_TEXT_SEE_RECEIPT_DETAILS: map[string]string{
		"ru-RU": "Посмотреть квитанцию",
		"en-US": "See receipt details",
	},
	COMMAND_TEXT_OTHER_WAYS_TO_SEND_INVITE: map[string]string{
		"ru-RU": "Другие способы отправить приглашение",
		"en-US": "Other ways to send an invite",
	},
	COMMAND_TEXT_SEND_MY_PHONE_NUMBER: map[string]string{
		"ru-RU": "Отправить мой номер",
		"en-US": "Send my phone number",
	},
	COMMAND_TEXT_SEND_BY_EMAIL: map[string]string{
		"ru-RU": "Через Email",
		"en-US": "By Email",
	},
	COMMAND_TEXT_SEND_BY_SMS: map[string]string{
		"ru-RU": "Через SMS",
		"en-US": "By SMS",
	},
	COMMAND_TEXT_INVITE_BY_TELEGRAM: map[string]string{
		"ru-RU": "Пригласить через Telegram",
		"en-US": "InviteBy Telegram",
	},
	MESSAGE_TEXT_INVITE_CREATED: map[string]string{
		"ru-RU": "Код приглашения: %v",
		"en-US": "Invite code: %v",
	},
	MESSAGE_TEXT_INVITE_BY_EMAIL: map[string]string{
		"ru-RU": "Введите email вашего друга на который мы отправим код приглашения.",
		"en-US": "Please enter emaill address of your friend where we should send an invite code.",
	},
	MESSAGE_TEXT_INVITE_ASK_EMAIL_FOR_RECEIPT: map[string]string{
		"ru-RU": "Введите email вашего друга (%v) на который мы отправим квитанцию подтверждения.",
		"en-US": "Please enter emaill address of your friend (%v) where we should send the receipt.",
	},
	MESSAGE_TEXT_INVITE_BY_SMS: map[string]string{
		"ru-RU": "Введите номер телефона вашего друга на который мы отправим код приглашения.",
		"en-US": "Please share a contact or enter phone number of your friend where we should send an invite code.",
	},
	MESSAGE_TEXT_INVITE_BY_TELEGRAM: map[string]string{
		"ru-RU": "Вставьте в чат контакт вашего друга которому вы хотите отправить приглашение.",
		"en-US": "Please share a contact of your friend you wish to send an invite code:",
	},
	MESSAGE_TEXT_INVALID_EMAIL: map[string]string{
		"ru-RU": "Неверный email. Проверьте и попробуйте ещё раз? /меню",
		"en-US": "Invalid email. Check and try it again? /menu",
	},
	MESSAGE_TEXT_INVALID_PHONE_NUMBER: map[string]string{
		"ru-RU": "Неверный номер. Проверьте и попробуйте ещё раз? /меню",
		"en-US": "Invalid phone number. Check and try it again? /menu",
	},
	MESSAGE_TEXT_NO_CONTACT_RECEIVED: map[string]string{
		"ru-RU": "Мы не получили контакта. Тут инструкция как это сделать. /меню",
		"en-US": "We have not received any contacts. INSTRUCTION HOW TO DO IT. /menu",
	},
	MESSAGE_TEXT_HISTORY_ROW_FROM_USER_WITH_NAME: map[string]string{
		"ru-RU": "Вы ⇒ %s ⇒ %s",
		"en-US": "You ⇒ %s ⇒ %s",
	},
	MESSAGE_TEXT_BALANCE_HEADER: map[string]string{
		"ru-RU": "Баланс",
		"en-US": "Balance",
	},
	MESSAGE_TEXT_RECEIPT_SENT_THROW_TELEGRAM: map[string]string{
		"ru-RU": "Квитанция отправлена через телеграм",
		"en-US": "Receipt sent through Telegram",
	},
	MESSAGE_TEXT_RECEIPT_VIEWED_BY_COUNTERPARTY: map[string]string{
		"ru-RU": "Квитанция просмотрена",
		"en-US": "Receipt viewed",
	},
	INLINE_BUTTON_SHOW_FULL_HISTORY: map[string]string{
		"ru-RU": "Показать всю историю",
		"en-US": "Show full history",
	},
	MESSAGE_TEXT_INCORRECT_VALUE_NOT_A_NUMBER: map[string]string{
		"ru-RU": "Это не цифра",
		"en-US": "it is not a number",
	},
	MESSAGE_TEXT_INCORRECT_VALUE_IS_NEGATIVE: map[string]string{
		"ru-RU": "Цифра должна быть положительной (<i>больше нуля</i>)",
		"en-US": "The number should be positive (<i>greater than 0</i>)",
	},
	MESSAGE_TEXT_ASK_HOW_MUCH_HAS_BEEN_RETURNED: map[string]string{
		"ru-RU": "Сколько было возвращено?",
		"en-US": "How much have been returned?",
	},
	MESSAGE_TEXT_HISTORY_ROW_TO_USER_WITH_NAME: map[string]string{
		"ru-RU": "%s ⇒ %s ⇒ Вам",
		"en-US": "%s ⇒ %s ⇒ to you",
	},
	MESSAGE_TEXT_WELCOME: map[string]string{
		"ru-RU": `Привет, я Коллектиус - Ваш персональный счетовод и коллектор.

Могу записать кто кому чего должен и, и при необходимости, напомнить когда должок пора возвращать.

Чего изволит новый хозяин?`,
		"en-US": `Hi, I'm Collectius - your personal accountant & collector.

I can record who is owing to whom and remind when the return is due.

What would you like to do?`,
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE: map[string]string{
		"ru-RU": "Хочу получить приглашение",
		"en-US": "I want to get an invite",
	},
	COMMAND_TEXT_I_HAVE_INVITE: map[string]string{
		"ru-RU": "У меня есть код приглашения",
		"en-US": "I have the invitation code",
	},
	MESSAGE_TEXT_ONBOARDING_TELL_ABOUT_INVITES: map[string]string{
		"ru-RU": `<b>%v</b>,

На данный момент наш бот доступен только тем кто получил приглашение от друзей.

Если у вас нет кода можете оставить свои контакты и мы вышлем вам приглашение как только наступит ваша очередь.

Мы высылаем по 10 кодов в день первоочередникам + 1 случайным образом.`,
		"en-US": `<b>%v</b>,

At the momnet our bot is available just by invitation from friends.

If you have no code you can leave you contact details and we'll send you an invite as soon as your queue is due.

We send 10 invites per day to those in the head of the queue and 1 randomly.`,
	},
	EMAIL_INVITE_SUBJ: map[string]string{
		"ru-RU": "Приглашение от {{.FromName}}",
		"en-US": "An invite from {{.FromName}}",
	},
	EMAIL_INVITE_TEXT: map[string]string{
		"ru-RU": `Привет {{.ToName}},


{{.FromName}} приглашает тебя попробовать приложение для учёта долгов - https://debtstracker.io/#invite={{.InviteCode}}

Ваш код приглашения: {{.InviteCode}}`,
		"en-US": `Hi{{.ToName}},
{{.FromName}} п is inviting you to use debts tracking app - https://debtstracker.io/#invite={{.InviteCode}}

You invitation code is: {{.InviteCode}}`,
	},
	EMAIL_INVITE_HTML: map[string]string{
		"ru-RU": "",
		"en-US": "",
	},
	EMAIL_RECEIPT_SUBJ: map[string]string{
		"ru-RU": "Запись о долге - {{.FromName}}",
		"en-US": "Debt record - {{.FromName}}",
	},
	EMAIL_RECEIPT_BODY_TEXT: map[string]string{
		"ru-RU": "{{.FromName}} создал запись о долге: {{.ReceiptURL}}",
		"en-US": "{{.FromName}} created a debt record: {{.ReceiptURL}}",
	},
	INLINE_RECEIPT_TITLE: map[string]string{
		"ru-RU": "Квитанция: %v",
		"en-US": "Receipt: %v",
	},
	INLINE_RECEIPT_DESCRIPTION: map[string]string{
		"ru-RU": "Нажмите здесь чтобы отправить квитанцию",
		"en-US": "Click here to send the receipt",
	},
	INLINE_RECEIPT_MESSAGE: map[string]string{
		"ru-RU": "%v создал(а) запись о долге.\n\n<i>Выберите язык чтобы посмотреть подробности.</i>",
		"en-US": "%v recorded a debt associated with you.\n\n<i>Choose language to see details.</i>",
	},
	INLINE_INVITE_TITLE: map[string]string{
		"ru-RU": "Приглашение в %v",
		"en-US": "Invitation to %v",
	},
	INLINE_INVITE_DESCRIPTION: map[string]string{
		"ru-RU": "Нажмите здесь для отправки приглашения",
		"en-US": "Click here to send an invite",
	},
	INLINE_INVITE_MESSAGE: map[string]string{
		"ru-RU": "%v пригласил вас попробовать %v",
		"en-US": "%v invited you to try %v",
	},
	LINK_VIEW_RECEIPT: map[string]string{
		"ru-RU": "Посмотреть квитанцию",
		"en-US": "View receipt online",
	},
	SMS_RECEIPT_YOU_GOT: map[string]string{
		"ru-RU": "Вы получили %v от %v. Подробнее тут: %v",
		"en-US": "You've got %v from %v. Details here: %v",
	},
	SMS_RECEIPT_YOU_GAVE: map[string]string{
		"ru-RU": "Вы дали %v - взял %v. Подробнее тут: %v",
		"en-US": "You've given %v to %v. Details here: %v",
	},
	HTML_RECEIPT: map[string]string{
		"ru-RU": "Квитанция",
		"en-US": "Receipt",
	},
	HTML_AMOUNT: map[string]string{
		"ru-RU": "Сумма",
		"en-US": "Amount",
	},
	HTML_FROM: map[string]string{
		"ru-RU": "Дал",
		"en-US": "From",
	},
	HTML_TO: map[string]string{
		"ru-RU": "Получил",
		"en-US": "To",
	},
	TELEGRAM_RECEIPT: map[string]string{
		"ru-RU": "{{.FromName}} создал запись о долге ({{.TransferCurrency}})",
		"en-US": "{{.FromName}} created a debtrecord ({{.TransferCurrency}})",
	},
	MESSAGE_TEXT_PLEASE_CHOOSE_FROM_OPTIONS_PROVIDED: map[string]string{
		"ru-RU": "Пожалуйста выберете из предоставленных опций.",
		"en-US": "Please choose from provided options.",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_NOTE_OR_COMMENT: map[string]string{
		"ru-RU": "<b>Хотите добавить заметку или комментарий?</b>\n%v Заметки хранятся для вашего личго пользования.\n%v Комментарий виден всем кому разрешён просмотр этой транзакции.",
		"en-US": "<b>Do you want to add a note or comment?</b>\n%v Memos are private records for yoru own reference.\n%v Comments are available to everyone who has permission to view this transaction.",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_NOTE: map[string]string{
		"ru-RU": "Напишите заметку:",
		"en-US": "Please write a note:",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_COMMENT: map[string]string{
		"ru-RU": "Напишите комментарий:",
		"en-US": "Please write the comment:",
	},
	MESSAGE_TEXT_TRANSFER_NOTE_ADDED_ASK_FOR_COMMENT: map[string]string{
		"ru-RU": "Заметка добавлена. Хотите написать комментарий?",
		"en-US": "Memo have been added. Do you want to write a comment?",
	},
	MESSAGE_TEXT_TRANSFER_COMMENT_ADDED_ASK_FOR_NOTE: map[string]string{
		"ru-RU": "Комментарий добавлен. Хотите написать заметку?",
		"en-US": "Comment have been added. Do you want to write a note?",
	},
	COMMAND_TEXT_NO_COMMENT_OR_NOTE_FOR_TRANSFER: map[string]string{
		"ru-RU": "Без заметок и комментариев",
		"en-US": "Without notes or comments",
	},
	COMMAND_TEXT_NO_COMMENT_FOR_TRANSFER: map[string]string{
		"ru-RU": "Без комментариев",
		"en-US": "No comments",
	},
	COMMAND_TEXT_NO_NOTE_FOR_TRANSFER: map[string]string{
		"ru-RU": "Без заметок",
		"en-US": "Without notes",
	},
	COMMAND_TEXT_ADD_NOTE_TO_TRANSFER: map[string]string{
		"ru-RU": "Добавить заметку",
		"en-US": "Add a note (private)",
	},
	COMMAND_TEXT_ADD_COMMENT_TO_TRANSFER: map[string]string{
		"ru-RU": "Добавить комментарий",
		"en-US": "Add a comment (public)",
	},
}
