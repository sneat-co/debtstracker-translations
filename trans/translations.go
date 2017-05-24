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
		"fa-IR": "نمونه",
	},
	"Jan": {
		"en-US": "Jan",
		"ru-RU": "Янв.",
		"fa-IR": "ژانویه",
	},
	"Feb": {
		"en-US": "Feb",
		"ru-RU": "Фев.",
		"fa-IR": "فوریه",
	},
	"Mar": {
		"en-US": "Mar",
		"ru-RU": "Мрт.",
		"fa-IR": "مارس",
	},
	"Apr": {
		"en-US": "Apr",
		"ru-RU": "Апр.",
		"fa-IR": "آوریل",
	},
	"May": {
		"en-US": "May",
		"ru-RU": "Май ",
		"fa-IR": "مه",
	},
	"Jun": {
		"en-US": "Jun",
		"ru-RU": "Июнь",
		"fa-IR": "ژوئن",
	},
	"Jul": {
		"en-US": "Jul",
		"ru-RU": "Июль",
		"fa-IR": "ژوئیه",
	},
	"Aug": {
		"en-US": "Aug",
		"ru-RU": "Авг.",
		"fa-IR": "اوت",
	},
	"Sep": {
		"en-US": "Sep",
		"ru-RU": "Сен.",
		"fa-IR": "سپتامبر",
	},
	"Oct": {
		"en-US": "Oct",
		"ru-RU": "Окт.",
		"fa-IR": "اکتبر",
	},
	"Nov": {
		"en-US": "Nov",
		"ru-RU": "Нбр.",
		"fa-IR": "نوامبر",
	},
	"Dec": {
		"en-US": "Dec",
		"ru-RU": "Дек.",
		"fa-IR": "دسامبر",
	},

	adsCommandTitle: {
		"ru-RU": adsCommandTitle,
		"en-US": adsCommandTitle,
		"fa-IR": adsCommandTitle,
	},
	" and ": {
		"ru-RU": " и ",
		"en-US": " and ",
		"fa-IR": " و ",
	},
	"MESSAGE_TEXT_OOPS_SOMETHING_WENT_WRONG": {
		"ru-RU": "Упс, что-то пошло не так... \xF0\x9F\x98\xB3",
		"en-US": "Oops, something went wrong... \xF0\x9F\x98\xB3",
		"fa-IR": "اوه، یک جای کار مشکل دارد ...  \xF0\x9F\x98\xB3",
	},
	MESSAGE_TEXT_ASK_DUE: {
		"ru-RU": "Когда дата возврата?",
		"en-US": "When is the due date?",
		"fa-IR": "سررسید چه زمانی است؟",
	},
	MESSAGE_TEXT_ASK_DATE_TO_REMIND: {
		"ru-RU": `Чтобы задать дату напопинания напишите её в формате <i>ДД.MM.ГОД</i>.
<b>Например</b> для 20 января 2017 г. отправьте:
    <i>20.01.2017</i>`,

		"en-US": `To set date for next reminder please send it as a text in format of <i>DD.MM.YEAR</i>.
<b>For example</b> for 20th of January 2017 submit:
    <i>20.01.2017</i>`,
		"fa-IR": `لطفاً برای تنظیم تاریخ یادآوری بعدی این فرمت را رعایت فرمایید <i>DD.MM.YEAR</i>.
<i>روز.ماه.سال</i>.
<b>برای مثال</b> برای تاریخ 20 ژانویه 2017 ثبت کنید:
    <i>20.01.2017</i>`,
	},
	MESSAGE_TEXT_ASK_DUE_DATE: {
		"ru-RU": `Чтобы задать дату возврата напишите её в формате <i>ДД.MM.ГОД</i>.
<b>Например</b> для 20 января 2017 г. отправьте:
    <i>20.01.2017</i>`,

		"en-US": `To set due date please send it as a text in format of <i>DD.MM.YEAR</i>.
<b>For example</b> for 20th of January 2017 submit:
    <i>20.01.2017</i>`,
		"fa-IR": `لطفاً برای تنظیم تاریخ سررسید این فرمت را رعایت فرمایید. <i>روز.ماه.سال</i>.
<b>برای مثال</b> برای 20 ژانویه 2017 ثبت کنید:
    <i>20.01.2017</i>`,

	},
	MESSAGE_TEXT_WRONG_DATE: {
		"ru-RU": "Извините, что-то не так с датой которую вы отправили.",
		"en-US": "Sorry, there is something wrong with the date you've provided.",
		"fa-IR": "متاسفم، در تاریخی که وارد نمودید مشکلی وجود دارد.",
	},
	COMMAND_TEXT_DISABLE_REMINDER: {
		"ru-RU": "Не напоминать",
		"en-US": "No reminder",
		"fa-IR": "بدون یادآور",
	},
	COMMAND_TEXT_TOMORROW: {
		"ru-RU": "Завтра",
		"en-US": "Tomorrow",
		"fa-IR": "فردا",
	},
	COMMAND_TEXT_DAY_AFTER_TOMORROW: {
		"ru-RU": "Послезавтра",
		"en-US": "Day after tomorrow",
		"fa-IR": "پس فردا",
	},
	COMMAND_TEXT_THIS_WEEK: {
		"ru-RU": "На этой неделе",
		"en-US": "This week",
		"fa-IR": "این هفته",
	},
	COMMAND_TEXT_YES_IT_HAS_RETURN_DEADLINE: {
		"ru-RU": "Да, есть срок возврата!",
		"en-US": "Yes, it has a deadline!",
		"fa-IR": "بله، دارای آخرین فرصت می باشد!",
	},
	COMMAND_TEXT_NO_IT_CAN_BE_RETURNED_ANYTIME: {
		"ru-RU": "Нет, срок возврата не важен.",
		"en-US": "No, whenever is fine.",
		"fa-IR": "خیر، هر زمانی مناسب است.",
	},
	COMMAND_TEXT_IT_CAN_BE_RETURNED_ANYTIME: {
		"ru-RU": "Когда-нибудь",
		"en-US": "Whenever is fine",
		"fa-IR": "هر زمانی مناسب است.",
	},
	COMMAND_TEXT_IN_FEW_MINUTES: {
		"ru-RU": "Через минуту",
		"en-US": "In few minutes",
		"fa-IR": "در چند دقیقه",
	},
	COMMAND_TEXT_IN_1_WEEK: {
		"ru-RU": "Через неделю",
		"en-US": "In 1 week",
		"fa-IR": "ظرف یک هفته",
	},
	COMMAND_TEXT_IN_1_MONTH: {
		"ru-RU": "Через месяц",
		"en-US": "In 1 month",
		"fa-IR": "ظرف یک ماه",
	},
	COMMAND_TEXT_SET_DATE: {
		"ru-RU": "Задать дату",
		"en-US": "Set date",
		"fa-IR": "تاریخ را تنظیم کنید",
	},
	COMMAND_TEXT_MONDAY: {
		"ru-RU": "Понедельник",
		"en-US": "Monday",
		"fa-IR": "دوشنبه",
	},
	COMMAND_TEXT_TUESDAY: {
		"ru-RU": "Вторник",
		"en-US": "Tuesday",
		"fa-IR": "سه شنبه",
	},
	COMMAND_TEXT_WEDNESDAY: {
		"ru-RU": "Среда",
		"en-US": "Wednesday",
		"fa-IR": "چهارشنبه",
	},
	COMMAND_TEXT_THURSDAY: {
		"ru-RU": "Четверг",
		"en-US": "Thursday",
		"fa-IR": "پنج شنبه",
	},
	COMMAND_TEXT_FRIDAY: {
		"ru-RU": "Пятница",
		"en-US": "Friday",
		"fa-IR": "جمعه",
	},
	COMMAND_TEXT_SATURDAY: {
		"ru-RU": "Суббота",
		"en-US": "Saturday",
		"fa-IR": "شنبه",
	},
	COMMAND_TEXT_SUNDAY: {
		"ru-RU": "Воскресенье",
		"en-US": "Sunday",
		"fa-IR": "یکشنبه",
	},
	COMMAND_TEXT_DO_NOT_SEND_RECEIPT: {
		"ru-RU": "Не отправлять квитанцию",
		"en-US": "Do not send the receipt",
		"fa-IR": "رسید را ارسال نکنید",
	},
	MESSAGE_TEXT_RECEIPT_WILL_NOT_BE_SENT: {
		"ru-RU": "Вы решили не отправлять квитанцию.",
		"en-US": "You've decided not to send the receipt.",
		"fa-IR": "شما تصمیم گرفتید که رسید را ارسال نکنید.",
	},
	COMMAND_TEXT_I_HAVE_CHANGED_MY_MIND: {
		"ru-RU": "Я передумал(а)",
		"en-US": "I've changed my mind",
		"fa-IR": "نظرم را عوض کردم.",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_TELEGRAM: {
		"ru-RU": "Отправить через Telelgram",
		"en-US": "Send by Telegram",
		"fa-IR": "با تلگرام ارسال شود",
	},
	COMMAND_TEXT_COUNTERPARTY_HAS_NO_TELEGRAM: {
		"ru-RU": "Отправить через Viber, VK, FB, ...",
		"en-US": "Send by FB, WhatsApp, Viber, etc.",
		"fa-IR": "با فیسبوک، واتس آپ، وایبر و ... ارسال شود.",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_SMS: {
		"ru-RU": "Отправить через SMS",
		"en-US": "Send by SMS",
		"fa-IR": "با پیام کوتاه ارسال شود",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_VK: {
		"ru-RU": "Отправить через ВКонтакте",
		"en-US": "Send throw VK.com",
		"fa-IR": "ارسال شود VK.com از طریق ",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_OK: {
		"ru-RU": "Отправить через Одноклассники",
		"en-US": "Send throw OK",
		"fa-IR": "ارسال شود OK از طریق ",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_FB: {
		"ru-RU": "Отправить через Facebook",
		"en-US": "Send throw Facebook",
		"fa-IR": "از طریق فیسبوک ارسال شود.",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_TWT: {
		"ru-RU": "Отправить через Twitter",
		"en-US": "Send throw Twiter",
		"fa-IR": "از طریق توئیتر ارسال شود.",
	},
	COMMAND_TEXT_CANCEL_SENDING_RECEIPT_BY_TELEGRAM: {
		"ru-RU": "Отменить отправку через Telelgram",
		"en-US": "Cancel sending receipt by Telegram",
		"fa-IR": "ارسال رسید با تلگرام کنسل شود",
	},
	COMMAND_TEXT_MAIN_MENU_TITLE: {
		"ru-RU": "Главное /меню",
		"en-US": "Main /menu",
		"fa-IR": "منو / اصلی ",
	},
	MESSAGE_TEXT_NOTHING_TO_CANCEL: {
		"ru-RU": "Нечего отменять.",
		"en-US": "Nothing to cancel.",
		"fa-IR": "چیزی برای کنسل شدن وجود ندارد",
	},
	MESSAGE_TEXT_TRANSFER_CREATION_CANCELED: {
		"ru-RU": "Создание записи о долге отменено.",
		"en-US": "Creation of debt record has been canceled.",
		"fa-IR": "ایجاد سابقه بدهی کنسل شده است.",
	},
	COMMAND_TEXT_SHOW_ALL_CONTACTS: {
		"ru-RU": "Показать все...",
		"en-US": "Show all...",
		"fa-IR": "نمایش تمام موارد ...",
	},
	COMMAND_TEXT_ADD_YOUR_OWN_OPTION: {
		"ru-RU": "Что-то другое",
		"en-US": "Something else",
		"fa-IR": "چیزی دیگر",
	},
	MESSAGE_TEXT_REMINDER_ASK_IF_RETURNED: {
		"ru-RU": "Был ли этот долг возвращён?",
		"en-US": "Have this debt been returned?",
		"fa-IR": "آیا این بدهی بازپرداخت شده است؟",
	},
	MESSAGE_TEXT_ASK_WHEN_TO_REMIND_AGAIN: {
		"ru-RU": "Когда вам напомнить об этом долге ещё раз?",
		"en-US": "When should we remind you about this debt again?",
		"fa-IR": "چه زمانی لازم است مجدداً در مورد این بدهی به شما یادآوری نماییم؟",
	},
	MESSAGE_TEXT_REPLIED_DEBT_RETURNED_FULLY: {
		"ru-RU": "Вы ответили что долг возвращён полностью.",
		"en-US": "You've replied that debt has been returned fully.",
		"fa-IR": "شما پاسخ داده اید که بدهی به صورت کامل بازپرداخت شده است.",
	},
	MESSAGE_TEXT_DEBT_IS_RETURNED: {
		"ru-RU": "Долг возвращён полностью.",
		"en-US": "The debt has been returned fully.",
		"fa-IR": "بدهی به صورت کامل بازپرداخت شده است",
	},
	MESSAGE_TEXT_DETAILS_ARE_HERE: {
		"ru-RU": "Подробности тут: %v",
		"en-US": "Details here: %v",
		"fa-IR": "جزئیات در اینجا: %v",
	},
	MESSAGE_TEXT_REMINDER: {
		"ru-RU": "Напоминание",
		"en-US": "Reminder",
		"fa-IR": "یادآور",
	},
	MESSAGE_TEXT_REMINDER_SET: {
		"ru-RU": "Напоминание установлено на: %v",
		"en-US": "Reminder set for: %v",
		"fa-IR": "یادآور تنظیم شده است برای: %v",
	},
	MESSAGE_TEXT_REMINDER_DISABLED: {
		"ru-RU": "Напоминания об этом долге отключены.",
		"en-US": "You've disabled reminders for this debt.",
		"fa-IR": "شما یادآور را برای این بدهی غیرفعال نموده اید.",
	},
	MESSAGE_TEXT_REMINDER_ALREADY_RESCHEDULED: {
		"ru-RU": "Напоминание об этом долге уже перенесено.",
		"en-US": "You've already rescheduled this reminder.",
		"fa-IR": "شما قبلا به صورت مجدد این یادآور را زمانبندی نموده اید.",
	},
	COMMAND_TEXT_REMINDER_RETURNED_IN_FULL: {
		"ru-RU": "Да, возвращено полностью",
		"en-US": "Yes, returne in full",
		"fa-IR": "بله، بازپرداخت به صورت کامل",
	},
	COMMAND_TEXT_REMINDER_RETURNED_PARTIALLY: {
		"ru-RU": "Возврашено частично",
		"en-US": "Returned partially",
		"fa-IR": "تا اندازه ای بازپرداخت شده است",
	},
	COMMAND_TEXT_REMINDER_NOT_RETURNED: {
		"ru-RU": "Не возвращено",
		"en-US": "Not returned",
		"fa-IR": "بازپرداخت نشده است",
	},
	MESSAGE_TEXT_YOU_REPLIED: {
		"ru-RU": "Вы ответили: %v",
		"en-US": "You've replied: %v",
		"fa-IR": "شما پاسخ داده اید: %v",
	},
	"book": {
		"ru-RU": "книгу",
		"en-US": "book",
		"fa-IR": "کتاب",
	},
	"MESSAGE_TEXT_I_DID_NOT_UNDERSTAND_THE_COMMAND": {
		"ru-RU": "\xF0\x9F\x98\xB3 Извините, я не понял вашу команду. Возможно я немного туповат...\n\nЧтобы начать сначала нажмите /menu",
		"en-US": "\xF0\x9F\x98\xB3 Sorry, I did not understand your order. May be I'm a little bit dumb...\n\nYou can return to main /menu",
		"fa-IR": "\xF0\x9F\x98\xB3 ببخشید، من دستور شما را نفهمیدم. احتمالا کمی کند ذهن هستم...\n\nشما میتوانید به منو /اصلی بازگردید",
	},
	"COMMAND_TEXT_LANGUAGE": {
		"ru-RU": "/Язык приложения",
		"en-US": "App /language",
		"fa-IR": "App /زبان",
	},
	"/start": {
		"ru-RU": "/старт",
		"en-US": "/start",
		"fa-IR": "/شروع",
	},
	COMMAND_TEXT_DUE_RETURNS: {
		"ru-RU": "Предстоящие платежи",
		"en-US": "Due returns",
		"fa-IR": "بازپرداخت بدهی",
	},
	MESSAGE_TEXT_OVERDUE_RETURNS_HEADER: {
		"ru-RU": "<b>Просроченные долги:</b>",
		"en-US": "<b>Overdue debts:</b>",
		"fa-IR": "<b>بدهی های معوق:</b>",
	},
	MESSAGE_TEXT_DUE_RETURNS_HEADER: {
		"ru-RU": "<b>Ближайшие долги к возрату:</b>",
		"en-US": "<b>Closest debts to return:</b>",
		"fa-IR": "<b>نزدیک ترین بدهی برای بازپرداخت:</b>",
	},
	MESSAGE_TEXT_DUE_RETURNS_ROW_BY_USER: {
		"ru-RU": "%v ожидает от вас возврата %v через %v",
		"en-US": "%v expects %v from you in %v",
		"fa-IR": "%v انتظار دارد %v از شما در %v",
	},
	MESSAGE_TEXT_DUE_RETURNS_ROW_BY_COUNTERPARTY: {
		"ru-RU": "Вы ожидаете от %v возврата %v через %v",
		"en-US": "You expect %v to retun %v to you in %v",
		"fa-IR": "شما انتظار دارید %v بازگرداند %v به شما در %v",
	},

	MESSAGE_TEXT_DUE_RETURNS_EMPTY: {
		"ru-RU": "У вас нет долгов с назначеным сроком к возврату.",
		"en-US": "You have no debts with set due date.",
		"fa-IR": "شما بدهی ای با ثبت سررسید ندارید.",
	},
	COMMAND_TEXT_GAVE: {
		"ru-RU": "/Дал",
		"en-US": "/Gave",
		"fa-IR": "/داد",
	},
	COMMAND_TEXT_GOT: {
		"ru-RU": "/Взял",
		"en-US": "/Got",
		"fa-IR": "/گرفت",
	},
	COMMAND_TEXT_RETURN: {
		"ru-RU": "/Вернул",
		"en-US": "/Returned",
		"fa-IR": "/بازگشت",
	},
	COMMAND_TEXT_BALANCE: {
		"ru-RU": "/Баланс",
		"en-US": "/Balance",
		"fa-IR": "/تراز",
	},
	COMMAND_TEXT_SETTING: {
		"ru-RU": "/Настройки",
		"en-US": "/Settings",
		"fa-IR": "/تنظیمات",
	},
	COMMAND_TEXT_HIGH_FIVE: {
		"ru-RU": "Дать пять!",
		"en-US": "High five!",
		"fa-IR": "بزن قدش!",
	},
	COMMAND_TEXT_CHANGE_LANG: {
		"ru-RU": "/Язык",
		"en-US": "/Language",
		"fa-IR": "/زبان",
	},
	COMMAND_TEXT_HELP: {
		"ru-RU": "/Помощь",
		"en-US": "/Help",
		"fa-IR": "/کمک",
	},
	COMMAND_TEXT_HISTORY: {
		"ru-RU": "/История",
		"en-US": "/History",
		"fa-IR": "/پیشینه",
	},
	COMMAND_TEXT_CANCEL: {
		"ru-RU": "/Отменить",
		"en-US": "/Cancel",
		"fa-IR": "/کنسل",
	},
	COMMAND_TEXT_SETTINGS_PRIMARY_CURRENCY: {
		"ru-RU": "Основная валюта",
		"en-US": "Primary currency",
		"fa-IR": "واحد پول اولیه",
	},
	COMMAND_TEXT_NEW_COUNTERPARTY: {
		"ru-RU": "Добавить",
		"en-US": "Add new",
		"fa-IR": "اضافه کردن مورد جدید",
	},
	MESSAGE_TEXT_LOGIN_CODE: {
		"ru-RU": "Ваш код для входа в приложение: <b>%v</b>",
		"en-US": "Your code for signing in to app: <b>%v</b>",
		"fa-IR": "کد شما برای ورود به برنامه: <b>%v</b>",
	},
	MESSAGE_TEXT_ASK_NEW_COUNTERPARTY_NAME: {
		"ru-RU": `<b>Имя для нового контакта</b>
Напишите сами или выберите из своей адресной книги (<i>через иконку "скрепка"</i>).

<i>Отправьте '.' для отмены</i>`,
		"en-US": `Please enter a name for the new contact:
You can type manually or choose from your address book (<i>throw "clip" icon</i>).

<i>Send '.' to cancel</i>`,
		"fa-IR": `لطفا برای مخاطب جدید یک نام وارد کنید:
میتوانید به صورت دستی تایپ نموده یا از لیست مخاطبین خود انتخاب نمایید (<i>throw "clip" icon</i>).

<i>Send '.' برای کنسل کردن</i>`,
	},
	MESSAGE_TEXT_TRANSFER_IS_CREATING: {
		"ru-RU": "Создаю запись...",
		"en-US": "Creating transfer...",
		"fa-IR": "ایجاد انتقال ...",
	},
	COMMAND_TEXT_PLEASE_WAIT: {
		"ru-RU": "Пожалуйста подождите",
		"en-US": "Please wait",
		"fa-IR": "لطفا صبر کنید",
	},
	MESSAGE_TEXT_PLEASE_WAIT: {
		"ru-RU": "Пожалуйста подождите...",
		"en-US": "Please wait...",
		"fa-IR": "لطفا صبر کنید ...",
	},
	MESSAGE_TEXT_SELF_ACKNOWLEDGEMENT: {
		"ru-RU": "Подтверждение ожидается от %v",
		"en-US": "Acknowledgement is expected from %v",
		"fa-IR": "انتظار تصدیق می رود از %v",
	},
	MESSAGE_TEXT_TRANSFER_ACCEPTED_BY_YOU: {
		"ru-RU": "Вы подтвердили эту транзакцию.",
		"en-US": "You've accepted this transaction.",
		"fa-IR": ".شما این تراکنش را قبول کردید ",
	},
	MESSAGE_TEXT_TRANSFER_DECLINED_BY_YOU: {
		"ru-RU": "Вы НЕ подтвердили эту транзакцию.",
		"en-US": "You've declined this transaction.",
		"fa-IR": ".شما این تراکنش را رد کردید",
	},
	MESSAGE_TEXT_TRANSFER_ACCEPTED_BY_COUNTERPARTY: {
		"ru-RU": "%v подтвердил(a) вашу транзакцию:",
		"en-US": "%v accepted your transaction:",
		"fa-IR": ": تراکنش شمارا تایید کرد %v ",
	},
	MESSAGE_TEXT_TRANSFER_DECLINED_BY_COUNTERPARTY: {
		"ru-RU": "%v НЕ подтвердил(a) вашу транзакцию.",
		"en-US": "%v declined your transaction.",
		"fa-IR": "تراکنش شما را رد کرد  %v declined your transaction.",
	},
	COMMAND_TEXT_SUBSCRIBE_TO_APP: {
		"ru-RU": "Хочу приложение!",
		"en-US": "I want the app!",
		"fa-IR": "!من برنامه را می خواهم",
	},
	COMMAND_TEXT_I_AM_FINE_WITH_BOT: {
		"ru-RU": "Меня вполне устраивает бот!",
		"en-US": "I'm fine with just the bot!",
		"fa-IR": "! ربات به تنهایی برای من کافی است",
	},
	MESSAGE_TEXT_SUBSCRIBED_TO_APP: {
		"ru-RU": "Мы сообщим вам когда приложение будет доступно для загруки.",
		"en-US": "We'll let you know once the app is available for download.",
		"fa-IR": ".وقتی برنامه برای دانلود دردسترس بود به شما اطلاع می دهیم",
	},
	MESSAGE_TEXT_NOT_INTERESTED_IN_APP: {
		"ru-RU": "Чтож, мы рады что вас устраивает наш бот и нет необходимости загружать приложение.",
		"en-US": "Well, we are happy our bot is good enough and there is no need to download an app.",
		"fa-IR": ".خب، ما خوشحال هستیم که ربات برای شما کافی است و نیازی به دانلود برنامه نیست",
	},
	MESSAGE_TEXT_YOUR_AD_COULD_BE_HERE: {
		"ru-RU": "Здесь можно <a href>разместить рекламу</a>",
		"en-US": "You can <a href>advertise here</a>",
		"fa-IR": "شما میتوانید <a href>در اینجا تبلیغ کنید</a>",
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
		"fa-IR": "ببخشید، اما شما تنها میتوانید از اعداد بعنوان مقادیر / اندازه ها استفاده کنید (<i>با دو رقم اعشار</i>).",
	},
	MESSAGE_TEXT_ASK_LENDING_TYPE: {
		"ru-RU": "<b>Что вы дали в долг?</b>",
		"en-US": "<b>What did you lend to someone?</b>",
		"fa-IR": "<b> چه چیزی به کسی قرض داده اید؟</b>",
	},
	MESSAGE_TEXT_CHOOSE_CURRENCY: {
		"ru-RU": `Выберите из меню внизу экрана или <a>выберите валюту из списка</a>.

Если ни один из стандартных вариантов не подходит просто напишите текстом. Например: "<i>яблоко</i>".`,

		"en-US": `Please choose from the options below or <a>select a currency from the list</a>.

If standard options are not enough simply send a text. For example: "<i>apple</i>".`,
		"fa-IR": `لطفا از بین گزینه های زیر انتخاب کنید یا <a>یک واحد پولی از لیست انتخاب کنید</a>.

اگر گزینه های استاندارد کافی نبودند به سادگی یک متن بفرستید ، برای مثال: . "<i>سیب</i>".`,
	},
	MESSAGE_TEXT_ASK_LENDING_AMOUNT: {
		"ru-RU": "Сколько <b>%v</b> вы дали в долг?\n(<i>отправьте '.' чтобы отменить</i>)",
		"en-US": "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)",
		"fa-IR": "چه مقدار <b>%v</b> قرض داده اید؟\n(<i>ارسال '.' برای کنسل کردن</i>)",
	},
	MESSAGE_TEXT_ASK_LENDING_COUNTERPARTY: {
		"ru-RU": "Кому вы дали в долг <b>%v</b>?\n(<i>отправьте '.' чтобы отменить</i>)",
		"en-US": "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		"fa-IR": "چه کسی از شما قرض گرفته است <b>%v</b>?\n(<i>ارسال '.' برای کنسل کردن</i>)",
	},
	MESSAGE_TEXT_ASK_BORROWING_TYPE: {
		"ru-RU": "Что вы взяли в долг?",
		"en-US": "What did you lend?",
		"fa-IR": "چه چیزی قرض داده اید؟",
	},
	MESSAGE_TEXT_ASK_BORROWING_AMOUNT: {
		"ru-RU": "Сколько <b>%v</b> вы взяли в долг?\n(<i>отправьте '.' чтобы отменить</i>)",
		"en-US": "How much <b>%v</b> did you borrow?\n(<i>send '.' to cancel</i>)",
		"fa-IR": "چه مقدار <b>%v</b> قرض گرفته اید؟\n(<i>ارسال '.' برای کنسل کردن</i>)",
	},
	MESSAGE_TEXT_ASK_BORROWING_COUNTERPARTY: {
		"ru-RU": "У кого вы взяли в долг <b>%v</b>?\n(<i>отправьте '.' чтобы отменить</i>)",
		"en-US": "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		"fa-IR": "چه کسی به شما قرض داده است <b>%v</b>?\n(<i>ارسال '.' برای کنسل کردن</i>)",
	},
	MESSAGE_TEXT_YOU_CAN_SEND_RECEIPT: {
		"ru-RU": "Отправить <a receipt>квитанцию</a> для <a counterparty>%v</a>?",
		"en-US": "Should we send a <a receipt>receipt</a> to <a counterparty>%v</a>?",
		"fa-IR": "آیا لازم است ماارسال کنیم یک <a receipt>رسید</a> به <a counterparty>%v</a>?",
	},
	MESSAGE_TEXT_YOU_CAN_SEND_RECEIPT_TO_YOURSELF_BY_SMS: {
		"ru-RU": "К сожалению отправка квитанцию себе по СМС в данный момент отключена. Но вы можете отправить её для %v.",
		"en-US": "Sorry, sending a receipt to yourself by SMS is not available at the moment. You can send it to %v though.",
		"fa-IR": "متاسفم، درحال حاضرارسال یک رسید به خودتان بوسیله پیام کوتاه امکان پذیر نیست. شما میتوانید آنرا ارسال کنید به  %v از طریق.",
	},
	MESSAGE_TEXT_RECEIPT_IS_SENDING_BY_TELEGRAM: {
		"ru-RU": "Отправляем для %v извещение через Telegram...",
		"en-US": "We are sending receipt to %v by Telegram...",
		"fa-IR": "مادرحال ارسال رسید به %v از طریق تلگرام هستیم...",
	},
	MESSAGE_TEXT_RECEIPT_NEW_DEBT_FROM_USER: {
		"ru-RU": "{{.Counterparty}} взял(а) в долг {{.Amount}}.",
		"en-US": "{{.Counterparty}} borrowed from you {{.Amount}}.",
		"fa-IR": "{{.Counterparty}} از شما قرض گرفته است {{.Amount}}.",
	},
	MESSAGE_TEXT_RECEIPT_NEW_DEBT_TO_USER: {
		"ru-RU": "{{.Counterparty}} дал(а) вам в долг {{.Amount}}.",
		"en-US": "{{.Counterparty}} lended to you {{.Amount}}.",
		"fa-IR": "{{.Counterparty}} به شما قرض داده است {{.Amount}}.",
	},
	MESSAGE_TEXT_RECEIPT_RETURN_FROM_USER: {
		"ru-RU": "Вы вернули долг - {{.Counterparty}} получил(а) {{.Amount}}.",
		"en-US": "You returned {{.Amount}} to {{.Counterparty}}.",
		"fa-IR": "شما بازگردانده اید {{.Amount}} به {{.Counterparty}}.",
	},
	MESSAGE_TEXT_RECEIPT_RETURN_TO_USER: {
		"ru-RU": "{{.Counterparty}} вернул вам {{.Amount}}.",
		"en-US": "{{.Counterparty}} returned to you {{.Amount}}.",
		"fa-IR": "{{.Counterparty}} به شما بازپرداخت کرده است {{.Amount}}.",
	},
	MESSAGE_TEXT_DUE_ON: {
		"ru-RU": "<b>Вернуть до</b>: %v",
		"en-US": "<b>Return till</b>: %v",
		"fa-IR": "<b>بازگردانده شود تا</b>: %v",
	},
	MESSAGE_TEXT_NOTE: {
		"ru-RU": "Заметка",
		"en-US": "Note",
		"fa-IR": "نکته",
	},
	MESSAGE_TEXT_COMMENT: {
		"ru-RU": "Комментарий",
		"en-US": "Comment",
		"fa-IR": "نظر",
	},
	MESSAGE_TEXT_LOGIN_TO_WEB_APP: {
		"ru-RU": `Перейдите по <a>ссылке</a> чтобы запустить web-приложение.`,
		"en-US": `Click to <a>sign in</a> to web-app.`,
		"fa-IR": `کلیک کنید تا <a>وارد شوید</a> برنامه وب.`,
	},
	MESSAGE_TEXT_DO_YOU_LIKE_OUR_BOT: {
		"ru-RU": "Вам нравится наш бот?",
		"en-US": "Do you like our bot?",
		"fa-IR": "آیا ربات ما را می پسندید؟",
	},
	MESSAGE_TEXT_ASK_FOR_FEEDBAK: {
		"ru-RU": "Будем признетельны если вы оцените работу нашего приложения. Это займёт всего несколько секунд.",
		"en-US": "We would appreciate if tell us how we doing. It takes just few seconds.",
		"fa-IR": "سپاسگزار خواهیم بود اگر به ما بگویید کارمان چطور بوده است. این تنها چند ثانیه زمان میبرد.",
	},
	COMMAND_TEXT_GIVE_FEEDBACK: {
		"ru-RU": "Оценить приложение",
		"en-US": "Rate this bot",
		"fa-IR": "به این ربات امتیاز بدهید",
	},
	COMMAND_TEXT_OPEN_STOREBOT_FOR_FEEDBACK: {
		"ru-RU": "Оценить на  @Storebot",
		"en-US": "Leave rating at @Storebot",
		"fa-IR": "امتیاز خود را اینجا وارد کنید @Storebot",
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
		/*------------------------------------------------------------*/
		"fa-IR": `
		:شما به ما کمک بسیاری می کنید اگر

		  * ثبت بازخورد مثبت در <a storebot>دایرکتوری روبات ها</a>.

		  * به دوستان خود در مورد برنامه اطلاع رسانی کنید.
		    برای مثال در <a share-fb>فیسبوک</a> یا <a share-twitter>توئیتر</a>.

		  * از توسعه بیشتر برنامه پشتیبانی کنید - <a href="https://goo.gl/Qhh0yL">€2 از طریق پی پال</a> (<i>$2.2 حدود </i>)
		`,
	},
	MESSAGE_TEXT_COUNTERPARTY_HAS_EMPTY_BALANCE: {
		"ru-RU": `Нулевой баланс для %v`,
		"en-US": `Balance is empty for %v`,
		"fa-IR": `تراز خالی است برای %v`,
	},
	MESSAGE_TEXT_ASK_TO_TRANSLATE: {
		"ru-RU": `Хотите чтобы наш бот разговаривал на другом языке? Вы можете <a>помочь с переводом</a>.`,
		"en-US": `Do you want our bot to talk in other language? You can <a>help with translation</a>.`,
		"fa-IR": `آیا می خواهید ربات ما به زبان دیگری صحبت کند؟ شما می توانید <a>با ترجمه به ما کمک کنید</a>.`,
	},
	MESSAGE_TEXT_ON_FEEDBACK_NEUTRAL: {
		"ru-RU": `Чтож, мы очень старались. Ваша оценка будет передана разработчикам.

Может быть вы <a submit-bug>сообщите нам что не работает</a> или подскажите <a suggest-idea>как можно улучшить</a>?`,
		/*------------------------------------------------------------*/
		"en-US": `Well, we worked hard. You feedback will be passed to developers.

Maybe you can <a submit-bug>report your issue</a> or <a suggest-idea>suggest how we can improve</a>?`,
		/*------------------------------------------------------------*/
		"fa-IR": `خب، ما سخت کارکردیم. بازخورد شما به توسعه دهندگان برنامه منعکس می شود.

شما می توانید <a submit-bug>مشکل خود را گزارش دهید</a> یا <a suggest-idea>پیشنهاد دهید چطور میتوانیم بهبود ایجاد کنیم</a>?`,
	},
	MESSAGE_TEXT_ON_FEEDBACK_NEGATIVE: {
		"ru-RU": `Нам очень стыдно. Может быть вы <a submit-bug>подскажите что не так</a> или <a suggest-idea>предложите усовершенствования</a>?`,
		/*------------------------------------------------------------*/
		"en-US": `We are very sorry. Maybe you can <a submit-bug>let us know what is wrong</a> or <a suggest-idea>suggest how we can improve</a>?`,
		/*------------------------------------------------------------*/
		"fa-IR": `ما بسیار متاسفیم. شما می توانید <a submit-bug>به ما بگویید مشکلتان چیست</a> یا <a suggest-idea>پیشنهاد دهید چطور میتوانیم بهبود ایجاد کنیم</a>?`,
	},
	COMMAND_TEXT_ASK_FOR_FEEDBACK: {
		"ru-RU": "Оцените наше приложение?",
		"en-US": "Please rate our app",
		"fa-IR": "لطفاً به برنامه ما امتیاز دهید",
	},
	COMMAND_TEXT_FEEDBACK_POSITIVE: {
		"ru-RU": "Да, отличное приложение!",
		"en-US": "Yes, it's a great app!",
		"fa-IR": "بله، این برنامه عالی است",
	},
	COMMAND_TEXT_FEEDBACK_NEUTRAL: {
		"ru-RU": "Неплохо, но можно лучше.",
		"en-US": "Not bad but can be better.",
		"fa-IR": "بد نیست ولی می تواند بهتر باشد.",
	},
	COMMAND_TEXT_FEEDBACK_NEGATIVE: {
		"ru-RU": "Не нравится",
		"en-US": "Don't like it",
		"fa-IR": "از این برنامه را نمی پسندم",
	},
	COMMAND_TEXT_FEEDBACK_NOT_READY: {
		"ru-RU": "Пока не понятно",
		"en-US": "Not decided yet",
		"fa-IR": "هنوز تصمیم نگرفته ام.",
	},
	MESSAGE_TEXT_SETTINGS: {
		"ru-RU": "Что будем настраивать?",
		"en-US": "What do you want to adjust?",
		"fa-IR": "می خواهید چه چیزی را تنظیم کنید؟",
	},
	MESSAGE_TEXT_NOT_IMPLEMENTED_YET: {
		"ru-RU": "Извините, данный функционал ещё не реализован.",
		"en-US": "Sorry, this functionality is not implemented yet.",
		"fa-IR": "متاسفم، این عملکرد هنوز پیاده سازی نشده است.",
	},
	MESSAGE_TEXT_ASK_INVITE_CHANNEL: {
		"ru-RU": "Как вы хотите получить код приглашения?",
		"en-US": "How do you want to get an invite?",
		"fa-IR": "می خواهید چگونه دعوت شوید؟",
	},
	MESSAGE_TEXT_PLEASE_ENTER_INVITE_CODE: {
		"ru-RU": "Пожалуйста введите код приглашения:",
		"en-US": "Please enter an invite code:",
		"fa-IR": "لطفاً یک کد دعوت وارد کنید:",
	},
	MESSAGE_TEXT_USER_EMAIL_FOR_INVITE_RECEIVED: {
		"ru-RU": "Мы отправили письмо на %v.\n\nПожалуйста откройте его и кликните на ссылку для подтверждения адреса.",
		"en-US": "We've sent a message to %v.\n\nPlease open the email and click a link to confirm your email address.",
		"fa-IR": "ما یک پیام ارسال کردیم به %v.\n\nلطفاً ایمیل خود را باز کرده و روی لینک کلیک کنید تا آدرس ایمیل شما تایید شود.",
	},
	MESSAGE_TEXT_USER_EMAIL_FOR_INVITE_SENT_TELEGRAM: {
		"ru-RU": "После того как откроется Telegram нажмите кнопку <b>Start</b>.",
		"en-US": "Once Telegram pop ups click the <b>Start</b> button.",
		"fa-IR": "وقتی تلگرام اجرا شد برروی دکمه  <b>شروع</b> کلیک کنید.",
	},
	MESSAGE_TEXT_USER_CONTACT_FOR_INVITE_RECEIVED: {
		"ru-RU": "Спасибо, вы записаны в очередь на получение приглашения.\n\nТекущее время ожидания 2-3 дня.\n\nВы можете получить приглашение сегодня если расскажите о нашем боте на Facebook.",
		"en-US": "Thanks, you've been queued for an invite.\n\nCurrent awaiting time is 2-3 days.\n\nYou can get an invite code today by sharing a link to the bot on Facebook.",
		"fa-IR": "سپاسگزاریم، شما در نوبت دعوت قرار گرفتید\n\nزمان انتظار شما در حال حاضر 2-3 روز می باشد.\n\n شما می توانید با به اشتراک گذاری لینک روبات در فیسبوک امروز یک کد دعوت دریافت کنید. ",
	},
	MESSAGE_TEXT_PLEASE_PROVIDE_YOUR_EMAIL: {
		"ru-RU": "Пожалуйста напишите ваш email адрес:",
		"en-US": "Please provide your email address",
		"fa-IR": "لطفاً آدرس ایمیل خود را وارد کنید.",
	},
	MESSAGE_TEXT_PLEASE_PROVIDE_YOUR_PHONE_NUMBER: {
		"ru-RU": "Пожалуйста напишите номер вашего телефона:",
		"en-US": "Please provide your phone number",
		"fa-IR": "لطفاً شماره تلفن خود را وارد نمایید.",
	},
	MESSAGE_TEXT_WRONG_INVITE_CODE: {
		"ru-RU": "Неправильный код приглашения: %v",
		"en-US": "Wrong invite code: %v",
		"fa-IR": "کد دعوت اشتباه است %v",
	},
	MESSAGE_TEXT_WRONG_EMAIL: {
		"ru-RU": "Неправильный email адрес.",
		"en-US": "Wrong email address.",
		"fa-IR": "آدرس ایمیل اشتباه است.",
	},
	MESSAGE_TEXT_WRONG_PHONE_NUMBER: {
		"ru-RU": "Неправильный номер телефона.",
		"en-US": "Wrong phone number.",
		"fa-IR": "شماره تلفن اشتباه است",
	},
	MESSAGE_TEXT_OK_PLEASE_TRY_AGAIN: {
		"ru-RU": "Хорошо, попробуйте ещё раз.",
		"en-US": "Ok, please try again.",
		"fa-IR": "بسیار خوب، لطفا مجدداً سعی کنید.",
	},

	COMMAND_TEXT_MISTYPE_WILL_TRY_AGAIN: {
		"ru-RU": "Я опечатался, попробую ещё раз.",
		"en-US": "I've mistyped, will try again.",
		"fa-IR": "اشتباه تایپ کردم، دوباره امتحان می کنم.",
	},
	COMMAND_TEXT_TELL_ME_MORE_ABOUT_INVITES: {
		"ru-RU": "Расскажите ка мне об этих кодах",
		"en-US": "Tell me more about the codes",
		"fa-IR": "در مورد کدها بیشتر برای من توضیح دهید.",
	},
	COMMAND_TEXT_INVITE_ME_BY_EMAIL: {
		"ru-RU": "Хочу код приглашения на email",
		"en-US": "Send me invite code by email",
		"fa-IR": "کد دعوت را برای من از طریق ایمیل ارسال کنید.",
	},
	COMMAND_TEXT_INVITE_ME_BY_SMS: {
		"ru-RU": "Хочу код приглашения по SMS",
		"en-US": "Send me invite code by SMS",
		"fa-IR": "کد دعوت را برای من از طریق پیام کوتاه ارسال کنید.",
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_EMAIL: {
		"ru-RU": "Новый код приглашения на email",
		"en-US": "Send me new invite code by email",
		"fa-IR": "یک کد دعوت جدیداز طریق ایمیل برای من  ارسال کنید.",
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_SMS: {
		"ru-RU": "Новый код приглашения по SMS",
		"en-US": "Send me new invite code by SMS",
		"fa-IR": "یک کد دعوت جدید از طریق پیام کوتاه برای من ارسال کنید.",
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_TELEGRAM: {
		"ru-RU": "Получить приграшение в Telegram",
		"en-US": "Send me new invite by Telegram",
		"fa-IR": "یک کد دعوت جدید از طریق تلگرام برای من ارسال کنید.",
	},
	MESSAGE_TEXT_UNKNOWN_LANGUAGE: {
		"ru-RU": "Незнакомый язык. Пожалуйста выберете один из предоставленных:",
		"en-US": "Unknown language. Please choose 1 from the options:",
		"fa-IR": "زبان ناشناخته. لطفاً یکی از گزینه ها را انتخاب کنید:",
	},
	MESSAGE_TEXT_UNKNOWN_COUNTERPARTY: {
		"ru-RU": "Неизвестный контакт. Пожалуйста выберите <b>Добавить</b> если это новый контакт.",
		"en-US": "Unknown counterparty. Please choose <b>Add new</b> if it's a new contact.",
		"fa-IR": "مخاطب تراکنش شناسایی نشد. <b>یک مورد جدید اضافه کنید</b> اگر این ایشان یک مخاطب جدید هستند.",
	},
	MESSAGE_TEXT_UNKNOWN_COUNTERPARTY_ON_RETURN: {
		"ru-RU": "Неизвестный контакт. Пожалуйста выберите из списка.",
		"en-US": "Unknown counterparty. Please choose from the list.",
		"fa-IR": "مخاطب تراکنش شناسایی نشد. لطفاً از فهرست انتخاب کنید.",
	},
	MESSAGE_TEXT_UNKNOWN_DEBT: {
		"ru-RU": "Неизвестный долг. Пожалуйста выберите из списка.",
		"en-US": "Unknown debt. Please choose from the list.",
		"fa-IR": "بدهی ناشناخته است. لطفا از فهرست انتخاب کنید.",
	},

	MESSAGE_TEXT_HI: { // This is the same for all languages.
		"ru-RU": `¡Hola! Hi! Привет! سلام!`,
		"en-US": `¡Hola! Hi! Привет! سلام!`,
		"fa-IR": `¡Hola! Hi! Привет! سلام!`,
	},
	MESSAGE_TEXT_BACK_TO_MAIN_MENU: {
		"ru-RU": `Можно вернуться назад в главное /меню`,
		"en-US": `You can go back to main /menu`,
		"fa-IR": `شما میتوانید به منو /اصلی مراجعه کنید.`,
	},
	MESSAGE_TEXT_YOUR_SELECTED_PREFERRED_LANGUAGE: { // This is the same for all languages.
		"ru-RU": `Выбранный язык программы: %v`,
		"en-US": `Preferred app language: %v`,
		"fa-IR": `زبان برنامه: %v`,
	},
	MESSAGE_TEXT_ONBOARDING_ASK_TO_CHOOSE_LANGUAGE: {
		"ru-RU": `<b>%v</b>, на каком языке вы хотели бы общаться?

(What is your preferred language?)`,
		"en-US": `<b>%v</b>, what is your preferred language?`,
		"fa-IR": `<b>%v</b>, شما چه زبانی را ترجیح می دهید؟`,
	},
	MESSAGE_TEXT_CHOOSE_UI_LANGUAGE: {
		"ru-RU": "На каком языке вы хотели бы общаться со мной?",
		"en-US": "Which language you would like to talk to me?",
		"fa-IR": "دوست دارید به چه زبانی با من صحبت کنید؟",
	},
	MESSAGE_TEXT_LOCALE_CHANGED: {
		"ru-RU": "Вы поменяли язык на %v",
		"en-US": "You've switched language to %v",
		"fa-IR": "شما زبان را تغییر دادید به %v",
	},
	MESSAGE_TEXT_WHATS_NEXT: {
		"ru-RU": "Что будем делать дальше?",
		"en-US": "What's next?",
		"fa-IR": "بعد چه چیزی است؟",
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
		"fa-IR": `
اگر از کسی قرض گرفته اید برای ثبت مصرف / گرفتن.
اگر به کسی قرض داده اید برای ثبت مصرف / دادن.

یا از منوی پایین استفاده نمایید.
`,
	},
	MESSAGE_TEXT_HISTORY_HEADER: {
		"ru-RU": "История",
		"en-US": "History",
		"fa-IR": "سوابق",
	},
	MESSAGE_TEXT_HISTORY_NO_RECORDS: {
		"ru-RU": "У вас пока нет ни одной записи.",
		"en-US": "You don't have any records yet.",
		"fa-IR": "شما هنوز هیچ ثبت سابقه ای ندارید.",
	},
	MESSAGE_TEXT_HISTORY_LIST: {
		"ru-RU": `<b>%v</b> <i>(%d последних)</i>
─────────────
%v`,
		"en-US": `<b>%v</b> <i>(last %d):</i>
─────────────
%v`,
		"fa-IR": `<b>%v</b> <i>(آخرین %d):</i>
─────────────
%v`,
	},
	MESSAGE_TEXT_BALANCE_IS_ZERO: {
		"ru-RU": "Нет записей о текущих долгах.",
		"en-US": "You have no records on current debts.",
		"fa-IR": "شما در خصوص بدهی های اخیر ثبت سابقه ای ندارید.",
	},
	MESSAGE_TEXT_BALANCE_CURRENCY_TOTAL_INTRO: {
		"ru-RU": "Всего",
		"en-US": "Total",
		"fa-IR": "مجموع",
	},
	MESSAGE_TEXT_PRIMARY_CURRENCY_IS_SET_TO: {
		"ru-RU": "OK, теперь я буду использовать '%v' как основную валюту.",
		"en-US": "OK, from now on I will use '%v' as a primary currency.",
		"fa-IR": "بسیار خوب، از الان من از '%v' بعنوان واحد پول اولیه استفاده می کنم",
	},
	MESSAGE_TEXT_BALANCE_SINGLE_CURRENCY_COUNTERPARTY_DEBT_TO_USER: {
		"ru-RU": "<b>%v</b> - долг вам %v",
		"en-US": "<b>%v</b> - owes you %v",
		"fa-IR": "<b>%v</b> - به شما بدهکار است %v",
	},
	MESSAGE_TEXT_BALANCE_CURRENCY_ROW_DEBT_TO_USER: {
		"ru-RU": "Вам должны %v",
		"en-US": "Owes to you %v",
		"fa-IR": "به شما بدهکار است %v",
	},
	MESSAGE_TEXT_ON_RETURN_USER_DOES_NOT_OWE_ANYTHING_TO_COUNTERPARTY_ANYMORE: {
		"ru-RU": "Поздравляем! У вас не осталось долгов перед <b>%v</b>.",
		"en-US": "Congratulations! You don't owe anything more to <b>%v</b>.",
		"fa-IR": "تبریک! شما دیگر چیزی بدهکار نیستید به <b>%v</b>.",
	},
	MESSAGE_TEXT_ON_RETURN_COUNTERPARTY_DOES_NOT_OWE_ANYTHING_TO_USER_ANYMORE: {
		"ru-RU": "У <b>%v</b> больше не осталось долгов перед вами.",
		"en-US": "<b>%v</b> does not owe anything more to you.",
		"fa-IR": "<b>%v</b> دیگر چیزی به شما بدهکار نیست",
	},
	MESSAGE_TEXT_BALANCE_CURRENCY_ROW_DEBT_BY_USER: {
		"ru-RU": "Вы должны %v",
		"en-US": "You owe %v",
		"fa-IR": "شما بدهکار هستید %v",
	},
	MESSAGE_TEXT_BALANCE_SINGLE_CURRENCY_COUNTERPARTY_DEBT_BY_USER: {
		"ru-RU": "<b>%v</b> - вы должны %v",
		"en-US": "<b>%v</b> - you owe %v",
		"fa-IR": "<b>%v</b> - شما بدهکار هستید %v",
	},
	MESSAGE_TEXT_ASK_PRIMARY_CURRENCY: {
		"ru-RU": "Какая валюта для вас основная?",
		"en-US": "What is your primary currency?",
		"fa-IR": "واحد پولی اولیه شما چیست؟",
	},
	MESSAGE_TEXT_FAILED_TO_DELETE_USER: {
		"ru-RU": "Не удалось удалить данные пользователя: %v",
		"en-US": "Failed to delete user: %v",
		"fa-IR": "حذف کاربر ناموفق بود: %v",
	},
	MESSAGE_TEXT_USER_DELETED: {
		"ru-RU": "Данные пользователя удалены",
		"en-US": "User's data has been deleted",
		"fa-IR": "اطلاعات کاربر حذف شد.",
	},
	MESSAGE_TEXT_PLEASE_WAIT_WHILE_WE_GENERATE_INVITE_CODE: {
		"ru-RU": "Пожалуйста подождите пока мы генерируем секретный код доступа...",
		"en-US": "Please wait a moment while we are generating a security access code...",
		"fa-IR": "لطفاً کمی صبر کنید تا ما یک کد دسترسی امنیتی  ایجاد کنیم.",
	},
	MESSAGE_TEXT_RETURN_ASK_TO_CHOOSE_COUNTERPARTY: {
		"ru-RU": "Выберете кому вы вернули долг или кто вернул его вам.",
		"en-US": "Please choose who returned the debt or to who you returned it.",
		"fa-IR": "لطفاً انتخاب کنید چه کسی بدهی اش را به شما پرداخت کرده یا شما بدهیتان را به چه کسی بازپرداخت نموده اید.",
	},
	MESSAGE_TEXT_CHOOSE_DEBT_THAT_HAS_BEEN_RETURNED: {
		"ru-RU": "Выберите долг который был возвращён целиком или частично.",
		"en-US": "Please choose a debt that has been returned fully or partially.",
		"fa-IR": "لطفاً انتخاب کنید تمام یا بخشی از کدام بدهی پرداخت شده است.",
	},
	MESSAGE_TEXT_PLEASE_ACKNOWLEDGE_TRANSFER: {
		"ru-RU": "Пожалуйста подтвердите или отклоните эту транзакцию.",
		"en-US": "Please confirm or decline this transfer.",
		"fa-IR": "لطفاً این تراکنش را تایید را رد نمایید.",
	},
	MESSAGE_TEXT_ALREADY_ACCEPTED_TRANSFER: {
		"ru-RU": "Эта транзакция уже подтверждена.",
		"en-US": "This transfer has been accepted already.",
		"fa-IR": "این تراکنش قبلا قبول شده است.",
	},
	MESSAGE_TEXT_ALREADY_DECLINED_TRANSFER: {
		"ru-RU": "Эта транзакция уже отклонена.",
		"en-US": "This transfer has been declined already.",
		"fa-IR": "این تراکنش قبلاً رد شده است.",
	},
	MESSAGE_TEXT_RECEIPT_LINK: {
		"ru-RU": "Подробнее тут: %v",
		"en-US": "Details here: %v",
		"fa-IR": "جزئیات: %v",
	},
	MESSAGE_TEXT_ASK_PHONE_NUMBER_OF_COUNTERPARTY: {
		"ru-RU": "Пожалуйста напишите номер телефона <b>%v</b>.",
		"en-US": "Please provide phone number for <b>%v</b>",
		"fa-IR": "لطفا شماره تلفن ایشان را وارد کنید <b>%v</b>",
	},
	MESSAGE_TEXT_USE_CONTACT_TO_SEND_PHONE_NUMBER: {
		"ru-RU": "Если номер телефона есть в записной книжке <b>воспользуйтесь кнопкой %v</b> (скрепка) чтобы отправить контакт.",
		"en-US": "If phone number is in your address book you can <b>use %v button</b> to send the contact.",
		"fa-IR": "اگر شماره تلفن در فهرست مخاطبین شما وجود دارد شما می توانید <b> با استفاده از این %v دکمه</b> تماس را ارسال نمایید.",
	},
	MESSAGE_TEXT_ABOUT_PHONE_NUMBER_FORMAT: {
		"ru-RU": "Номер должен быть в международном формате:\n\t* Начинаться со знака '+' и кода страны\n\t* Состоять только из цифр\nПример: <pre>+</pre><b>7</b><code>999012345678</code>",
		"en-US": "The number should be in international standard:\n\t* Starts with '+' following by country code\n\t* Consist of numbers only\nExample: <pre>+</pre><b>1</b><code>999012345678</code>",
		"fa-IR": "شماره باید به صورت استاندارد بین المللی باشد\n\t* با '+' شروع شده و بدنبال آن کد کشور وارد شود\n\t* تنها شامل اعداد باشد\nمثال: <pre>+</pre><b>1</b><code>999012345678</code>",
	},
	MESSAGE_TEXT_THIS_NUMBER_WILL_BE_USED_TO_SEND_RECEIPT: {
		"ru-RU": "На этот номер мы отправим SMS:",
		"en-US": "Will send an SMS to this number:",
		"fa-IR": "یک پیام کوتاه به این شماره ارسال خواهد شد:",
	},
	MESSAGE_TEXT_COUNTERPARTY_OWES_YOU_SINGLE_DEBT: {
		"ru-RU": `<b>%v</b> одалживал(а) у вас <b>%v</b>.`,
		"en-US": `<b>%v</b> owed to you <b>%v</b>.`,
		"fa-IR": `<b>%v</b> به شما بدهکار بوده <b>%v</b>.`,
	},
	MESSAGE_TEXT_YOU_OWE_TO_COUNTERPARTY_SINGLE_DEBT: {
		"ru-RU": `<b>%v</b> одалживал(а) вам <b>%v</b>.`,
		"en-US": "You owe to <b>%v</b> <b>%v</b>.",
		"fa-IR": "شما بدهکار هستید به <b>%v</b> <b>%v</b>.",
	},
	MESSAGE_TEXT_IS_IT_RETURNED_IN_FULL: {
		"ru-RU": `Возвращено полностью?

		<i>Если частично можете сразу написать сумму.</i>`,

		"en-US": `Has this debt been returned in full?

		<i>If partially you can enter amount right away.</i>`,
		
		"fa-IR": `آیا این بدهی بصورت کامل بازپرداخت شده است؟

		<i>اگر بخشی از بدهی پرداخت شده است شما میتوانید مبلغ را وارد کنید.</i>`,
	},
	MESSAGE_TEXT_PLEASE_HELP_MAKE_IT_BETTER: {
		"ru-RU": `Эта программа <b>бесплатна</b>. <a href="https://debtstracker.io/">Помогите</a> сделать её лучше!`,

		"en-US": `This program is <b>free to use</b>. Please <a href="https://debtstracker.io/">help</a> to make it better!`,
		
		"fa-IR": `این برنامه <b>رایگان می باشد</b>.لطفاً <a href="https://debtstracker.io/">به ما کمک کنید</a>تا آنرا بهبود دهیم!`,
	},
	BUTTON_TEXT_YOU_OWE_AMOUNT_TO_SOMEONE: {
		"ru-RU": "%v | вы должны: %v",
		"en-US": "%v | you owe: %v",
		"fa-IR": "%v | شما بدهکارید: %v",
	},
	BUTTON_TEXT_SOMEONE_OWES_TO_YOU_AMOUNT: {
		"ru-RU": "%v | долг вам: %v",
		"en-US": "%v | owes to you: %v",
		"fa-IR": "%v | به شما بدهکار است: %v",
	},
	BUTTON_TEXT_DEBT_RETURNED_FULLY: {
		"ru-RU": "Да, целиком",
		"en-US": "Yes, fully",
		"fa-IR": "بله، به صورت کامل",
	},
	BUTTON_TEXT_DEBT_RETURNED_PARTIALLY: {
		"ru-RU": "Нет, только часть",
		"en-US": "No, just partially",
		"fa-IR": "خیر، تنها قسمتی",
	},
	MESSAGE_TEXT_ATTEMPT_TO_USE_OWN_INVITE: {
		"ru-RU": "Хорошая попытка пригласить самого себя ;)",
		"en-US": "You should not use your own invite ;)",
		"fa-IR": "نباید از دعوت خود استفاده کنید ;)",
	},
	MESSAGE_TEXT_WELCOME_ONBOARDING_INVITE_ACCEPTED: {
		"ru-RU": "Спасибо за то что воспользовались приглашением!",
		"en-US": "Welcome and thanks for accepting the invite!",
		"fa-IR": "خوش آمدید و سپاسگزاریم که دعوت را پذیرفتید!",
	},
	MESSAGE_TEXT_FOR_COUNTERPARTY_ONLY: {
		"ru-RU": "Это действие доступно только для %v",
		"en-US": "This action for %v only.",
		"fa-IR": "این عمل تنها برای %v می باشد.",
	},
	BUTTON_TEXT_SEE_RECEIPT_DETAILS: {
		"ru-RU": "Показать детали",
		"en-US": "Show receipt details",
		"fa-IR": "جزئیات رسید را نشان بده",
	},
	MESSAGE_TEXT_YOU_SELECTED_INVITE_BY_EMAIL: {
		"ru-RU": "Вы решили пригласить друга через email.",
		"en-US": "You've selected to invite friend by email.",
		"fa-IR": "شما انتخاب کردید که یک دوست را بوسیله ایمیل دعوت کنید.",
	},
	MESSAGE_TEXT_YOU_SELECTED_INVITE_BY_SMS: {
		"ru-RU": "Вы решили пригласить друга через SMS.",
		"en-US": "You've selected to invite friend by SMS.",
		"fa-IR": "شما انتخاب کردید که یک دوست را بوسیله پیام کوتاه دعوت کنید",
	},
	MESSAGE_TEXT_ABOUT_INVITES: {
		"ru-RU": `На данный момент доступ к нашему боту ограничен, но вы можете пригласить друга.

Как вы хотите передать код приглашение?`,
		"en-US": `At the moment access to our bot is limited but you can invite your friend.

How do you want to pass the invite code?`,
		"fa-IR": `در حال حاضر دسترسی به ربات محدود می باشد ولی شما می توانید دوست خود را دعوت کنید.

How do آیا میخواهید کد دعوت را ارسال کنید؟`,
	},
	MESSAGE_TEXT_USER_BLOCKED_TRANSFER_NOTIFICATIONS_BY: {
		"ru-RU": "%v заблокировал получение оповешений о транзакиях через: %v.",
		"en-US": "%v blocked reminders about transactions by: %v",
		"fa-IR": "%v یادآور تراکنش مسدود شده است بوسیله ی: %v",
	},
	COMMAND_TEXT_WAIT_A_SECOND: {
		"ru-RU": "Секундочку...",
		"en-US": "Wait a second...",
		"fa-IR": "یک ثانیه صبر کنید ...",
	},
	HTML_USING_TELEGRAM: {
		"ru-RU": "используя Telegram мессенджер",
		"en-US": "using Telegram messenger",
		"fa-IR": "استفاده از پیام رسان تلگرام",
	},
	COMMAND_TEXT_ACCEPT: {
		"ru-RU": "Согласиться",
		"en-US": "Accept",
		"fa-IR": "قبول",
	},
	//BUTTON_TEXT_ACCEPT_TRANSFER_USING_TELEGRAM:{
	//	"ru-RU": "Подтвердить ",
	//	"en-US": "Accept ",
	//	"fa-IR": "قبول ",
	//},
	//BUTTON_TEXT_DECLINE_TRANSFER_USING_TELEGRAM:{
	//	"ru-RU": "Отказаться (используя Telegram)",
	//	"en-US": "Decline (using Telegram messenger)",
	//	"fa-IR": "رد ( استفاده از پیام رسان تلگرام)",
	//},
	COMMAND_TEXT_DECLINE: {
		"ru-RU": "Отклонить",
		"en-US": "Decline",
		"fa-IR": "رد",
	},
	COMMAND_TEXT_ACCEPT_INVITE: {
		"ru-RU": "Принять приглашение",
		"en-US": "Accept invite",
		"fa-IR": "قبول دعوت",
	},
	COMMAND_TEXT_VIEW_RECEIPT_DETAILS: {
		"ru-RU": "Посмотреть квитанцию",
		"en-US": "See receipt details",
		"fa-IR": "دیدن جزئیات رسید",
	},
	COMMAND_TEXT_OTHER_WAYS_TO_SEND_INVITE: {
		"ru-RU": "Другие способы отправить приглашение",
		"en-US": "Other ways to send an invite",
		"fa-IR": "سایر راههای ارسال دعوت",
	},
	COMMAND_TEXT_SEND_MY_PHONE_NUMBER: {
		"ru-RU": "Отправить мой номер",
		"en-US": "Send my phone number",
		"fa-IR": "شماره تلفن مرا ارسال کنید",
	},
	COMMAND_TEXT_SEND_BY_EMAIL: {
		"ru-RU": "Через Email",
		"en-US": "By Email",
		"fa-IR": "بوسیله ی ایمیل",
	},
	COMMAND_TEXT_SEND_BY_SMS: {
		"ru-RU": "Через SMS",
		"en-US": "By SMS",
		"fa-IR": "بوسیله پیام کوتاه",
	},
	COMMAND_TEXT_INVITE_BY_TELEGRAM: {
		"ru-RU": "Пригласить через Telegram",
		"en-US": "Invite By Telegram",
		"fa-IR": "دعوت با تلگرام",
	},
	MESSAGE_TEXT_INVITE_CREATED: {
		"ru-RU": `Мы отправили код приглашения на указынный вами адрес. (#%v)

Когда ваш друг потдвердит приглашение у вас будут общий баланс и транзакции (только между вами). Это поможет вам минимизировать усилия по ведению учёта.`,

		"en-US": `We've sent an invite code to your friend. (#%v)

Once your friend accepts invitation you'll share balance & transfers between you to make sure you both are on the same page with minimum efforts.`,
		
		"fa-IR": `ما برای دوست شما یک  پیام دعوت ارسال کردیم. (#%v)

وقتی دوست شما دعوت را بپذیرد شما تراز و مبادلات بین خود را به اشتراک می گذارید تا با کمترین تلاش از درک مشترک میان خود اطمینان حاصل کنید. `,
	},
	MESSAGE_TEXT_INVITE_BY_EMAIL: {
		"ru-RU": "Введите email вашего друга на который мы отправим код приглашения.",
		"en-US": "Please enter emaill address of your friend where we should send an invite code.",
		"fa-IR": "لطفاً آدرس ایمیل دوست خود را وارد کنید تا ما از آن طریق کد دعوت را ارسال نماییم.",
	},
	MESSAGE_TEXT_INVITE_ASK_EMAIL_FOR_RECEIPT: {
		"ru-RU": "Введите email вашего друга (%v) на который мы отправим квитанцию подтверждения.",
		"en-US": "Please enter emaill address of your friend (%v) where we should send the receipt.",
		"fa-IR": "لطفا آدرس ایمیل دوست خود را وارد کنید (%v) تا ما از آن طریق کد دعوت را ارسال نماییم.",
	},
	MESSAGE_TEXT_INVITE_BY_SMS: {
		"ru-RU": "Введите номер телефона вашего друга на который мы отправим код приглашения.",
		"en-US": "Please share a contact or enter phone number of your friend where we should send an invite code.",
		"fa-IR": "لطفا شماره دوستان را که می خواهید برای او کد دعوت بفرستیم از لیست مخاطبین به اشتراک گذاشته یا آن را وارد کنید.",
	},
	MESSAGE_TEXT_INVITE_BY_TELEGRAM: {
		"ru-RU": "Вставьте в чат контакт вашего друга которому вы хотите отправить приглашение.",
		"en-US": "Please share a contact of your friend you wish to send an invite code:",
		"fa-IR": "لطفا اطلاعات تماس دوستتان را که میخواهید برایشان کد دعوت ارسال شود به اشتراک بگذارید.",
	},
	MESSAGE_TEXT_INVALID_EMAIL: {
		"ru-RU": "Неверный email. Проверьте и попробуйте ещё раз? /menu",
		"en-US": "Invalid email. Check and try it again? /menu",
		"fa-IR": "ایمیل غیر معتبر می باشد. آیا بررسی نموده، دوباره سعی می کنید؟/منو",
	},
	MESSAGE_TEXT_INVALID_YEAR: {
		"ru-RU": "Неверный год. Год должен быть 2 или 4 цифры (<i>например 2016 или 16)</i>).",
		"en-US": "Invalid year. The year part should be 2 or 4 numbers (<i>e.g. 2016 or 16</i>).",
		"fa-IR": "سال غیرمعتبر می باشد. سال باید به صورت 2 یا 4 رقمی وارد شود (<i>برای مثال 16 یا 2016</i>).",
	},
	MESSAGE_TEXT_INVALID_MONTH: {
		"ru-RU": "Неверный месяц. Месяц должен быть задан целым числом от 1 до 12.",
		"en-US": "Invalid month. Month should be an integer from 1 to 12.",
		"fa-IR": "ماه غیر معتبر می باشد. ماه باید به صورت عددی صحیح بین 1 تا 12 باشد.",
	},
	MESSAGE_TEXT_INVALID_DAY: {
		"ru-RU": "Неверный день. День должен быть задан целым числом от 1 до 31.",
		"en-US": "Invalid day. The day should be an integer from 1 to 31.",
		"fa-IR": "روز غیر معتبر می باشد. روز باید عددی صحیح بین 1 تا 31 باشد.",
	},
	MESSAGE_TEXT_INVALID_DATE: {
		"ru-RU": "Неверный формат даты. Например для 20 февраля 2019 года надо ввести: 20.02.2019 или 20.02.19",
		"en-US": "Invalid date format. For exampel for 20 February 2019 please submit: 20.02.2019 or 20.02.19",
		"fa-IR": "فرمت تاریخ غیر معتبر می باشد. برای مثال برای 20 فوریه 2019 لطفا اینگونه وارد کنید: 20.02.2019 یا 20.02.19",
	},
	MESSAGE_TEXT_INVALID_PHONE_NUMBER: {
		"ru-RU": "Неверный номер. Проверьте и попробуйте ещё раз? /menu",
		"en-US": "Invalid phone number. Check and try it again? /menu",
		"fa-IR": "شماره تلفن غیر معتبر می باشد. آیا بررسی نموده، مجدداً سعی می کنید؟/منو",
	},
	MESSAGE_TEXT_PHONE_NUMBER_IS_NOT_SMS_CAPABLE: {
		"ru-RU": "Данный номер не принимает SMS. Попробуйте другой номер? /menu",
		"en-US": "This phone number not able to receive SMS. Try another number? /menu",
		"fa-IR": "این شماره تلفن قادر به دریافت پیام کوتاه نمی باشد. آیا شماره دیگری را امتحان میکنید؟/منو",
	},
	MESSAGE_TEXT_NO_CONTACT_RECEIVED: {
		"ru-RU": "Мы не получили контакта. Тут инструкция как это сделать. /menu",
		"en-US": "We have not received any contacts. INSTRUCTION HOW TO DO IT. /menu",
		"fa-IR": "ما هیچ اطلاعات تماسی دریافت نکردیم. دستورالعمل چگونگی انجام این کار./منو",
	},
	MESSAGE_TEXT_CONTACT_NAME_IS_NUMBER: {
		"ru-RU": "Вы ввели только цифры в качестве имени контакта. Пожалуйста используйте текстовые символы.",
		"en-US": "You've entered just digits for a contact name. Please use some text characters.",
		"fa-IR": "شما تنها اعداد را برای نام مخاطب وارد کرده اید. لطفا کاراکتر های متنی وارد کنید.",
	},
	MESSAGE_TEXT_CURRENCY_NAME_IS_NUMBER: {
		"ru-RU": "Вы ввели только цифры в качестве номинала. Пожалуйста используйте текстовые символы.",
		"en-US": "You've entered just digits for currency. Please use some text characters.",
		"fa-IR": "شما تنها اعداد را برای واحد پولی وارد کرده اید. لطفا کاراکترهای متنی وارد کنید.",
	},
	MESSAGE_TEXT_HISTORY_ROW_TO_USER_WITH_NAME: {
		"ru-RU": "%v - %s ⇒ Вам : %s",
		"en-US": "%v - %s ⇒ to you: %s",
		"fa-IR": "%v - %s ⇒ به شما: %s",
	},
	MESSAGE_TEXT_HISTORY_ROW_FROM_USER_WITH_NAME: {
		"ru-RU": "%v - Вы ⇒ %s : %s",
		"en-US": "%v - You ⇒ %s : %s",
		"fa-IR": "%v - شما ⇒ %s : %s",
	},
	MESSAGE_TEXT_LETS_SEND_SMS: {
		"ru-RU": "Давайте отправим SMS",
		"en-US": "Let's send SMS",
		"fa-IR": "پیام کوتاه ارسال کنید",
	},
	MESSAGE_TEXT_SMS_QUEUING_FOR_SENDING: {
		"ru-RU": "SMS ставится в очередь на отправку на номер %v...",
		"en-US": "Queuing SMS for sending to number %v...",
		"fa-IR": "پیام کوتاه برای ارسال به شماره مقابل در حال قرارگیری در نوبت ارسال می باشد %v...",
	},
	MESSAGE_TEXT_SMS_QUEUED_FOR_SENDING: {
		"ru-RU": "SMS поставлена в очередь на отправку на номер %v",
		"en-US": "SMS is queued for sending to number %v",
		"fa-IR": "پیام کوتاه برای شماره مقابل در نوبت ارسال قرار گرفت %v",
	},
	MESSAGE_TEXT_BALANCE_HEADER: {
		"ru-RU": "Баланс",
		"en-US": "Balance",
		"fa-IR": "تراز",
	},
	MESSAGE_TEXT_RECEIPT_AVAILABLE_CHANNELS: {
		"ru-RU": "Извините, в данный момент доступны только эти каналы для отправки квитанции:",
		"en-US": "Sorry, at the moment just this channels are available for sending a receipt:",
		"fa-IR": "متاسفانه، در حال حاضر تنها این کانالها برای ارسال رسید در دسترس می باشند.",
	},
	MESSAGE_TEXT_RECEIPT_SENT_THROW_TELEGRAM: {
		"ru-RU": "Квитанция отправлена через телеграм.",
		"en-US": "Receipt sent through Telegram.",
		"fa-IR": "رسید از طریق تلگرام ارسال شد.",
	},
	MESSAGE_TEXT_RECEIPT_NOT_SENT_AS_COUNTERPARTY_HAS_DISABLED_TG_BOT: {
		"ru-RU": "Квитанция НЕ отправлена через телеграм так как %v удалил чат с ботом.",
		"en-US": "Receipt NOT sent through Telegram as %v has deleted chat with the bot.",
		"fa-IR": "از آنجایی که %v چت انجام شده با روبات را حذف کرده است رسید از طریق تلگرام ارسال نشد.",
	},
	MESSAGE_TEXT_RECEIPT_SENT_THROW_EMAIL: {
		"ru-RU": "Квитанция отправлена через email. (id: %v)",
		"en-US": "Receipt sent through email. (id: %v)",
		"fa-IR": "رسید از طریق ایمیل ارسال شد. (id: %v)",
	},
	MESSAGE_TEXT_RECEIPT_SENT_THROW_SMS: {
		"ru-RU": "Квитанция отправлена через SMS.",
		"en-US": "Receipt sent through SMS.",
		"fa-IR": "رسید از طریق پیام کوتاه ارسال شد.",
	},
	MESSAGE_TEXT_SWITCH_TO_PM_TO_VIEW_RECEIPT: {
		"ru-RU": "Переключитьсь на чат с ботом чтобы посмотреть квитанцию",
		"en-US": "Switch to private mode to see receipt details.",
		"fa-IR": "انتقال به حالت خصوصی جهت رویت جزئیات رسید.",
	},
	MESSAGE_TEXT_RECEIPT_VIEWED_BY_COUNTERPARTY: {
		"ru-RU": "Квитанция просмотрена",
		"en-US": "Receipt viewed",
		"fa-IR": "رسید رویت شد.",
	},
	MESSAGE_TEXT_VIEW_MY_NUMBER_IN_INTERNATIONAL_FORMAT: {
		"ru-RU": "Вы можете посмотреть свой номер телефона в ожидаемоем нами формате.",
		"en-US": "You can view your own phone number in the format we are expecting.",
		"fa-IR": "شما می توانید شماره تلفن خود را با فرمتی که ما انتظار داریم ببینید.",
	},
	COMMAND_TEXT_VIEW_MY_NUMBER_IN_INTERNATIONAL_FORMAT: {
		"ru-RU": "Посмотреть мой номер в ожидаемоем формате",
		"en-US": "View my number in the expectd format",
		"fa-IR": "رویت شماره خودم با فرمت مورد انتظار",
	},
	INLINE_BUTTON_SHOW_FULL_HISTORY: {
		"ru-RU": "Показать всю историю",
		"en-US": "Show full history",
		"fa-IR": "نمایش کامل سوابق",
	},
	MESSAGE_TEXT_INCORRECT_VALUE_NOT_A_NUMBER: {
		"ru-RU": "Это не цифра",
		"en-US": "it is not a number",
		"fa-IR": "این یک شماره نیست",
	},
	MESSAGE_TEXT_INCORRECT_VALUE_IS_NEGATIVE: {
		"ru-RU": "Цифра должна быть положительной (<i>больше нуля</i>)",
		"en-US": "The number should be positive (<i>greater than 0</i>)",
		"fa-IR": "شماره باید مثبت باشد (<i>بزرگتر از 0</i>)",
	},
	MESSAGE_TEXT_ASK_HOW_MUCH_HAS_BEEN_RETURNED: {
		"ru-RU": "Сколько было возвращено?",
		"en-US": "How much have been returned?",
		"fa-IR": "چه مقدار بازپرداخت شده است؟",
	},
	MESSAGE_TEXT_HELP: {
		"ru-RU": "Вы можете сообщить о проблеме или сделать предложения по улучшению программы на нашем сайте.",
		"en-US": "Please report any issue or submit a feature request at our website.",
		"fa-IR": "لطفاً در وب سایت ما هرگونه مسئله ای را اعلام فرموده یا درخواست خود را مطرح نمایید.",
	},
	COMMAND_TEXT_OPEN_USER_REPORT: {
		"ru-RU": "Cтраница поддержки ",
		"en-US": "Support page",
		"fa-IR": "صفحه پشتیبانی",
	},
	COMMAND_TEXT_REPORT_A_BUG: {
		"ru-RU": "Сообщить об ошибке",
		"en-US": "Report a bug",
		"fa-IR": "گزارش یک باگ",
	},
	COMMAND_TEXT_SUBMIT_AN_IDEA: {
		"ru-RU": "Предложить идею",
		"en-US": "Add an idea",
		"fa-IR": "یک ایده اضافه کنید.",
	},
	MESSAGE_TEXT_WELCOME: {
		"ru-RU": `Привет, я Коллектиус - Ваш персональный счетовод и коллектор.

Могу записать кто кому чего должен и, и при необходимости, напомнить когда должок пора возвращать.

Чего изволит новый хозяин?`,
		"en-US": `Hi, I'm Collectius - your personal accountant & collector.

I can record who is owing to whom and remind when the return is due.

What would you like to do?`,
		"fa-IR": `سلام، من کالکتیوس هستم - حسابدار شخصی و مامور وصول شما

من میتوانم اینکه چه کسی به چه کسی بدهکار است را ثبت کرده و زمان بازپرداخت را یادآوری کنم.

دوست دارید چکار کنید؟`,
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE: {
		"ru-RU": "Хочу получить приглашение",
		"en-US": "I want to get an invite",
		"fa-IR": "می خواهم یک دعوت دریافت کنم.",
	},
	COMMAND_TEXT_I_HAVE_INVITE: {
		"ru-RU": "У меня есть код приглашения",
		"en-US": "I have the invitation code",
		"fa-IR": "من کد دعوت را دارم.",
	},
	COMMAND_TEXT_I_HAVE_NOT_GOT_EMAIL: {
		"ru-RU": "Я не получил письма на email",
		"en-US": "I have not got any emails",
		"fa-IR": "من ایمیلی دریافت نکردم",
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
		"fa-IR": `<b>%v</b>,

درحال حاضر ربات ما تنها با دریافت دعوت از دوستان در دسترس می باشد.

اگر شما کدی در اختیار ندارید می توانید اطلاعات تماس خود را برای من وارد نموده و من به محض اینکه نوبت شما فرارسید یک دعوتنامه برایتان ارسال می کنم. 

ما روزانه 10 دعوتنامه برای نفرات اول لیست انتظار و همچنین یک دعوتنامه تصادفی ارسال میکنیم.`,
	},
	EMAIL_INVITE_SUBJ: {
		"ru-RU": "Приглашение от {{.FromName}} - код: {{.InviteCode}}",
		"en-US": "An invite from {{.FromName}} - code: {{.InviteCode}}",
		"fa-IR": "دعوت از طرف {{.FromName}} - کد: {{.InviteCode}}",
	},
	SMS_INVITE_TEXT: {
		"ru-RU": `Привет {{.ToName}}, {{.FromName}} рекомендует приложение для учёта долгов - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Код приглашения: {{.InviteCode}}`,
		"en-US": `Hi{{.ToName}}, {{.FromName}} is inviting you to try debts tracking app - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

You invitation code is: {{.InviteCode}}`,
		"fa-IR": `سلام{{.ToName}}, {{.FromName}} شما را دعوت کرده تا برنامه ردیابی بدهی ها را امتحان کنید. - https://debtstracker.io/invite#id={{.InviteCode}}&ربات تلگرام={{.TgBot}}&{{.Utm}}

کد دعوت شما: {{.InviteCode}}`,
	},
	EMAIL_INVITE_TEXT: {
		"ru-RU": `Привет {{.ToName}},

{{.FromName}} приглашает тебя попробовать приложение для учёта долгов - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Ваш код приглашения: {{.InviteCode}}`,
		"en-US": `Hi{{.ToName}},

{{.FromName}} is inviting you to use debts tracking app - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

You invitation code is: {{.InviteCode}}`,
		"fa-IR": `سلام{{.ToName}},

{{.FromName}} شما را دعوت کرده تا از برنامه ردیابی بدهی ها استفاده کنید. - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

کد دعوت شما: {{.InviteCode}}`,
	},
	EMAIL_INVITE_HTML: {
		"ru-RU": `<p>Привет {{.ToName}},</P

<p>{{.FromName}} приглашает тебя <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">попробовать приложение для учёта долгов</a>.</p>

<p>Ваш код приглашения: <b>{{.InviteCode}}</b></p>`,
		"en-US": `<p>Hi{{.ToName}},</p>

<p>{{.FromName}} п is inviting you to <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">try debts tracking app</a>.</p>

<p>You invitation code is: <b>{{.InviteCode}}</b></p>`,
		"fa-IR": `<p>سلام{{.ToName}},</p>

<p>{{.FromName}} п شما را دعوت کرده به <a href="https://debtstracker.io/invite#id={{.InviteCode}}&ربات تلگرام={{.TgBot}}&{{.Utm}}"> امتحان برنامه ردیابی بدهی ها</a>.</p>

<p>You invitation code is: <b>{{.InviteCode}}</b></p>`,
	},
	EMAIL_RECEIPT_SUBJ: {
		"ru-RU": "Запись о долге - {{.FromName}}",
		"en-US": "Debt record - {{.FromName}}",
		"fa-IR": "سوابق بدهی - {{.FromName}}",
	},
	EMAIL_RECEIPT_BODY_TEXT: {
		"ru-RU": "{{.FromName}} создал(а) запись о долге: {{.ReceiptURL}}",
		"en-US": "{{.FromName}} created a debt record: {{.ReceiptURL}}",
		"fa-IR": "{{.FromName}} یک سابقه بدهی ایجاد کرده است: {{.ReceiptURL}}",
	},
	MESSENGER_RECEIPT_TEXT: {
		"ru-RU": "Я создал(а) запись о долге, подробности тут - {{.ReceiptURL}}",
		"en-US": "I've created a debt record regards you, see details at {{.ReceiptURL}}",
		"fa-IR": "من یک سابقه بدهی برای شما ایجاد کرده ام، لطفا جزئیات را ملاحظه فرمایید در {{.ReceiptURL}}",
	},
	EMAIL_RECEIPT_BODY_HTML: {
		"ru-RU": "{{.FromName}} создал(а) запись о долге: {{.ReceiptURL}}",
		"en-US": "{{.FromName}} created a debt record: {{.ReceiptURL}}",
		"fa-IR": "{{.FromName}} یک سابقه بدهی ایجاد کرده است: {{.ReceiptURL}}",
	},
	INLINE_RECEIPT_TITLE: {
		"ru-RU": "Квитанция: %v",
		"en-US": "Receipt: %v",
		"fa-IR": "رسید: %v",
	},
	INLINE_RECEIPT_DESCRIPTION: {
		"ru-RU": "Нажмите здесь чтобы отправить квитанцию",
		"en-US": "Click here to send the receipt",
		"fa-IR": "برای ارسال رسید اینجا کلیک کنید.",
	},
	INLINE_RECEIPT_CHOOSE_LANGUAGE: {
		"ru-RU": "<b>Выберите язык чтобы посмотреть подробности записи о долге</b> которую создал(а) {{.Creator}}.",
		"en-US": "<b>Please choose language to see details of the debt</b> that has been recorded by {{.Creator}}.",
		"fa-IR": "<b لطفا زبان خود را انتخاب کنید تا جزئیات بدهی رویت شود </b> ثبت سوابق بدهی توسط {{.Creator}}.",
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
		//-------------------------------------------------------
		"fa-IR": `<b>{{.Creator}} بدهی ثبت نموده</b> مرتبط با شما.

{{.SiteLink}} — یک برنامه ردیابی بدهی به شما کمک می کند تا:

  - همیشه از سود و زیان خود مطلع باشید.

  - بدهی های خود را به موقع پرداخت کنید
    <i>(یادآوری به  شما و بدهکاران به شما)</i>`,
	},
	INLINE_INVITE_TITLE: {
		"ru-RU": "Приглашение в %v",
		"en-US": "Invitation to %v",
		"fa-IR": "دعوت به %v",
	},
	INLINE_INVITE_DESCRIPTION: {
		"ru-RU": "Нажмите здесь для отправки приглашения",
		"en-US": "Click here to send an invite",
		"fa-IR": "برای ارسال یک دعوتنامه اینجا کلیک کنید.",
	},
	INLINE_INVITE_MESSAGE: {
		"ru-RU": "%v пригласил вас попробовать %v",
		"en-US": "%v invited you to try %v",
		"fa-IR": "%v شمارا دعوت کرده است به امتحان %v",
	},
	SMS_RECEIPT_YOU_GOT: {
		"ru-RU": "Вы получили %v от %v.",
		"en-US": "You've got %v from %v.",
		"fa-IR": "شما دریافت کرده اید %v از %v.",
	},
	SMS_RECEIPT_YOU_GAVE: {
		"ru-RU": "Вы дали %v - взял(а) %v.",
		"en-US": "You've given %v to %v.",
		"fa-IR": "شما پرداخت کرده اید %v به %v.",
	},
	SMS_CLICK_TO_CONFIRM_OR_DECLINE: {
		"ru-RU": "Нажмите %v чтобы подтвердить или отвергнуть.",
		"en-US": "Click %v to confirm or decline.",
		"fa-IR": "کلیک کنید %v تا رد خود را تایید نمایید",
	},
	HTML_DATE: {
		"ru-RU": "Дата",
		"en-US": "Date",
		"fa-IR": "تاریخ",
	},
	HTML_RECEIPT: {
		"ru-RU": "Квитанция",
		"en-US": "Receipt",
		"fa-IR": "رسید",
	},
	HTML_AMOUNT: {
		"ru-RU": "Сумма",
		"en-US": "Amount",
		"fa-IR": "مقدار",
	},
	HTML_FROM: {
		"ru-RU": "Дал",
		"en-US": "From",
		"fa-IR": "از",
	},
	HTML_TO: {
		"ru-RU": "Получил",
		"en-US": "To",
		"fa-IR": "به",
	},
	TELEGRAM_RECEIPT: {
		"ru-RU": "{{.FromName}} создал запись о долге ({{.TransferCurrency}})",
		"en-US": "{{.FromName}} created a debtrecord ({{.TransferCurrency}})",
		"fa-IR": "{{.FromName}} ایجاد یک سابقه بدهی ({{.TransferCurrency}})",
	},
	MESSAGE_TEXT_PLEASE_CHOOSE_FROM_OPTIONS_PROVIDED: {
		"ru-RU": "Пожалуйста выберете из предоставленных опций.",
		"en-US": "Please choose from provided options.",
		"fa-IR": "لطفاً از گزینه های ارائه شده انتخاب نمایید.",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_NOTE_OR_COMMENT: {
		"ru-RU": "<b>Хотите добавить заметку или комментарий?</b>\n%v Заметки хранятся для вашего личго пользования.\n%v Комментарий виден всем кому разрешён просмотр этой транзакции.",
		"en-US": "<b>Do you want to add a note or comment?</b>\n%v Memos are private records for yoru own reference.\n%v Comments are available to everyone who has permission to view this transaction.",
		"fa-IR": "<b>آیا میخواهید یادداشت یا نظری اضافه کنید؟</b>\n%v یادداشت ها نوشته های خصوصی برای مراجعه خود شما هستند.\n%v نظرات در دسترس تمام کسانی که مجاز رویت این تراکنش هستند میباشد.",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_NOTE: {
		"ru-RU": "Напишите заметку:",
		"en-US": "Please write a note:",
		"fa-IR": "لطفاً یک یادداشت بنویسید:",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_COMMENT_ONLY: {
		"ru-RU": `Если хотите добавить комментарий просто отправьте текст.`,
		"en-US": `If you want to add a comment just send a text now.`,
		"fa-IR": `اگر شما می خواهید یک نظر اضافه کنید فقط الان یک متن ارسال کنید`,
	},
	MESSAGE_TEXT_VISIBLE_TO_YOU_AND_COUNTERPARTY: {
		"ru-RU": "виден вам и %v",
		"en-US": "visible to you & %v",
		"fa-IR": "قابل مشاهده برای شما & %v",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_COMMENT: {
		"ru-RU": "Напишите комментарий:",
		"en-US": "Please write the comment:",
		"fa-IR": "لطفاً نظرتان را ثبت کنید:",
	},
	MESSAGE_TEXT_TRANSFER_NOTE_ADDED_ASK_FOR_COMMENT: {
		"ru-RU": "Заметка добавлена. Хотите написать комментарий?",
		"en-US": "Memo have been added. Do you want to write a comment?",
		"fa-IR": "یادداشت اضافه شد. آیا میخواهید یک نظر ثبت کنید؟",
	},
	MESSAGE_TEXT_TRANSFER_COMMENT_ADDED_ASK_FOR_NOTE: {
		"ru-RU": "Комментарий добавлен. Хотите написать заметку?",
		"en-US": "Comment have been added. Do you want to write a note?",
		"fa-IR": "نظر شما ثبت شد. آیا می خواهید یک یادداشت بنویسید؟",
	},
	COMMAND_TEXT_NO_COMMENT_OR_NOTE_FOR_TRANSFER: {
		"ru-RU": "Без заметок и комментариев",
		"en-US": "Without notes or comments",
		"fa-IR": "بدون یادداشت یا نظر",
	},
	COMMAND_TEXT_NO_COMMENT_FOR_TRANSFER: {
		"ru-RU": "Без комментариев",
		"en-US": "No comments",
		"fa-IR": "بدون نظر",
	},
	COMMAND_TEXT_NO_NOTE_FOR_TRANSFER: {
		"ru-RU": "Без заметок",
		"en-US": "Without notes",
		"fa-IR": "بدون یادداشت",
	},
	COMMAND_TEXT_ADD_NOTE_TO_TRANSFER: {
		"ru-RU": "Добавить заметку",
		"en-US": "Add a note (private)",
		"fa-IR": "یک یادداشت اضافه کنید (خصوصی)",
	},
	COMMAND_TEXT_ADD_COMMENT_TO_TRANSFER: {
		"ru-RU": "Добавить комментарий",
		"en-US": "Add a comment (public)",
		"fa-IR": "یک نظر اضافه کنید (عمومی)",
	},
	DUE_IN_NOW: {
		"ru-RU": "прямо сейчас",
		"en-US": "now",
		"fa-IR": "حالا",
	},
	DUE_IN_A_MINUTE: {
		"ru-RU": "через минуту",
		"en-US": "in a minute",
		"fa-IR": "ظرف یک دقیقه",
	},
	DUE_IN_X_MINUTES: {
		"ru-RU": "через %v минут(ы)",
		"en-US": "in %v minutes",
		"fa-IR": "در %v دقیقه",
	},
	DUE_IN_AN_HOUR: {
		"ru-RU": "через час",
		"en-US": "in an hour",
		"fa-IR": "ظرف یک ساعت",
	},
	DUE_IN_X_HOURS: {
		"ru-RU": "через %v часа/часов",
		"en-US": "in %v hours",
		"fa-IR": "در %v ساعت",
	},
	DUE_IN_X_DAYS: {
		"ru-RU": "через %v дня/дней",
		"en-US": "in %v days",
		"fa-IR": "در %v روز",
	},
	//-------------------------------------------------------------------------------------------------------------------
	WS_ALEX_T: {
		"ru-RU": "Александр Трахимёнок",
		"en-US": "Alexander Trakhimenok",
		"fa-IR": "الکساندر تراخیمِنوک",
	},

	WS_INDEX_TITLE: {
		"ru-RU": "DebtsTracker.io - программа для учёта личных долгов и активов",
		"en-US": "DebtsTracker.io - an IOU app to track your personal debts & assets",
		"es-ES": "DebtsTracker.io - una aplicación para el seguimiento de sus deudas personales",
		"fa-IR": "DebtsTracker.io - برنامه ای برای ردیابی بدهی ها و دارایی های شما",
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
		"fa-IR": "ربات چت برای پیام رسان تلگرام",
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
		"fa-IR": "بازکردن در تلگرام &#x1F680;",
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
		"fa-IR": "درحال حاضر برنامه ما فقط در دسترس است در <a href='https://telegram.org/'>تلگرام</a>",
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
		"fa-IR": "از سود و زیان خود مطلع باشید و هرگز پرداخت بدهی ای را از قلم نیندازید",
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
		"fa-IR": "DebtsTracker.io یک برنامه موبایل و سرویس یادآور می باشد که به شما کمک می کند تا بدهی ها و اعتبارات خود را ردیابی نمایید. همچنین ایمیل و پیام کوتاه یادآوری به بدهکاران ارسال می کند.",
		"pl-PL": "DebtsTracker.io to aplikacje mobilne i przypomnienia usługa, która pozwala na śledzenie swoich długów i kredytów. Wysyła automatycznych powiadomień e-mail i SMS do swoich dłużników.",
		"pt-PT": "DebtsTracker.io é um serviço de aplicativos móveis e lembrete de que ajuda a controlar seus débitos e créditos. Envia e-mail e SMS notificações automáticas aos seus devedores.",
		"de-DE": "DebtsTracker.io ist eine mobile Apps und Erinnerungs-Service, die Ihre Schulden und Kredite zu verfolgen hilft. Sendet automatisierte E-Mail und SMS-Benachrichtigungen an Ihre Schuldner.",
		"fr-FR": "DebtsTracker.io est une des applications mobiles et rappel service qui permet de suivre vos dettes et crédits. Envoie automatisés email & SMS reminders à vos débiteurs.",
		"it-IT": "DebtsTracker.io è un servizio di applicazioni mobili e per ricordare che aiuta a monitorare i debiti e crediti. Invia notifiche e-mail e SMS automatici per i vostri debitori.",
		"ko-KO": "DebtsTracker.io 은 채무 및 크레딧을 추적하는 데 도움이 모바일 앱 및 알림 서비스입니다. 당신의 채무자에 자동화 된 이메일 및 SMS 알림을 보냅니다.",
		"ja-JP": "DebtsTracker.io は、あなたの借金＆クレジットを追跡するのに役立ちますモバイルアプリ＆リマインダーサービスです。あなたの債務者に自動メール＆SMS通知を送信します。",
		"zh-CN": "DebtsTracker.io 是一个移动应用和提醒服务，帮助跟踪你的债务和信用。发送自动电子邮件和短信通知到您的债务人。",
	},

	WS_ADS_TITLE: {
		"en-US": "Ads @ DebtsTracker.IO",
		"ru-RU": "Реклама на DebtsTracker.IO",
		"fa-IR": "تبلیغات @ DebtsTracker.IO",
	},
	WS_ADS_CONTENT: {
		"en-US": `To place ads in our app please send us an email to <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
		"ru-RU": `Чтобы разместить рекламу в нашем приложении пишите нам на <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
		"fa-IR": `برای قراردادن تبلیغات در برنامه ما، درخواست خود را به این آدرس ایمیل فرمایید <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
	},
	FB_MAKE_RECORD_HEADER: {
		"en-US": "Record debt",
		"ru-RU": "Записать долг",
		"fa-IR": "ثبت بدهی",
	},
	FB_MAKE_RECORD_HEADLINE: {
		"en-US": "Scroll left to see other options.",
		"ru-RU": "Пролистайте карточки влево чтобы увидеть другие опции.",
		"fa-IR": "برای دیدن سایر گزینه ها به سمت چپ اسکرول نمایید.",
	},
	FB_HOW_ARE_THINGS_HEADER: {
		"en-US": "How are you doing?",
		"ru-RU": "Как идут дела?",
		"fa-IR": "حال شما چطوره؟",
	},
}
