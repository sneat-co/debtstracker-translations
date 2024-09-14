package trans

import "github.com/bots-go-framework/bots-fw/botsfw"

const adsCommandTitle = "\xE2\xAD\x90\xE2\xAD\x90\xE2\xAD\x90"

// TRANS - translation string
var TRANS = map[string]map[string]string{
	"EXAMPLE": {
		"de-DE": "BEISPIEL",
		"en-UK": "SAMPLE",
		"es-ES": "EJEMPLO",
		"fa-IR": "نمونه",
		"it-IT": "ESEMPIO",
		"ru-RU": "ПРИМЕР",
	},

	"Jan": {
		"de-DE": "Jan",
		"en-UK": "Jan",
		"es-ES": "Enero",
		"fa-IR": "ژانویه",
		"it-IT": "Gen",
		"ru-RU": "Янв.",
	},
	"Feb": {
		"de-DE": "Feb",
		"en-UK": "Feb",
		"es-ES": "Feb",
		"fa-IR": "فوریه",
		"it-IT": "Feb",
		"ru-RU": "Фев.",
	},
	"Mar": {
		"de-DE": "Mär",
		"en-UK": "Mar",
		"es-ES": "Mar",
		"fa-IR": "مارس",
		"it-IT": "Mar",
		"ru-RU": "Мрт.",
	},
	"Apr": {
		"de-DE": "Apr",
		"en-UK": "Apr",
		"es-ES": "Abr",
		"fa-IR": "آوریل",
		"it-IT": "Apr",
		"ru-RU": "Апр.",
	},
	"May": {
		"de-DE": "Mai",
		"en-UK": "May",
		"es-ES": "May",
		"fa-IR": "مه",
		"it-IT": "Mag",
		"ru-RU": "Май ",
	},
	"Jun": {
		"de-DE": "Jun",
		"en-UK": "Jun",
		"es-ES": "Jun",
		"fa-IR": "ژوئن",
		"it-IT": "Giu",
		"ru-RU": "Июнь",
	},
	"Jul": {
		"de-DE": "Jul",
		"en-UK": "Jul",
		"es-ES": "Jul",
		"fa-IR": "ژوئیه",
		"it-IT": "Lug",
		"ru-RU": "Июль",
	},
	"Aug": {
		"de-DE": "Aug",
		"en-UK": "Aug",
		"es-ES": "Ago",
		"fa-IR": "اوت",
		"it-IT": "Ago",
		"ru-RU": "Авг.",
	},
	"Sep": {
		"de-DE": "Sep",
		"en-UK": "Sep",
		"es-ES": "Sep",
		"fa-IR": "سپتامبر",
		"it-IT": "Sett",
		"ru-RU": "Сен.",
	},
	"Oct": {
		"de-DE": "Okt",
		"en-UK": "Oct",
		"es-ES": "Oct",
		"fa-IR": "اکتبر",
		"it-IT": "Ott",
		"ru-RU": "Окт.",
	},
	"Nov": {
		"de-DE": "Nov",
		"en-UK": "Nov",
		"es-ES": "Nov",
		"fa-IR": "نوامبر",
		"it-IT": "Nov",
		"ru-RU": "Нбр.",
	},
	"Dec": {
		"de-DE": "Dez",
		"en-UK": "Dec",
		"es-ES": "Dic",
		"fa-IR": "دسامبر",
		"it-IT": "Dic",
		"ru-RU": "Дек.",
	},
	COMMAND_START: {
		"de-DE": "start",
		"en-UK": "start",
		"es-ES": "inicio",
		"fa-IR": "شروع",
		"it-IT": "inizio",
		"ru-RU": "старт",
	},
	COMMAND_MENU: {
		"de-DE": "menu",
		"en-UK": "menu",
		"es-ES": "menú",
		"fa-IR": "منو",
		"it-IT": "menu", // TODO(IT): Google translated
		"ru-RU": "меню",
	},
	COMMAND_GAVE: {
		"de-DE": "verleihen",
		"en-UK": "gave",
		"es-ES": "prestado_a_ti",
		"fa-IR": "قرض_دادن",
		"it-IT": "debito",
		"ru-RU": "дать",
	},
	COMMAND_GOT: {
		"de-DE": "anleihen",
		"en-UK": "got",
		"es-ES": "prestado_por_ti",
		"fa-IR": "قرض_گرفتن",
		"it-IT": "credito",
		"ru-RU": "взять",
	},
	COMMAND_RETURNED: {
		"de-DE": "beglichen",
		"en-UK": "return",
		"es-ES": "devuelto",
		"fa-IR": "بازگردانده_شده",
		"it-IT": "rientra",
		"ru-RU": "вернуть",
	},
	COMMAND_BALANCE: {
		"de-DE": "ausstehend",
		"en-UK": "balance",
		"es-ES": "balance",
		"fa-IR": "تراز",
		"it-IT": "bilancio",
		"ru-RU": "баланс",
	},
	COMMAND_HISTORY: {
		"de-DE": "verlauf",
		"en-UK": "history",
		"es-ES": "cronología",
		"fa-IR": "سوابق",
		"it-IT": "cronologia",
		"ru-RU": "история",
	},
	COMMAND_SETTINGS: {
		"de-DE": "einstellungen",
		"en-UK": "settings",
		"es-ES": "ajustes",
		"fa-IR": "تنظیمات",
		"it-IT": "impostazioni",
		"ru-RU": "настройки",
	},
	COMMAND_HELP: {
		"de-DE": "hilfe",
		"en-UK": "help",
		"es-ES": "ayuda",
		"fa-IR": "کمک",
		"it-IT": "aiuto",
		"ru-RU": "помощь",
	},
	COMMAND_CANCEL: {
		"de-DE": "abbrechen",
		"en-UK": "cancel",
		"es-ES": "cancelar",
		"fa-IR": "کنسل",
		"it-IT": "annulla",
		"ru-RU": "/отменить",
	},
	COMMAND_CLEAR: {
		"de-DE": "leeren",
		"en-UK": "clear",
		"es-ES": "borrar",
		"fa-IR": "پاک_کردن",
		"it-IT": "chiaro",
		"ru-RU": "очистить",
	},
	adsCommandTitle: {
		"de-DE": adsCommandTitle,
		"en-UK": adsCommandTitle,
		"es-ES": adsCommandTitle,
		"fa-IR": adsCommandTitle,
		"it-IT": adsCommandTitle,
		"ru-RU": adsCommandTitle,
	},
	" and ": {
		"de-DE": " und ",
		"en-UK": " and ",
		"es-ES": " y ",
		"fa-IR": " و ",
		"it-IT": " e ",
		"ru-RU": " и ",
	},
	botsfw.MessageTextOopsSomethingWentWrong: {
		"de-DE": "Ups, etwas ist schiefgelaufen... \xF0\x9F\x98\xB3",
		"en-UK": "Oops, something went wrong... \xF0\x9F\x98\xB3",
		"es-ES": "Ops,  algo ha salido mal... \xF0\x9F\x98\xB3",
		"fa-IR": "اوه، یک جای کار مشکل دارد ...  \xF0\x9F\x98\xB3",
		"it-IT": "Ops, qualcosa e' andato storto... \xF0\x9F\x98\xB3",
		"ru-RU": "Упс, что-то пошло не так... \xF0\x9F\x98\xB3",
	},
	MESSAGE_TEXT_ASK_DUE: {
		"de-DE": "Wann ist die Schuld fällig?",
		"en-UK": "When is the due date?",
		"es-ES": "¿Cuándo es la fecha de devolución?",
		"fa-IR": "سررسید چه زمانی است؟",
		"it-IT": "Data di scadenza?",
		"ru-RU": "Когда дата возврата?",
	},
	MESSAGE_TEXT_ASK_DATE_TO_REMIND: {
		"de-DE": `Sende mir das Datum, an welches du <b>erneut</b> erinnert werden möchtest, in der Form <i>DD.MM.YEAR</i>.
<b>Zum Beispiel</b> für den 20. Januar 2017, schreibe:
    <i>20.01.2017</i>`,

		"en-UK": `To set date for next reminder please send it as a text in format of <i>DD.MM.YEAR</i>.
<b>For example</b> for 20th of January 2017 submit:
    <i>20.01.2017</i>`,

		"es-ES": `Para establecer la fecha recordatoria escribela en el siguiente formato <i>DD.MM.AÑO</i>.
<b>For example</b> para 20 de Enero 2017 envia:
    <i>20.01.2017</i>`,

		"fa-IR": `لطفاً برای تنظیم یادآور بعدی آنرا با متنی با این فرمت ارسال نمایید. <i>روز.ماه.سال</i>.
<b>برای مثال</b> برای 15 خرداد 1396 ثبت کنید:
    <i>15.03.1396</i>`,

		"it-IT": `Per impostare la data per il promemoria successivo invia il messaggio con la data nel seguente formato <i>GG.MM.ANNO</i>.
		<b>Esempio</b> per indicare la data 20 Gennaio 2017 inserisci:

		<i>20.01.2017</i>`,

		"ru-RU": `Чтобы задать дату напопинания напишите её в формате <i>ДД.MM.ГОД</i>.
<b>Например</b> для 20 января 2017 г. отправьте:
    <i>20.01.2017</i>`,
	},
	MESSAGE_TEXT_ASK_DUE_DATE: {
		"de-DE": `Sende mir das Datum, an welches du erinnert werden möchtest, in der Form <i>DD.MM.YEAR</i>.
<b>Zum Beispiel</b> für den 20. Januar 2017, schreibe:
    <i>20.01.2017</i>`,

		"en-UK": `To set due date please send it as a text in format of <i>DD.MM.YEAR</i>.
		<b>For example</b> for 20th of January 2017 submit:
		<i>20.01.2017</i>`,

		"es-ES": `Para establecer la fecha de devolución escribela en el siguiente formato <i>DD.MM.AÑO</i>.
                 <b>For example</b> para 20 de Enero 2017 envia:
                 <i>20.01.2017</i>`,

		"fa-IR": `لطفاً برای تنظیم تاریخ سررسید این فرمت را رعایت فرمایید.<i>روز.ماه.سال</i>.
		<b>برای مثال</b> برای 20 ژانویه 2017 ثبت کنید:
		<i>20.01.2017</i>`,

		"it-IT": `Per impostare la data di scadenza invia il messaggio con la data nel seguente formato <i>GG.MM.ANNO</i>.
		<b>Esempio</b> per indicare la data 20 Gennaio 2017 inserisci:
		<i>20.01.2017</i>`,

		"ru-RU": `Чтобы задать дату возврата напишите её в формате <i>ДД.MM.ГОД</i>.
		<b>Например</b> для 20 января 2017 г.отправьте:
		<i>20.01.2017</i>`,
	},
	MESSAGE_TEXT_WRONG_DATE: {
		"de-DE": "Entschuldigung, aber mit diesem Datum stimmt etwas nicht.",
		"en-UK": "Sorry, there is something wrong with the date you've provided.",
		"es-ES": "Lo siento, algo no es correcto con la fecha que has puesto",
		"fa-IR": "متاسفم، در تاریخی که وارد نمودید مشکلی وجود دارد.",
		"it-IT": "Mi spiace, ma c'e' qualcosa di sbagliato nella data che hai inserito.",
		"ru-RU": "Извините, что-то не так с датой которую вы отправили.",
	},
	COMMAND_TEXT_DISABLE_REMINDER: {
		"de-DE": "Nicht erinnern",
		"en-UK": "No reminder",
		"es-ES": "No recordar",
		"fa-IR": "بدون یادآور",
		"it-IT": "Nessun promemoria",
		"ru-RU": "Не напоминать",
	},
	COMMAND_TEXT_TOMORROW: {
		"de-DE": "Morgen",
		"en-UK": "Tomorrow",
		"es-ES": "Mañana",
		"fa-IR": "فردا",
		"it-IT": "Domani",
		"ru-RU": "Завтра",
	},
	COMMAND_TEXT_DAY_AFTER_TOMORROW: {
		"de-DE": "Übermorgen",
		"en-UK": "Day after tomorrow",
		"es-ES": "Pasada mañana",
		"fa-IR": "پس فردا",
		"it-IT": "Dopo domani",
		"ru-RU": "Послезавтра",
	},
	COMMAND_TEXT_THIS_WEEK: {
		"de-DE": "Diese Woche",
		"en-UK": "This week",
		"es-ES": "Esta semana",
		"fa-IR": "این هفته",
		"it-IT": "Questa settimana",
		"ru-RU": "На этой неделе",
	},
	COMMAND_TEXT_YES_IT_HAS_RETURN_DEADLINE: {
		"de-DE": "Ja, es hat eine Frist!",
		"en-UK": "Yes, it has a deadline!",
		"es-ES": "Sí, hay una fecha de devolución!",
		"fa-IR": "بله، دارای آخرین فرصت می باشد!",
		"it-IT": "Si, c'e' una data di scadenza",
		"ru-RU": "Да, есть срок возврата!",
	},
	COMMAND_TEXT_NO_IT_CAN_BE_RETURNED_ANYTIME: {
		"de-DE": "Nein, sobald möglich.",
		"en-UK": "No, whenever is fine.",
		"es-ES": "No, sin fecha límite.",
		"fa-IR": "خیر، هر زمانی مناسب است.",
		"it-IT": "No, nessuna scadenza",
		"ru-RU": "Нет, срок возврата не важен.",
	},
	COMMAND_TEXT_IT_CAN_BE_RETURNED_ANYTIME: {
		"de-DE": "Unbefristet",
		"en-UK": "Whenever is fine",
		"es-ES": "Cualquier día",
		"fa-IR": "هر زمانی مناسب است.",
		"it-IT": "No, Nessuna scadenza",
		"ru-RU": "Когда-нибудь",
	},
	COMMAND_TEXT_IN_FEW_MINUTES: {
		"de-DE": "In wenigen Minuten",
		"en-UK": "In few minutes",
		"es-ES": "En unos minutos",
		"fa-IR": "در چند دقیقه",
		"it-IT": "Fra alcuni minuti",
		"ru-RU": "Через минуту",
	},
	COMMAND_TEXT_IN_1_WEEK: {
		"de-DE": "In einer Woche",
		"en-UK": "In 1 week",
		"es-ES": "En una semana",
		"fa-IR": "ظرف یک هفته",
		"it-IT": "Fra una settimana",
		"ru-RU": "Через неделю",
	},
	COMMAND_TEXT_IN_1_MONTH: {
		"de-DE": "In einem Monat",
		"en-UK": "In 1 month",
		"es-ES": "En un mes",
		"fa-IR": "ظرف یک ماه",
		"it-IT": "Fra un mese",
		"ru-RU": "Через месяц",
	},
	COMMAND_TEXT_SET_DATE: {
		"de-DE": "Datum setzen",
		"en-UK": "Set date",
		"es-ES": "Establecer la fecha",
		"fa-IR": "تاریخ را تنظیم کنید",
		"it-IT": "Imposta la data",
		"ru-RU": "Задать дату",
	},
	COMMAND_TEXT_MONDAY: {
		"de-DE": "Montag",
		"en-UK": "Monday",
		"es-ES": "Lunes",
		"fa-IR": "دوشنبه",
		"it-IT": "Lunedi'",
		"ru-RU": "Понедельник",
	},
	COMMAND_TEXT_TUESDAY: {
		"de-DE": "Dienstag",
		"en-UK": "Tuesday",
		"es-ES": "Martes",
		"fa-IR": "سه شنبه",
		"it-IT": "Martedi'",
		"ru-RU": "Вторник",
	},
	COMMAND_TEXT_WEDNESDAY: {
		"de-DE": "Mittwoch",
		"en-UK": "Wednesday",
		"es-ES": "Miercoles",
		"fa-IR": "چهارشنبه",
		"it-IT": "Mercoledi'",
		"ru-RU": "Среда",
	},
	COMMAND_TEXT_THURSDAY: {
		"de-DE": "Donnerstag",
		"en-UK": "Thursday",
		"es-ES": "Jueves",
		"fa-IR": "پنج شنبه",
		"it-IT": "Giovedi'",
		"ru-RU": "Четверг",
	},
	COMMAND_TEXT_FRIDAY: {
		"de-DE": "Freitag",
		"en-UK": "Friday",
		"es-ES": "Viernes",
		"fa-IR": "جمعه",
		"it-IT": "Venerdi'",
		"ru-RU": "Пятница",
	},
	COMMAND_TEXT_SATURDAY: {
		"de-DE": "Samstag",
		"en-UK": "Saturday",
		"es-ES": "Sabado",
		"fa-IR": "شنبه",
		"it-IT": "Sabato",
		"ru-RU": "Суббота",
	},
	COMMAND_TEXT_SUNDAY: {
		"de-DE": "Sonntag",
		"en-UK": "Sunday",
		"es-ES": "Domingo",
		"fa-IR": "یکشنبه",
		"it-IT": "Domenica",
		"ru-RU": "Воскресенье",
	},
	COMMAND_TEXT_DO_NOT_SEND_RECEIPT: {
		"de-DE": "Keine Quittung schicken",
		"en-UK": "Do not send the receipt",
		"es-ES": "No enviar el recibo",
		"fa-IR": "رسید را ارسال نکنید",
		"it-IT": "Non inviare la ricevuta",
		"ru-RU": "Не отправлять квитанцию",
	},
	MESSAGE_TEXT_RECEIPT_WILL_NOT_BE_SENT: {
		"de-DE": "Du hast dich gegen eine Quittung entschieden.",
		"en-UK": "You've decided not to send the receipt.",
		"es-ES": "Has decidido no enviar el recibo",
		"fa-IR": "شما تصمیم گرفتید که رسید را ارسال نکنید.",
		"it-IT": "Hai scelto di non inviare la ricevuta",
		"ru-RU": "Вы решили не отправлять квитанцию.",
	},
	COMMAND_TEXT_I_HAVE_CHANGED_MY_MIND: {
		"de-DE": "Ich habe meine Meinung geändert",
		"en-UK": "I've changed my mind",
		"es-ES": "He cambiado de opinion",
		"fa-IR": "نظرم را عوض کردم.",
		"it-IT": "Ho cambiato idea",
		"ru-RU": "Я передумал(а)",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_TELEGRAM: {
		"de-DE": "Senden per Telegram",
		"en-UK": "Send by Telegram",
		"es-ES": "Enviar a través de Telegram",
		"fa-IR": "با تلگرام ارسال شود",
		"it-IT": "Invia tramite Telegram",
		"ru-RU": "Отправить через Telelgram",
	},
	COMMAND_TEXT_GET_LINK_FOR_RECEIPT_IN_TELEGRAM: {
		"en-UK": "Get link for a receipt in Telegram",
		"es-ES": "Obtener enlace para recibirlo en Telegram", // TODO(ES) verify
		"fa-IR": "دریافت پیوند برای دریافت در Telegram",      // TODO(FA) verify
		"it-IT": "Link per la ricevuta nel Telegram",         // TODO(IT)
		"ru-RU": "Ссылка для квитанции в Телеграмме",
		"de-DE": "Erhalten sie einen link für eine quittung in Telegram", // TODO(DE) verify
	},
	COMMAND_TEXT_COUNTERPARTY_HAS_NO_TELEGRAM: {
		"de-DE": "Senden per FB, WhatsApp, Viber, etc.",
		"en-UK": "Send by FB, WhatsApp, Viber, etc.",
		"es-ES": "Enviar a través de FB, WhatsApp, Viber, etc.",
		"fa-IR": "با فیسبوک، واتس آپ، وایبر و ... ارسال شود.",
		"it-IT": "Invia con FB, WhatsCrap, Viber, etc.",
		"ru-RU": "Отправить через Viber, VK, FB, ...",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_SMS: {
		"de-DE": "Senden per SMS",
		"en-UK": "Send by SMS",
		"es-ES": "Enviar a través de SMS",
		"fa-IR": "با پیام کوتاه ارسال شود",
		"it-IT": "Invia tramite SMS",
		"ru-RU": "Отправить через SMS",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_VK: {
		"de-DE": "Senden per VK.com",
		"en-UK": "Send throw VK.com",
		"es-ES": "Enviar vía VK.com",
		"fa-IR": "ارسال شود VK.com از طریق ",
		"it-IT": "Invia tramite VK.com (Facebook russo)",
		"ru-RU": "Отправить через ВКонтакте",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_OK: {
		"de-DE": "Senden per OK",
		"en-UK": "Send throw OK",
		"es-ES": "Enviar a través de OK",
		"fa-IR": "ارسال شود OK از طریق ",
		"it-IT": "Invia tramite OK",
		"ru-RU": "Отправить через Одноклассники",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_FB: {
		"de-DE": "Senden per Facebook",
		"en-UK": "Send throw Facebook",
		"es-ES": "Enviar a través de Facebook",
		"fa-IR": "از طریق فیسبوک ارسال شود.",
		"it-IT": "Invia tramite Facebook",
		"ru-RU": "Отправить через Facebook",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_TWT: {
		"de-DE": "Senden per Twitter",
		"en-UK": "Send throw Twitter",
		"es-ES": "Enviar a través de Twitter",
		"fa-IR": "از طریق توئیتر ارسال شود.",
		"it-IT": "Invia tramite Twitter",
		"ru-RU": "Отправить через Twitter",
	},
	COMMAND_TEXT_CANCEL_SENDING_RECEIPT_BY_TELEGRAM: {
		"de-DE": "Sendung der Quittung per Telegram abbrechen",
		"en-UK": "Cancel sending receipt by Telegram",
		"es-ES": "Cancelar el envío a través de Telegram",
		"fa-IR": "ارسال رسید با تلگرام کنسل شود",
		"it-IT": "Annulla l'invio tramite Telegram",
		"ru-RU": "Отменить отправку через Telegram",
	},
	COMMAND_TEXT_MAIN_MENU_TITLE: {
		"de-DE": "Menü /menu",
		"en-UK": "Main /menu",
		"es-ES": "Inicio /menú",
		"fa-IR": "/منو ی اصلی ",
		"it-IT": "Menu' /menu",
		"ru-RU": "Главное /меню",
	},
	MESSAGE_TEXT_NOTHING_TO_CANCEL: {
		"de-DE": "Nichts zum abbrechen.",
		"en-UK": "Nothing to cancel.",
		"es-ES": "No hay nada que anular.",
		"fa-IR": "چیزی برای کنسل شدن وجود ندارد",
		"it-IT": "Nulla da annullare.",
		"ru-RU": "Нечего отменять.",
	},
	MESSAGE_TEXT_TRANSFER_CREATION_CANCELED: {
		"de-DE": "Erstellung des Schuldscheins abgebrochen.",
		"en-UK": "Creation of debt record has been canceled.",
		"es-ES": "La creación del recordatorio se ha cancelado.",
		"fa-IR": "ایجاد سابقه بدهی کنسل شد.",
		"it-IT": "Creazione record annullata",
		"ru-RU": "Создание записи о долге отменено.",
	},
	COMMAND_TEXT_SHOW_ALL_CONTACTS: {
		"de-DE": "Zeige alle...",
		"en-UK": "Show all...",
		"es-ES": "Mostrar todo...",
		"fa-IR": "نمایش تمام موارد ...",
		"it-IT": "Mostra tutto...",
		"ru-RU": "Показать все...",
	},
	COMMAND_TEXT_CONTACTS: {
		"de-DE": "Kontakte",
		"en-UK": "Contacts",
		"es-ES": "Contactos",
		"fa-IR": "لیست تماس",
		"it-IT": "Сontatti",
		"ru-RU": "Контакты",
	},
	COMMAND_TEXT_REFRESH: {
		"de-DE": "Aktualisieren",
		"en-UK": "Refresh",
		"es-ES": "Recargar",
		"fa-IR": "تازه کردن",
		"it-IT": "Ricaricare",
		"ru-RU": "Обновить",
	},
	COMMAND_TEXT_ADD_YOUR_OWN_OPTION: {
		"de-DE": "Etwas anderes",
		"en-UK": "Something else",
		"es-ES": "Otra cosa",
		"fa-IR": "چیزی دیگر",
		"it-IT": "Qualcos'altro",
		"ru-RU": "Что-то другое",
	},
	MESSAGE_TEXT_REMINDER_ASK_IF_RETURNED: {
		"de-DE": "Wurde diese Schuld beglichen?",
		"en-UK": "Have this debt been returned?",
		"es-ES": "¿Se ha devuelto esta deuda?",
		"fa-IR": "آیا این بدهی بازپرداخت شده است؟",
		"it-IT": "Questo debito e' stato saldato?",
		"ru-RU": "Был ли этот долг возвращён?",
	},
	MESSAGE_TEXT_ASK_WHEN_TO_REMIND_AGAIN: {
		"de-DE": "Wann willst du wieder an diese Schuld erinnert werden?",
		"en-UK": "When should we remind you about this debt again?",
		"es-ES": "¿Cuándo recordarte de esta deuda otra vez?",
		"it-IT": "Quando devo ricordarti di questo debito?",
		"fa-IR": "چه زمانی لازم است مجدداً در مورد این بدهی به شما یادآوری نماییم؟",
		"ru-RU": "Когда вам напомнить об этом долге ещё раз?",
	},
	MESSAGE_TEXT_REPLIED_DEBT_RETURNED_FULLY: {
		"de-DE": "Du hast angegeben, dass diese Schuld vollständig beglichen ist.",
		"en-UK": "You've replied back that debt has been returned fully.",
		"es-ES": "Has confirmado que la deuda se ha saldado totalmente",
		"fa-IR": "شما پاسخ داده اید که بدهی به صورت کامل بازپرداخت شده است.",
		"it-IT": "Hai confermato che il debito e' stato saldato.",
		"ru-RU": "Вы ответили что долг возвращён полностью.",
	},
	MESSAGE_TEXT_DEBT_IS_RETURNED: {
		"de-DE": "Die Schuld ist vollständig beglichen.",
		"en-UK": "The debt has been returned fully.",
		"es-ES": "La deuda se ha saldado totalmente",
		"fa-IR": "بدهی به صورت کامل بازپرداخت شده است",
		"it-IT": "Il debito e' stato saldato.",
		"ru-RU": "Долг возвращён полностью.",
	},
	MESSAGE_TEXT_DETAILS_ARE_HERE: {
		"de-DE": "Details hier: %v",
		"en-UK": "Details here: %v",
		"es-ES": "Detalles aquí: %v",
		"fa-IR": "جزئیات در اینجا: %v",
		"it-IT": "Dettagli qui: %v",
		"ru-RU": "Подробности тут: %v",
	},
	MESSAGE_TEXT_REMINDER: {
		"de-DE": "Erinnerung",
		"en-UK": "Reminder",
		"es-ES": "Recordatorio",
		"fa-IR": "یادآور",
		"it-IT": "Promemoria",
		"ru-RU": "Напоминание",
	},
	MESSAGE_TEXT_REMINDER_SET: {
		"de-DE": "Erinnerung am: %v",
		"en-UK": "Reminder set for: %v",
		"es-ES": "Recordatorio establecito para: %v",
		"fa-IR": "یادآور تنظیم شده است برای: %v",
		"it-IT": "Imposta promemoria per: %v",
		"ru-RU": "Напоминание установлено на: %v",
	},
	MESSAGE_TEXT_REMINDER_DISABLED: {
		"de-DE": "Du hast die Erinnerung an diese Schuld deaktiviert.",
		"en-UK": "You've disabled reminders for this debt.",
		"es-ES": "Recordatorio para esta deuda se ha deshabilitado.",
		"fa-IR": "شما یادآور را برای این بدهی غیرفعال نموده اید.",
		"it-IT": "Hai disabilitato il promemoria per questo debito.",
		"ru-RU": "Напоминания об этом долге отключены.",
	},
	COMMAND_TEXT_REMINDER_ENABLE: {
		"de-DE": "Erinnerung aktivieren",
		"en-UK": "Turn-on reminder",
		"es-ES": "Recordatorio de encendido", // TODO(es) verify
		"fa-IR": "یادآوری روشن",              // TODO(fa) verify
		"it-IT": "Ricordo promozionale",      // TODO(it) verify
		"ru-RU": "Включить напоминание",
	},
	MESSAGE_TEXT_REMINDER_ALREADY_RESCHEDULED: {
		"de-DE": "Du wirst bereits erneut erinnert.",
		"en-UK": "You've already rescheduled this reminder.",
		"es-ES": "Recordatorio para esta deuda se ha reprogramado ya.",
		"it-IT": "Hai gia' impostato questo promemoria",
		"fa-IR": "شما قبلا به صورت مجدد این یادآور را زمانبندی نموده اید.",
		"ru-RU": "Напоминание об этом долге уже перенесено.",
	},
	COMMAND_TEXT_REMINDER_RETURNED_IN_FULL: {
		"de-DE": "Ja, vollständig beglichen",
		"en-UK": "Yes, returne in full",
		"es-ES": "Sí, devuelto totalmente",
		"fa-IR": "بله، بازپرداخت به صورت کامل",
		"it-IT": "Si, completamento saldato",
		"ru-RU": "Да, возвращено полностью",
	},
	COMMAND_TEXT_REMINDER_RETURNED_PARTIALLY: {
		"de-DE": "Teilweise beglichen",
		"en-UK": "Returned partially",
		"es-ES": "Devuelto parcialmente",
		"fa-IR": "تا اندازه ای بازپرداخت شده است",
		"it-IT": "Parzialmente saldato",
		"ru-RU": "Возврашено частично",
	},
	COMMAND_TEXT_REMINDER_NOT_RETURNED: {
		"de-DE": "Nicht beglichen",
		"en-UK": "Not returned",
		"es-ES": "No devuelto",
		"fa-IR": "بازپرداخت نشده است",
		"it-IT": "Debito non saldato",
		"ru-RU": "Не возвращено",
	},
	MESSAGE_TEXT_YOU_REPLIED: {
		"de-DE": "Beantwortet: %v",
		"en-UK": "You've replied: %v",
		"es-ES": "Has respondido: %v",
		"fa-IR": "شما پاسخ داده اید: %v",
		"it-IT": "Hai risposto: %v",
		"ru-RU": "Вы ответили: %v",
	},
	"book": {
		"de-DE": "buchen",
		"en-UK": "book",
		"es-ES": "libro",
		"fa-IR": "کتاب",
		"it-IT": "libro",
		"ru-RU": "книгу",
	},
	botsfw.MessageTextBotDidNotUnderstandTheCommand: {
		"de-DE": "\xF0\x9F\x98\xB3 Entschuldigung, aber ich habe deinen Befehl nicht verstanden. Vielleicht bin ich ein bisschen dumm...\n\nDu kannst zurück ins /menu",
		"en-UK": "\xF0\x9F\x98\xB3 Sorry, I did not understand your order. May be I'm a little bit dumb...\n\nYou can return to main /menu",
		"es-ES": "\xF0\x9F\x98\xB3 Disculpa, no he entendido tu orden. Tal vez soy un poco tonto...\n\nPuedes volver al Menu principal /menu",
		"fa-IR": "\xF0\x9F\x98\xB3 ببخشید، من دستور شما را نفهمیدم. احتمالا کمی کند ذهن هستم...\n\nشما میتوانید به /منو ی اصلی بازگردید",
		"it-IT": "\xF0\x9F\x98\xB3 Scusami ma non ho capito cosa vuoi. Sono ancora un po' sciocco...\n\nPuoi ritornare al Menu con /menu",
		"ru-RU": "\xF0\x9F\x98\xB3 Извините, я не понял вашу команду. Возможно я немного туповат...\n\nЧтобы начать сначала нажмите /menu",
	},
	COMMAND_TEXT_LANGUAGE: {
		"de-DE": "Sprache / Language",
		"en-UK": "Bot language",
		"es-ES": "Idioma / Language",
		"fa-IR": "زبان",
		"it-IT": "Lingua / Language",
		"ru-RU": "Язык / Language",
	},
	"/start": {
		"de-DE": "/start",
		"en-UK": "/start",
		"es-ES": "/comienzo",
		"fa-IR": "/شروع",
		"it-IT": "/start",
		"ru-RU": "/старт",
	},
	COMMAND_TEXT_DUE_RETURNS: {
		"de-DE": "Fällige Schulden",
		"en-UK": "Due returns",
		"es-ES": "Devoluciones",
		"fa-IR": "بازپرداخت بدهی",
		"it-IT": "Debiti",
		"ru-RU": "Предстоящие платежи",
	},
	MESSAGE_TEXT_OVERDUE_RETURNS_HEADER: {
		"de-DE": "<b>Überfällige Schulden:</b>",
		"en-UK": "<b>Overdue debts:</b>",
		"es-ES": "<b>Deudas atrasadas:</b>",
		"fa-IR": "<b>بدهی های معوق:</b>",
		"it-IT": "<b>Debiti in ritardo:</b>",
		"ru-RU": "<b>Просроченные долги:</b>",
	},
	MESSAGE_TEXT_DUE_RETURNS_HEADER: {
		"de-DE": "<b>Bald fällige Schulden:</b>",
		"en-UK": "<b>Closest debts to return:</b>",
		"es-ES": "<b>Deudas más cercanos que pagar:</b>",
		"fa-IR": "<b>نزدیک ترین بدهی برای بازپرداخت:</b>",
		"it-IT": "<b>Debiti in scadenza:</b>",
		"ru-RU": "<b>Ближайшие долги к возрату:</b>",
	},
	MESSAGE_TEXT_DUE_RETURNS_ROW_BY_USER: {
		"de-DE": "%v bekommt %v von dir, spätestens in %v",
		"en-UK": "%v expects %v from you in %v",
		"es-ES": "%v espera %v que devuelvas en %v",
		"it-IT": "%v aspetta %v da te entro il %v",
		"fa-IR": "%v انتظار دارد %v از شما در %v",
		"ru-RU": "%v ожидает от вас возврата %v через %v",
	},
	MESSAGE_TEXT_DUE_RETURNS_ROW_BY_COUNTERPARTY: {
		"de-DE": "%v gibt dir %v, spätestens in %v",
		"en-UK": "You expect %v to return %v to you in %v",
		"es-ES": "Estas esperando de %v que devuelva %v a ti en %v",
		"fa-IR": "شما انتظار دارید %v بازگرداند %v به شما در %v",
		"it-IT": "Stai aspettando %v che ti dia %v entro il %v",
		"ru-RU": "Вы ожидаете от %v возврата %v через %v",
	},
	MESSAGE_TEXT_DUE_RETURNS_EMPTY: {
		"de-DE": "Du hast keine Schulden mit Fälligkeitsdatum.",
		"en-UK": "You have no debts with set due date.",
		"es-ES": "No tienes deudas con la fecha señalada para devolver. ",
		"fa-IR": "شما بدهی ای با ثبت سررسید ندارید.",
		"it-IT": "Non hai debiti con una data di scadenza.",
		"ru-RU": "У вас нет долгов с назначеным сроком к возврату.",
	},
	COMMAND_TEXT_GAVE: {
		"de-DE": "Verleihen",
		"en-UK": "Gave",
		"es-ES": "Prestado por ti",
		"fa-IR": "قرض_دادن",
		"it-IT": "Credito",
		"ru-RU": "Дать",
	},
	COMMAND_TEXT_GOT: {
		"de-DE": "Anleihen",
		"en-UK": "Got",
		"es-ES": "Prestado a ti",
		"fa-IR": "قرض_گرفتن",
		"it-IT": "Debito",
		"ru-RU": "Взять",
	},
	COMMAND_TEXT_RETURN: {
		"de-DE": "Beglichen",
		"en-UK": "Return",
		"es-ES": "Devuelto",
		"fa-IR": "بازگشت",
		"it-IT": "Rientra",
		"ru-RU": "Вернуть",
	},
	COMMAND_TEXT_BALANCE: {
		"de-DE": "Ausstehend",
		"en-UK": "Balance",
		"es-ES": "Balance",
		"fa-IR": "تراز",
		"it-IT": "Bilancio",
		"ru-RU": "Баланс",
	},
	COMMAND_TEXT_SETTING: {
		"de-DE": "Einstellungen",
		"en-UK": "/Settings",
		"es-ES": "Ajustes",
		"fa-IR": "تنظیمات",
		"it-IT": "Settaggi",
		"ru-RU": "Настройки",
	},
	COMMAND_TEXT_HIGH_FIVE: {
		"de-DE": "Gib mir Fünf!",
		"en-UK": "High five!",
		"es-ES": "¡Choca esos cinco!",
		"fa-IR": "بزن قدش!",
		"it-IT": "Batti 5 bro!",
		"ru-RU": "Дать пять!",
	},
	COMMAND_TEXT_CHANGE_LANG: {
		"de-DE": "Sprache",
		"en-UK": "Language",
		"es-ES": "Idioma",
		"fa-IR": "زبان",
		"it-IT": "Lingua",
		"ru-RU": "Язык",
	},
	COMMAND_TEXT_HELP: {
		"de-DE": "Hilfe",
		"en-UK": "Help",
		"es-ES": "Ayuda",
		"fa-IR": "کمک",
		"it-IT": "Aiuto",
		"ru-RU": "Помощь",
	},
	COMMAND_TEXT_HISTORY: {
		"de-DE": "Verlauf",
		"en-UK": "History",
		"es-ES": "Cronología",
		"fa-IR": "پیشینه",
		"it-IT": "Cronologia",
		"ru-RU": "История",
	},
	COMMAND_TEXT_CANCEL: {
		"de-DE": "Abbrechen",
		"en-UK": "Cancel",
		"es-ES": "Cancelar",
		"fa-IR": "کنسل",
		"it-IT": "Annulla",
		"ru-RU": "Отменить",
	},
	COMMAND_TEXT_REFERRERS: {
		//"de-DE": "Abbrechen",
		"en-UK": "Referrers",
		//"es-ES": "Cancelar",
		//"fa-IR": "کنسل",
		//"it-IT": "Annulla",
		"ru-RU": "Нас рекомендуют",
	},
	MESSAGE_TEXT_HOW_TO_ADD_TG_CHANNEL: {
		"en-UK": `To add your channel to the list just write about us with a link as %v <code>&lt;-</code> replace <code>YOUR_CHANNEL</code> with your own channel.

It's better if you hide the link in HTML as:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

This should be rendered by Telegram clients as: <a href="%v">@%v</a>

Top 5 referrers for the last 100 new users will be shown here.`,

		"ru-RU": `Чтобы добавить ваш канал в этот список просто напишите об этом боте использую ссылку вида %v <code>&lt;-</code> замените <code>YOUR_CHANNEL</code> на ваш канал.

Будет лучше  если вы спрячете её в HTML как:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Такой код должен отобразиться в Телеграмме как: <a href="%v">@%v</a>

Топ-5 источников последних 100 пользователей будут показаны здесь.`,
	},
	BUTTON_TEXT_CANCEL: {
		"de-DE": "↩ Zurück",
		"en-UK": "↩ Cancel",
		"es-ES": "↩ Cancelar",
		"fa-IR": "↪ کنسل",
		"it-IT": "↩ Annulla",
		"ru-RU": "↩ Отменить",
	},
	BUTTON_TEXT_MAIN_MENU: {
		"de-DE": "↩ Hauptmenü",
		"en-UK": "↩ Main menu",
		"es-ES": "↩Menú principal",
		"ru-RU": "↩ Главное меню",
		"fa-IR": "↪ منوی اصلی",       // TODO(FA) verify
		"it-IT": "↩ Menu principale", // TODO(IT) verify
	},
	COMMAND_TEXT_SETTINGS_PRIMARY_CURRENCY: {
		"de-DE": "Hauptwährung",
		"en-UK": "Primary currency",
		"es-ES": "Moneda principal",
		"fa-IR": "واحد پول اولیه",
		"it-IT": "Valuta principale",
		"ru-RU": "Основная валюта",
	},
	COMMAND_TEXT_ADD_GROUP: {
		"de-DE": "Neue Gruppe", // TODO(DE) - change grom "New group" to "Add group"
		"en-UK": "Add group",
		"es-ES": "Añadir grupo",      // TODO(ES) verify
		"fa-IR": "اضافه کردن گروه",   // TODO(FA) verify
		"it-IT": "Aggiungere gruppo", // TODO(IT) verify
		"ru-RU": "Добавить группу",
	},
	COMMAND_TEXT_GROUPS: {
		"de-DE": "", // TODO(DE)
		"en-UK": "Groups",
		"es-ES": "", // TODO(ES) verify
		"fa-IR": "", // TODO(FA) verify
		"it-IT": "", // TODO(IT) verify
		"ru-RU": "Группы",
	},
	COMMAND_TEXT_BILLS: {
		"de-DE": "", // TODO(DE)
		"en-UK": "Bills",
		"es-ES": "", // TODO(ES) verify
		"fa-IR": "", // TODO(FA) verify
		"it-IT": "", // TODO(IT) verify
		"ru-RU": "Закрыть счёт",
	},
	COMMAND_TEXT_SETTLE_BILL: {
		"de-DE": "", // TODO(DE)
		"en-UK": "Settle bill",
		"es-ES": "", // TODO(ES) verify
		"fa-IR": "", // TODO(FA) verify
		"it-IT": "", // TODO(IT) verify
		"ru-RU": "Счета",
	},
	COMMAND_TEXT_SETTLE_BILLS: {
		"de-DE": "", // TODO(DE)
		"en-UK": "Settle bills",
		"es-ES": "", // TODO(ES) verify
		"fa-IR": "", // TODO(FA) verify
		"it-IT": "", // TODO(IT) verify
		"ru-RU": "Закрыть счета",
	},
	COMMAND_TEXT_INVITE_FIREND: {
		"de-DE": "Freund einladen",
		"en-UK": "Invite friend",
		"es-ES": "Invitar a un amigo",
		"fa-IR": "دوستی دعوت کن",   // TODO(FA) verify
		"it-IT": "Invita un amico", // TODO(IT) verify
		"ru-RU": "Пригласить друга",
	},
	COMMAND_TEXT_INVITE_MEMBER: {
		"de-DE": "Mitglied einladen",
		"en-UK": "Invite member",
		"es-ES": "Invitar miembro", // TODO(ES) verify
		"fa-IR": "دعوت از اعضا",    // TODO(FA) verify
		"it-IT": "Invita membro",   // TODO(IT) verify
		"ru-RU": "Пригласить участника",
	},
	COMMAND_TEXT_NEW_BILL: {
		"de-DE": "Neue Rechnung",
		"en-UK": "New bill",
		"es-ES": "Nuevo proyecto de ley",  // TODO(ES) verify
		"fa-IR": "لایحه جدید",             // TODO(FA) verify
		"it-IT": "Nuovo disegno di legge", // TODO(IT) verify
		"ru-RU": "Новый счёт",
	},
	COMMAND_TEXT_NEW_FUNDRAISING: {
		"en-UK": "New fundraising",
		"es-ES": "Nueva recaudación de fondos", // TODO(ES) verify
		"fa-IR": "جمع آوری پول جدید",           // TODO(FA) verify
		"it-IT": "Nueva recaudación de fondos", // TODO(IT) verify
		"ru-RU": "Новый сбор средств",
	},
	COMMAND_TEXT_NEW_COUNTERPARTY: {
		"de-DE": "neuer Kontakt",
		"en-UK": "Add new",
		"es-ES": "Añadir",
		"fa-IR": "اضافه کردن مورد جدید",
		"it-IT": "Aggiungi nuovo",
		"ru-RU": "Добавить",
	},
	MESSAGE_TEXT_LOGIN_CODE: {
		"de-DE": "Dein Code um dich an der App anzumelden: <b>%v</b>",
		"en-UK": "Your code for signing in to app: <b>%v</b>",
		"es-ES": "Tu código para entrar en la app: <b>%v</b>",
		"fa-IR": "کد شما برای ورود به برنامه: <b>%v</b>",
		"it-IT": "Il tuo codice per accedere all'app e': <b>%v</b>",
		"ru-RU": "Ваш код для входа в приложение: <b>%v</b>",
	},
	MESSAGE_TEXT_ASK_NEW_COUNTERPARTY_NAME: {
		"de-DE": `Bitte gib einen Namen für den neuen Kontakt ein:
		Du kannst in eintippen oder aus deinem Adressbuch wählen (<i>mit dem "Büroklammer"-Symbol und dann Kontakt</i>).

		<i>Send '.' to cancel</i>`,

		"en-UK": `Please enter a name for the new contact:
		You can type manually or choose from your address book (<i>through "clip" icon</i>).

		<i>Send '.' to cancel</i>`,

		"es-ES": `Escribe un nombre para el nuevo contacto:
		Puedes escribirlo o elegirlo de tus contactos (<i>a traves del icono "clip"</i>).

		<i>Enviar '.' para anular</i>`,

		"fa-IR": `لطفا برای مخاطب جدید یک نام وارد کنید:
		میتوانید به صورت دستی تایپ نموده یا از لیست مخاطبین خود انتخاب نمایید (<i>throw "clip" icon</i>).

		<i>Send '.' برای کنسل کردن</i>`,

		"it-IT": `Inserisci un nome per il nuovo contatto:
		Puoi digitarlo o sceglierlo dalla tua rubrica (<i>attraverso l'icona "clip"</i>).

		<i>Digita '.' ed invia per annullare</i>`,

		"ru-RU": `<b>Имя для нового контакта</b>
		Напишите сами или выберите из своей адресной книги (<i>через иконку "скрепка"</i>).
		<i>Отправьте '.' для отмены</i>`,
	},
	MESSAGE_TEXT_TRANSFER_IS_CREATING: {
		"de-DE": "Transferiere...",
		"en-UK": "Creating transfer...",
		"es-ES": "Estoy creando la nueva nota...",
		"fa-IR": "ایجاد انتقال ...",
		"it-IT": "Sto creando la nuova voce...",
		"ru-RU": "Создаю запись...",
	},
	COMMAND_TEXT_PLEASE_WAIT: {
		"de-DE": "Bitte warten",
		"en-UK": "Please wait",
		"es-ES": "Espera, por favor",
		"fa-IR": "لطفا صبر کنید",
		"it-IT": "Aspetta per favore",
		"ru-RU": "Пожалуйста подождите",
	},
	MESSAGE_TEXT_PLEASE_WAIT: {
		"de-DE": "Bitte warten...",
		"en-UK": "Please wait...",
		"es-ES": "Espera, por favor...",
		"it-IT": "Aspetta per favore...",
		"fa-IR": "لطفا صبر کنید ...",
		"ru-RU": "Пожалуйста подождите...",
	},
	MESAGE_TEXT_CREATING_BILL: {
		"de-DE": "Rechnung erstellen",
		"en-UK": "Creating bill",
		"es-ES": "Crear factura",        // TODO(ES) verify
		"fa-IR": "ایجاد لایحه",          // TODO(FA) verify
		"it-IT": "Creazione di fattura", // TODO(IT) verify
		"ru-RU": "Создаётся счёт",
	},
	MESSAGE_TEXT_ASK_BILL_CURRENCY: {
		"de-DE": "In welcher Währung ist die Rechnung?",
		"en-UK": "What currency this bill in?",
		"es-ES": "¿Qué moneda esta factura?",       // TODO(es) verify
		"it-IT": "Quale moneta ha questa fattura?", // TODO(it) verify
		"fa-IR": "این لایحه چه ارزشی دارد؟",        // TODO(fa) verify
		"ru-RU": "В какой валюте этот счёт?",
	},
	MESSAGE_TEXT_ASK_BILL_PAYER: {
		"de-DE": "Wer bezahlt die Rechnung?", // TODO(de) verify
		"en-UK": "Who paid for the bill?",
		"es-ES": "¿Quién paga la cuenta?",           // TODO(es) verify
		"it-IT": "Chi paga il conto?",               // TODO(it) verify
		"fa-IR": "چه کسی برای این لایحه می پردازد؟", // TODO(fa) verify
		"ru-RU": "Кто оплатил счёт?",
	},
	MESSAGE_TEXT_SELF_ACKNOWLEDGEMENT: {
		"de-DE": "%v muss dem zustimmen",
		"en-UK": "Acknowledgement is expected from %v",
		"es-ES": "Se espera la confirmación de %v",
		"it-IT": "Conferma in attesa da %v",
		"fa-IR": "انتظار تصدیق می رود از %v",
		"ru-RU": "Подтверждение ожидается от %v",
	},
	MESSAGE_TEXT_TRANSFER_ACCEPTED_BY_YOU: {
		"de-DE": "Du hast dem zugestimmt.",
		"en-UK": "You've accepted this transaction.",
		"es-ES": "Has confirmado esta transacción",
		"fa-IR": ".شما این تراکنش را قبول کردید ",
		"it-IT": "Hai accettato il debito/credito.",
		"ru-RU": "Вы подтвердили эту транзакцию.",
	},
	MESSAGE_TEXT_TRANSFER_DECLINED_BY_YOU: {
		"de-DE": `Du hast dieser Anfrage nicht zugestimmt.
		Der Vorgang wird zurückgestellt und die Gegenpartei benachrichtigt.`,
		"en-UK": `You do not agree with this transaction.
                The transaction will not be deleted but the creator will be notified.`,
		"es-ES": `No estas de acuerdo con la transacción.
		La transacción NO será cancelada, pero el creador será notificado.`,
		"fa-IR": ".شما این تراکنش را رد کردید", //TODO(FA)
		"it-IT": `Hai rifiutato il debito/credito.`,
		"ru-RU": `Вы НЕ согласны с этой транзакцией.

Сама транзакция НЕ будет отменена, но создатель будет оповещён.`,
	},
	MESSAGE_TEXT_TRANSFER_ACCEPTED_BY_COUNTERPARTY: {
		"de-DE": "%v hat deiner Anfrage <b>zugestimmt</b>:",
		"en-UK": "%v accepted your transaction:",
		"es-ES": "%v ha aceptado tu transacción",
		"fa-IR": ": تراکنش شمارا تایید کرد %v ",
		"it-IT": "%v ha accettato il tuo credito/debito:",
		"ru-RU": "%v подтвердил(a) вашу транзакцию:",
	},
	MESSAGE_TEXT_TRANSFER_DECLINED_BY_COUNTERPARTY: {
		"de-DE": "%v hat deine Anfrage <b>abgelehnt</b>. Wenn die Sache besprochen ist, kann die Anfrage erneut gesendet werden.",
		"en-UK": "%v did not agree with your transaction. The transaction is not canceled but you may want to discuss it.", //TODO(FA)
		"es-ES": "%v ha confirmado tu transacción.  La transacción no ha sido cancelada, igual mejor comentarlo.",          //TODO(ES)
		"fa-IR": "تراکنش شما را رد کرد  %v declined your transaction.",
		"it-IT": "%v ha rifiutato il tuo credito/debito.  The transaction is not canceled but you may want to discuss it.", //TODO(IT)
		"ru-RU": "%v <b>НЕ</b> подтвердил(a) вашу транзакцию. Транзакция не отменена, но возможно вам стоит это обсудить.",
	},
	COMMAND_TEXT_SUBSCRIBE_TO_APP: {
		"de-DE": "Ich will die App!",
		"en-UK": "I want the app!",
		"es-ES": "¡Quiero la aplicación!",
		"fa-IR": "!من برنامه را می خواهم",
		"it-IT": "Voglio l'aplicazione!",
		"ru-RU": "Хочу приложение!",
	},
	COMMAND_TEXT_I_AM_FINE_WITH_BOT: {
		"de-DE": "Mir reicht der Bot!",
		"en-UK": "I'm fine with just the bot!",
		"es-ES": "¡Estoy satisfecho con este bot!",
		"fa-IR": "! ربات به تنهایی برای من کافی است",
		"it-IT": "Mi accontento del bot per ora!",
		"ru-RU": "Меня вполне устраивает бот!",
	},
	MESSAGE_TEXT_SUBSCRIBED_TO_APP: {
		"de-DE": "Du wirst darüber informiert, wenn die App zum Download zur Verfügung steht.",
		"en-UK": "We'll let you know once the app is available for download.",
		"es-ES": "Te avisamos cuando la aplicación esté disponible para descargarla",
		"fa-IR": ".وقتی برنامه برای دانلود دردسترس بود به شما اطلاع می دهیم",
		"it-IT": "Ti faremo sapere non appena l'applicazione sara' disponibile al download.",
		"ru-RU": "Мы сообщим вам когда приложение будет доступно для загруки.",
	},
	MESSAGE_TEXT_NOT_INTERESTED_IN_APP: {
		"de-DE": "Gut, wir sind froh, dass dir der Bot reicht und wir uns mit der App nicht beeilen müssen.",
		"en-UK": "Well, we are happy our bot is good enough and there is no need to download an app.",
		"es-ES": "Bueno, estamos contentos de que te haya gustado nuestro bot y no hace falta descargar ninguna otra aplicación",
		"fa-IR": ".خب، ما خوشحال هستیم که ربات برای شما کافی است و نیازی به دانلود برنامه نیست",
		"it-IT": "Bene, siamo contenti che il nostro bot sia di tuo gradimento e non hai bisogno di scaricare l'applicazione.",
		"ru-RU": "Что ж, мы рады что вас устраивает наш бот и нет необходимости загружать приложение.",
	},
	MESSAGE_TEXT_YOUR_AD_COULD_BE_HERE: {
		"de-DE": "Hier könnte <a href>ihre Werbung</a> stehen",
		"en-UK": "You can <a href>advertise here</a>",
		"es-ES": "Aquí se puede <a href>publicar un anuncio</a>",
		"fa-IR": "شما میتوانید <a href>در اینجا تبلیغ کنید</a>",
		"it-IT": "Puoi <a href>pubblicizzare qui</a>",
		"ru-RU": "Здесь можно <a href>разместить рекламу</a>",
	},
	MESSAGE_TEXT_YOUR_ABOUT_ADS: {
		"de-DE": `🤖: Ich hin ein guter Roboter - klar. Aber manchmal kommt es besser eine eigene App für etwas zu haben. Es ist noch nicht ganz fertig, aber falls du schonmal reinschauen willst: <a href="https://debtstracker.io/de/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/en/</a>

	Möchtest du daran erinnert werden, wenn die App rauskommt?`,
		"en-UK": `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/en/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/en/</a>

	Do you want to get an invite when it gets released?`,

		"es-ES": `🤖: Claro que soy un robot encantador, pero más comodo usar una aplicación especial.No esta disponible ya pero se puede ver como será: <a href = "https://debtstracker.io/es/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/en/</a>

	¿Quieres que te avisemos cuando esté lista?`,

		"fa-IR": `🤖: مطمئناً من روبات خوبی هستم. اما بعضی وقت هاساده تر و مناسب تر است که از یک برنامه به خوبی تخصصی شده استفاده شود، این برنامه هنوز برای استفاده عموم آماده نیست ولی می توانید چک کنید که چگونه به نظر خواهد رسید: <a href="https://debtstracker.io/fa/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/en/</a>

	آیا می خواهید وقتی منتشر شد دعوتنامه ای دریافت کنید؟`,

		"it-IT": `🤖: Di sicuro son un bravo bot, ma alcune volte e' piu' conveniente usare un'applicazione specializzata. Non e' ancora pronta per la pubblicazione ma puoi controllare l'avanzamento a questo indirizzo: <a href="https://debtstracker.io/it/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/it/</a>

	Vuoi essere invitato non appena viene rilasciata?`,

		"ru-RU": `🤖: Я конечно обоятельный робот, но пользоваться специализированным приложением бывает удобнее. Оно ещё не готово для общего доступа, но уже сейчас можно посмотреть как будет выглядеть: <a href="https://debtstracker.io/ru/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/ru/</a>

		Хотите получить оповещение когда оно выйдет?`,
	},
	MESSAGE_TEXT_INVALID_FLOAT: {
		"de-DE": "Entschuldigung, aber du kannst nur Zahlen für Menge oder Wert wählen (<i>mit zwei Nachkommastellen</i>).",
		"en-UK": "Sorry, but you can use just numbers as amount/quantity (<i>with up to 2 digits after point</i>).",
		"es-ES": "Lo siento, solo puedes utilizar numeros como importe/cantidad (<i>con un maximo de 2 dígitos despues de la coma</i>).",
		"fa-IR": "ببخشید، اما شما تنها میتوانید از اعداد بعنوان مقادیر / اندازه ها استفاده کنید (<i>با دو رقم اعشار</i>).",
		"it-IT": "Spiacente, puoi utilizzare solo numeri (<i>con un massimo di 2 numeri dopo il punto</i>).",
		"ru-RU": "Извините, но вы можете использовать только числа в качестве суммы/количества (<i>до 2х знаков после запятой</i>).",
	},
	MESSAGE_TEXT_ASK_LENDING_TYPE: {
		"de-DE": "<b>Was hast du jemanden geliehen?</b>",
		"en-UK": "<b>What did you lend to someone?</b>",
		"es-ES": "<b>¿Qué has prestado?</b>",
		"fa-IR": "<b> چه چیزی به کسی قرض داده اید؟</b>",
		"it-IT": "<b>Cos'hai prestato?</b>",
		"ru-RU": "<b>Что вы дали в долг?</b>",
	},
	MESSAGE_TEXT_CHOOSE_CURRENCY: {

		"de-DE": `Bitte wähle <a>eine Währung aus der Liste</a>.

	Falls die Standardoptionen nicht reichen, sende mir einen Text. Zum Beispiel: <i>Äpfel</i>".`,

		"en-UK": `Please choose from the options below or <a>select a currency from the list</a>.

	If standard options are not enough simply send a text.For example: "<i>apple</i>".`,

		"es-ES": `Elige del menú abajo de la pantalla o <a>selecciona la moneda de la lista</a>.

	Si no encuentras la opción correcta simplemente envía un texto. Por ejemplo: "<i>manzana</i>".`,

		"fa-IR": `لطفا از بین گزینه های زیر انتخاب کنید یا <a>یک واحد پولی از لیست انتخاب کنید</a>.

	اگر گزینه های استاندارد کافی نبودند به سادگی یک متن بفرستید ، برای مثال:. "<i>سیب</i>".`,

		"it-IT": `Scegli dalle opzioni qui sotto o <a>seleziona una valuta dalla lista</a>.

	Se le opzioni standard non bastano semplicemente invia un testo.Per esempio: "<i>mele</i>".`,

		"ru-RU": `Выберите из меню внизу экрана или <a>выберите валюту из списка</a>.

	Если ни один из стандартных вариантов не подходит просто напишите текстом.Например: "<i>яблоко</i>".`,
	},
	MESSAGE_TEXT_ASK_LENDING_AMOUNT: {
		"de-DE": "Wie viel <b>%v</b> hast du verliehen?\n(<i>Sende '.' zum Abbrechen</i>)",
		"en-UK": "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)",
		"es-ES": "Cuanto(s) <b>%v</b> has prestado\n(<i>enviar '.' para cancelar</i>)",
		"fa-IR": "چه مقدار <b>%v</b> قرض داده اید؟\n(<i>ارسال '.' برای کنسل کردن</i>)",
		"it-IT": "Quanti <b>%v</b> hai prestato?\n(<i>Digita '.' ed invia per annullare</i>)",
		"ru-RU": "Сколько <b>%v</b> вы дали в долг?\n(<i>отправьте '.' чтобы отменить</i>)",
	},
	MESSAGE_TEXT_ASK_LENDING_COUNTERPARTY: {
		"de-DE": "Wer hat sich <b>%v</b> von dir geliehen?\n(<i>Sende '.' zum Abbrechen</i>)",
		"en-UK": "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		"es-ES": "A quién has prestado <b>%v</b>?\n(<i>enviar '.' para cancelar</i>)",
		"fa-IR": "چه کسی از شما <b>%v</b> قرض گرفته است؟\n(<i>ارسال '.' برای کنسل کردن</i>)",
		"it-IT": "Chi e' in debito di <b>%v</b> con te?\n(<i>Digita '.' ed invia per annullare</i>)",
		"ru-RU": "Кому вы дали в долг <b>%v</b>?\n(<i>отправьте '.' чтобы отменить</i>)",
	},
	MESSAGE_TEXT_ASK_BORROWING_TYPE: {
		"de-DE": "Was hast du dir geliehen?",
		"en-UK": "What did you lend?",
		"es-ES": "¿Qué te han prestado?",
		"fa-IR": "چه چیزی قرض گرفته اید؟",
		"it-IT": "Cosa ti hanno prestato?",
		"ru-RU": "Что вы взяли в долг?",
	},
	MESSAGE_TEXT_ASK_BORROWING_AMOUNT: {
		"de-DE": "Wie viel <b>%v</b> hast du geliehen?\n(<i>Sende '.' zum Abbrechen</i>)",
		"ru-RU": "Сколько <b>%v</b> вы взяли в долг?\n(<i>отправьте '.' чтобы отменить</i>)",
		"es-ES": "¿Cuánto <b>%v</b> has prestado?\n(<i>enviar '.' para cancelar</i>)",
		"en-UK": "How much <b>%v</b> did you borrow?\n(<i>send '.' to cancel</i>)",
		"fa-IR": "چه مقدار <b>%v</b> قرض گرفته اید؟\n(<i>ارسال '.' برای کنسل کردن</i>)",
		"it-IT": "Quanti <b>%v</b> ti hanno prestato?\n(<i>Digita '.' ed invia per annullare</i>)",
	},
	MESSAGE_TEXT_ASK_BORROWING_COUNTERPARTY: {
		"de-DE": "Wer hat dir <b>%v</b> geliehen?\n(<i>Sende '.' zum Abbrechen</i>)",
		"en-UK": "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		"es-ES": "¿Quién te ha prestado <b>%v</b>?\n(<i>enviar '.' para cancelar</i>)",
		"fa-IR": "چه کسی به شما <b>%v</b> قرض داده است؟ \n(<i>ارسال '.' برای کنسل کردن</i>)",
		"it-IT": "Chi ti ha prestato <b>%v</b>?\n(<i>Digita '.' ed invia per annullare</i>)",
		"ru-RU": "У кого вы взяли в долг <b>%v</b>?\n(<i>отправьте '.' чтобы отменить</i>)",
	},
	MESSAGE_TEXT_YOU_CAN_SEND_RECEIPT: {
		"de-DE": "Soll eine <a receipt>Quittung</a> an <a counterparty>%v</a> gesendet werden?",
		"en-UK": "Should we send a <a receipt>receipt</a> to <a counterparty>%v</a>?",
		"es-ES": "¿Debo enviar <a receipt> el recibo</a> a <a counterparty>%v</a>?",
		"fa-IR": "آیا لازم است ماارسال کنیم یک <a receipt>رسید</a> به <a counterparty>%v</a>?",
		"it-IT": "Devo inviare una <a receipt>notifica</a> a <a counterparty>%v</a>?",
		"ru-RU": "Отправить <a receipt>квитанцию</a> для <a counterparty>%v</a>?",
	},
	MESSAGE_TEXT_YOU_CAN_SEND_RECEIPT_TO_YOURSELF_BY_SMS: {
		"de-DE": "Entschuldigung, aber eine Quittung selber per SMS zu schicken ist im Moment noch nicht möglich. Aber dafür geht es mit %v.",
		"en-UK": "Sorry, sending a receipt to yourself by SMS is not available at the moment. You can send it to %v though.",
		"es-ES": "Lo siento, el envio del recibo a ti mismo a través de SMS en este momento está desactivado. Pero lo puedes enviar a %v.",
		"fa-IR": "متاسفم، درحال حاضرارسال یک رسید به خودتان بوسیله پیام کوتاه امکان پذیر نیست. شما میتوانید آنرا ارسال کنید به  %v از طریق.",
		"it-IT": "Spiacente ma inviarsi da soli una notifica tramite SMS non e' al momento disponibile. Pero' puoi inviarla a %v.",
		"ru-RU": "К сожалению отправка квитанцию себе по СМС в данный момент отключена. Но вы можете отправить её для %v.",
	},
	MESSAGE_TEXT_RECEIPT_IS_SENDING_BY_TELEGRAM: {
		"de-DE": "Die Quittung wird %v per Telegram geschickt...",
		"en-UK": "We are sending receipt to %v by Telegram...",
		"es-ES": "El recibo está enviando a%v a través de Telegram…",
		"fa-IR": "مادرحال ارسال رسید به %v از طریق تلگرام هستیم...",
		"it-IT": "Sto inviando la notifica a %v tramite Telegram...",
		"ru-RU": "Отправляем для %v извещение через Telegram...",
	},
	DAY: {
		"en-UK": "%v day",
		"ru-RU": "%v день",
	},
	DAYS_234: {
		"en-UK": "%v days",
		"ru-RU": "%v дня",
	},
	DAYS: {
		"en-UK": "%v days",
		"ru-RU": "%v дней",
	},
	MESSAGE_TEXT_INTEREST_PLEASE_SPECIFY_PERIOD: {
		"en-UK": "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?",
		"ru-RU": "Пожалуйста укажите также процентный период, т.е. уточните %%v%% это процент за какое количество дней?",
	},
	MESSAGE_TEXT_INTEREST: {
		"en-UK": "<b>Interest</b>: %v%% per %v",
		"ru-RU": "<b>Ставка</b>: %v%% за %v",
	},
	MESSAGE_TEXT_INTEREST_MIN_PERIOD: {
		"en-UK": "minimum period %v",
		"ru-RU": "минимальный период %v",
	},
	MESSAGE_TEXT_RECEIPT_NEW_DEBT_FROM_USER: {
		"de-DE": "{{.Counterparty}} schuldet dir {{.Amount}} .",
		"en-UK": "{{.Counterparty}} borrowed from you {{.Amount}}.",
		"es-ES": "{{.Counterparty}} prestado por tí {{.Amount}}.",
		"fa-IR": "{{.Counterparty}} از شما {{.Amount}} قرض گرفته است .",
		//"it-IT": "{{.Counterparty}} ha preso in prestito da te {{.Amount}}.",
		"it-IT": "{{.Counterparty}} e' in debito di {{.Amount}} con te.",
		"ru-RU": "{{.Counterparty}} взял(а) в долг {{.Amount}}.",
	},
	MESSAGE_TEXT_RECEIPT_NEW_DEBT_TO_USER: {
		"de-DE": "{{.Counterparty}} hat dir {{.Amount}} geliehen.",
		"en-UK": "{{.Counterparty}} lended to you {{.Amount}}.",
		"es-ES": "{{.Counterparty}} prestado a mí {{.Amount}}.",
		"fa-IR": "{{.Counterparty}} به شما {{.Amount}} قرض داده است .",
		"it-IT": "{{.Counterparty}} ti ha prestato {{.Amount}}.",
		"ru-RU": "{{.Counterparty}} дал(а) вам в долг {{.Amount}}.",
	},
	MESSAGE_TEXT_RECEIPT_RETURN_FROM_USER: {
		"de-DE": "Du hast {{.Amount}} an {{.Counterparty}} beglichen.",
		"en-UK": "You returned {{.Amount}} to {{.Counterparty}}.",
		"es-ES": "Has devuelto {{.Amount}} a {{.Counterparty}}.",
		"fa-IR": "شما بازگردانده اید {{.Amount}} به {{.Counterparty}}.",
		"it-IT": "Hai ridato {{.Amount}} a {{.Counterparty}}.",
		"ru-RU": "Вы вернули долг - {{.Counterparty}} получил(а) {{.Amount}}.",
	},
	MESSAGE_TEXT_RECEIPT_RETURN_TO_USER: {
		"de-DE": "{{.Counterparty}} hat dir {{.Amount}} beglichen.",
		"en-UK": "{{.Counterparty}} returned to you {{.Amount}}.",
		"es-ES": "{{.Counterparty}} te ha devuelto {{.Amount}}.",
		"fa-IR": "{{.Counterparty}} به شما بازپرداخت کرده است {{.Amount}}.",
		"it-IT": "{{.Counterparty}} ti ha ridato {{.Amount}}.",
		"ru-RU": "{{.Counterparty}} вернул вам {{.Amount}}.",
	},
	MESSAGE_TEXT_TRANSFER_ALREADY_FULLY_RETURNED: {
		"de-DE": "Diese Schuld ist bereits vollständig beglichen.",
		"en-UK": "This debts is already fully returned.",
		"es-ES": "Esta deuda se ha devuelta totalmente.",
		"it-IT": "Questi debiti sono già completamente restituiti.", // TODO(IT) verify
		"fa-IR": "این بدهی ها در حال حاضر به طور کامل بازگشته است.", // TODO(FA) verify
		"ru-RU": "Этот долг уже полностью возвращён.",
	},
	MESSAGE_TEXT_RECEIPT_ALREADY_RETURNED_AMOUNT: {
		"de-DE": "Bereits beglichen: {{.Amount}}.",
		"en-UK": "Already returned: {{.Amount}}.",
		"es-ES": "Se ha devuelto ya: {{.Amount}}.",
		"fa-IR": "قبلا برگشت: {{.Amount}}.",     // TODO(FA) - verify
		"it-IT": "Già restituito: {{.Amount}}.", // TODO(IT) - verify
		"ru-RU": "Уже возвращено: {{.Amount}}.",
	},
	MESSAGE_TEXT_RECEIPT_OUTSTANDING_AMOUNT: {
		"de-DE": "Ausstehend: {{.Amount}}.",
		"en-UK": "Outstanding: {{.Amount}}.",
		"es-ES": "Falta devolver: {{.Amount}}.",
		"fa-IR": "برجسته: {{.Amount}}",  // TODO(FA) verify
		"it-IT": "Inevaso: {{.Amount}}", // TODO(IT) verify
		"ru-RU": "Осталось вернуть: {{.Amount}}.",
	},
	MESSAGE_TEXT_DUE_ON: {
		"de-DE": "<b>Fällig am</b>: %v",
		"en-UK": "<b>Return till</b>: %v",
		"es-ES": "<b>Devolver hasta</b>: %v",
		"fa-IR": "<b>بازگردانده شود تا</b>: %v",
		"it-IT": "<b>Scadenza</b>: %v",
		"ru-RU": "<b>Вернуть до</b>: %v",
	},
	MESSAGE_TEXT_NOTE: {
		"de-DE": "Notiz",
		"en-UK": "Note",
		"es-ES": "Nota",
		"fa-IR": "نکته",
		"it-IT": "Nota",
		"ru-RU": "Заметка",
	},
	MESSAGE_TEXT_COMMENT: {
		"de-DE": "Bemerkung",
		"en-UK": "Comment",
		"es-ES": "Comentario",
		"fa-IR": "شرح",
		"it-IT": "Commenti",
		"ru-RU": "Комментарий",
	},
	MESSAGE_TEXT_LOGIN_TO_WEB_APP: {
		"de-DE": `<a>Hier klicken</a>, um sich an der Wep-App anzumelden.`,
		"en-UK": `Click to <a>sign in</a> to web-app.`,
		"es-ES": `Haz click para <a>acceder</a>la web-app.`,
		"it-IT": "Fai clic per <a>accedi</a> per app web.",
		"fa-IR": `کلیک کنید تا <a>وارد شوید</a> برنامه وب.`,
		"ru-RU": `Перейдите по <a>ссылке</a> чтобы запустить web-приложение.`,
	},
	MESSAGE_TEXT_DO_YOU_LIKE_OUR_BOT: {
		"de-DE": "Magst du @{{bot}}?",
		"en-UK": "Do you like @{{bot}}?",
		"es-ES": "¿Te gusta @{{bot}}?",
		"fa-IR": "آیا می پسندید @{{bot}}?",
		"it-IT": "Divertito con @{{bot}}?",
		"ru-RU": "Вам нравится @{{bot}}?",
	},
	COMMAND_TEXT_YES_EXCLAMATION: {
		"de-DE": "%v Ja!",
		"en-UK": "%v Yes!",
		"es-ES": "%v ¡Sí!",
		"fa-IR": "بله! %v",
		"it-IT": "%v Si!",
		"ru-RU": "%v Да!",
	},
	COMMAND_TEXT_YES: {
		"de-DE": "%v Ja",
		"en-UK": "%v Yes",
		"es-ES": "%v Sí",
		"it-IT": "%v Si",
		"fa-IR": "بله %v",
		"ru-RU": "%v Да",
	},
	COMMAND_TEXT_NO: {
		"de-DE": "%v Nein",
		"en-UK": "%v No",
		"es-ES": "%v No",
		"it-IT": "%v No",
		"fa-IR": "خیر %v",
		"ru-RU": "%v Нет",
	},
	COMMAND_TEXT_NOT_TOO_MUCH: {
		"de-DE": "%v Nicht so sehr",
		"en-UK": "%v Not too much",
		"es-ES": "%vNo mucho",
		"it-IT": "%v Non troppo",
		"fa-IR": "نه خیلی زیاد %v",
		"ru-RU": "%v Не очень",
	},
	COMMAND_TEXT_FEEDBACK: {
		"de-DE": "/Bewertung",
		"en-UK": "/Feedback",
		"es-ES": "/Respuesta",
		"it-IT": "/Risposta",
		"fa-IR": "/بازخورد",
		"ru-RU": "/Отзыв",
	},
	COMMAND_TEXT_WRITE_FEEDBACK: {
		"de-DE": "%v Bewertung schreiben",
		"en-UK": "%v Write feedback",
		"es-ES": "%v Escribir un comentario",
		"it-IT": "%v Scrivi commenti",
		"fa-IR": "ارسال بازخورد %v",
		"ru-RU": "%v Написать отзыв",
	},
	MESSAGE_TEXT_THANKS: {
		"de-DE": "🙏 Danke!",
		"en-UK": "🙏 Thanks!",
		"es-ES": "🙏 ¡Gracias!",
		"fa-IR": "🙏 تشکر!",
		"it-IT": "🙏 Grazie!",
		"ru-RU": "🙏 Спасибо!",
	},
	MESSGE_TEXT_DEBT_ERROR_FIXED_START_OVER: {
		"de-DE": "🙏 Entschuldigung, da gab es einen Fehler. Er wird bald behoben, aber du musst nochmal neu anfangen.",
		"en-UK": "🙏 Sorry, there was an error. It has been fixed but please re-enter your data for this debt.",
		"es-ES": "🙏 Lo siento, ha salido un error. Lo ha arreglado, pero para esta deuda hay que introducir los datos de nuevo. ",
		"ru-RU": "🙏 Извините, у нас была ошибка. Она была исправлено, но потребуется внести данные для этого долга заново.",
	},
	MESSAGE_TEXT_PLEASE_SEND_TEXT: {
		"de-DE": "Bitte senden sie einen Text.",
		"en-UK": "Please send text.",
		"es-ES": "Por favor, envia el texto.",
		"fa-IR": "لطفاً متن ارسال کنید.",
		"it-IT": "Si prega di inviare il testo.",
		"ru-RU": "Пожалуйста отправьте текст.",
	},
	MESSAGE_TEXT_CAN_YOU_RATE_AT_STOREBOT: {

		"de-DE": `🤖 Kannst du mich im Store Bot hoch bewerten und eine gute Bewertung schreiben?
		‎Es wird dich weniger als eine Minute kosten! 😇`,
		"en-UK": `🤖 Can you rate it high and write a good review in bots catalog Store Bot?
		‎It will take less than a minute of your time! 😇`,

		"es-ES": `🤖 Puedes valolarlo con una buena nota y una buena opinión en el catálogo Store Bot?
		‎Te costará menos de un minuto! 😇`,

		"fa-IR": `🤖  آیا می توانید در کاتالوگ روباتها در استور بوت امتیاز بالایی داده و اظهار نظر خوبی در مورد این روبات ثبت کنید؟
		این کار کمتر از یک دقیقه از وقت شما را می گیرد! 😇`,

		"it-IT": `🤖 Puoi votarlo in alto e scrivere una buona revisione nel catalogo Bot Store?
		Ci vorrà meno di un minuto del tuo tempo! 😇`,

		"ru-RU": `🤖 Можете поставить ему высокую оценку и хороший отзыв в каталоге ботов Store Bot?
		‎Это займет меньше минуты вашего времени! 😇`,
	},
	MESSAGE_TEXT_ASK_TO_WRITE_FEEDBACK_WITHIN_MESSENGER: {
		"de-DE": "‎Schreibe (auf Englisch oder Russisch) uns, was man am Bot besser machen kann:",
		"en-UK": "‎Share your thoughts (in English or Russian) about what could be done to make the bot better:",
		"es-ES": "‎Comparte tus pensamientos (en Inglés o Ruso) sobre qué podemos hacer para que el bot sea mejor:",
		"fa-IR": "نظرات خود را (به انگلیسی و روسی ) در مورد اینکه چه کاری می توان انجام داد تا این ربات بهتر شود، با ما به اشتراک بگذارید:",
		"it-IT": "Condividi i tuoi pensieri (in Inglese o Russo) su come sarebbe migliore secondo te il bot:",
		"ru-RU": "Поделитесь вашими мыслями (на русском или английском) о том, что нужно сделать, чтобы бот стал лучше:",
	},
	MESSAGE_TEXT_HOW_TO_RATE_AT_STOREBOT: {

		"de-DE": `<b>Wie man bewertet - in 3 einfachen Schritten:</b>

	1. Klick auf diesen Link, um eine Bewertung abzugeben:
	https://t.me/storebot?start={{bot}}

	‎2. Wähle "⭐️⭐️⭐️⭐️⭐️" 

	‎3. Schreib etwas Nettes auf Englisch oder wähle "Skip this step"

	Wirklich vielen Dank! Dank deiner Bewertung werden vielleicht mehr Leute auf diesen Bot aufmerksam. Das ist gut für die Motivation der Entwickler dieses Bots! 😎`,

		"en-UK": `<b>How to rate in 3 simple steps:</b>

	1. Click on this link to rate and review:
	https://t.me/storebot?start={{bot}}

	‎2. Click on the "⭐️⭐️⭐️⭐️⭐️" button

	‎3. Write your message or press "Skip this step" button

	Thank you very much! As a result of your actions, even more people will learn about the bot.All this will serve as the additional motivation for the developers! 😎`,

		"es-ES": `<b>Como valorar en 3 simples pasos:</b>

	1. Click este link para valorar y dejar tu opinión:
	https://t.me/storebot?start={{bot}}

	‎2. Click en "⭐️⭐️⭐️⭐️⭐️" botón

	‎3. Escribe tu mensage o apreta "Skip this step" botón

	¡Muchas gracias! Merced a tus acciones más gente conocerá a bot. Todo eso sirve para una motivación adicional a los creadores! 😎`,

		"it-IT": `<b>Come valutare in 3 semplici passaggi:</b>
	‎1. Clicca su questo link per votare e lasciare una recensione:
	‎https://t.me/storebot?start={{bot}}

	‎2. Clicca sul "⭐️⭐️⭐️⭐️⭐️" bottone

	‎3. Scrivi il tuo messaggio o premi "Salta questo step"

	Grazie infinitamente! Come risultato delle tue azioni, altre persone guarderanno il bot.Dando anche un motivo in più per continuare ai developers! 😎`,

		"fa-IR": `<b>چگونگی امتیازدهی در سه گام ساده :</b>

	1. برای امتیازدهی و ثبت نظرات بر روی لینگ زیر کلیک کنید
	https://t.me/storebot?start={{bot}}

	‎2. بر روی دکمه "⭐️⭐️⭐️⭐️⭐️" کلیک کنید

	‎3. پیام خودرا ثبت کنید یا روی دکمه "پرش از این مرحله" کلیک کنید

	بسیار سپاسگزاریم! عمل شما باعث می شود افراد بیشتری در مورد bot.All بیاموزند. این امر انگیزه مضاعفی به توسعه دهندگان این ربات می دهد ! 😎`,

		"ru-RU": `‎<b>Как поставить оценку в три простых шага:</b>

	1. Перейдите по ссылке, чтобы оставить оценку и отзыв:
	‎https://t.me/storebot?start={{bot}}

	‎2. Нажмите на кнопку "⭐️⭐️⭐️⭐️⭐️"

	‎3. Напишите сообщение или нажмите кнопку "Пропустить этот шаг"

	Спасибо вам большое! Благодаря этому о боте узнает больше людей — это служит дополнительной мотивацией для разработчиков! 😎`,
	},
	MESSAGE_TEXT_ASK_FOR_FEEDBAСK: {
		"de-DE": "Über ein kleines Feedback wie der Bot so ist, würden wir uns freuen. Es dauert nur ein paar Sekunden.",
		"en-UK": "We would appreciate if tell us how we doing. It takes just few seconds.",
		"es-ES": "Te agredecemos si valoras el funccionamiento de nuestro applicación. Te costará solo unos segundos.",
		"fa-IR": "سپاسگزار خواهیم بود اگر به ما بگویید کارمان چطور بوده است. این تنها چند ثانیه زمان میبرد.",
		"it-IT": "Ci farebbe piacere se lasciassi un voto per il nostro lavoro. Ti bastano solo alcuni secondi.",
		"ru-RU": "Будем признетельны если вы оцените работу нашего приложения. Это займёт всего несколько секунд.",
	},
	COMMAND_TEXT_GIVE_FEEDBACK: {
		"de-DE": "Bewerte diesen Bot",
		"en-UK": "Rate this bot",
		"es-ES": "Valora a bot",
		"fa-IR": "به این ربات امتیاز بدهید",
		"it-IT": "Vota questo bot",
		"ru-RU": "Оценить приложение",
	},
	COMMAND_TEXT_OPEN_STOREBOT_FOR_FEEDBACK: {
		"de-DE": "Mache eine Bewertung auf @Storebot",
		"en-UK": "Leave rating at @Storebot",
		"es-ES": "Valorar en @Storebot",
		"fa-IR": "امتیاز خود را اینجا وارد کنید @Storebot",
		"it-IT": "Lascia un voto a @Storebot",
		"ru-RU": "Оценить на  @Storebot",
	},
	MESSAGE_TEXT_ON_REFUSED_TO_RATE: {
		"ru-RU": `ОК, возможно вы смоежете поставить оценку в другой раз.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Так же будем признательны если вы предложите любые улучшения.
	`,
		/*------------------------------------------------------------*/
		"de-DE": `Okay, vielleicht werden wir wann anders bewertet.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Ansonsten freuen wir uns immer zu hören, was man besser machen kann.
	`,

		"en-UK": `OK, maybe you can rate us another time.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	We also will appreciate if you suggest any improvements.
	`,
		/*------------------------------------------------------------*/
		"es-ES": `OK, Quizás puedas valorar en otro momento.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	También te agredeceríamos si ofrecieras alguna mejora.
	`,
		/*------------------------------------------------------------*/

		// TODO(IT): Google translated
		"it-IT": `OK, forse ci puoi valutare un'altra volta.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Apprezzeremo anche se suggeriamo qualche miglioramento.
	`,
		/*------------------------------------------------------------*/
		"fa-IR": `بسیار خوب، ممکن است شما بتوانید زمان دیگری به ما امتیاز بدهید.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	همچنین سپاسگزار خواهیم بود اگر شما هرگونه امکان بهبود را با ما در میان بگذارید.
	`,
	},
	MESSAGE_TEXT_ON_FEEDBACK_POSITIVE: {
		"ru-RU": `Спасибо, мы очень старались!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Так же будем признательны если вы <a suggest-idea>предложите улучшения</a>.
	`,
		/*------------------------------------------------------------*/
		"de-DE": `Danke, wir arbeiten hart dran!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Wir freuen uns auch immer über <a suggest-idea>neue Ideen</a>.
	`,
		"en-UK": `Thanks, we worked hard!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	We also will appreciate if you <a suggest-idea>suggest improvements</a>.
	`,
		/*------------------------------------------------------------*/
		"es-ES": `Gracias, hemos trabajado duro!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Te agradeceríamos si <a suggest-idea>propusieras alguna mejora</a>.
	`,
		/*------------------------------------------------------------*/

		"it-IT": `GRAZIE MILLE, abbiamo lavoro duro!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Sarebbe ancora piu' apprezzatto se ci <a suggest-idea>suggerisci qualche miglioramento</a>.
	`,
		/*------------------------------------------------------------*/
		"fa-IR": `ممنونیم، ما سخت کارکرده ایم!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	ما همچنین سپاسگزار خواهیم بود اگر شما<a suggest-idea> هرگونه امکان بهبود را با ما در میان بگذارید </a>را.
	`,
	},
	MESSAGE_TEXT_YOU_CAN_HELP_BY: {
		"ru-RU": `
	Вы нам очень поможете если:

	* Поставите нам 5⭐ в <a storebot>каталоге ботов</a>.

	* Расскажите о боте своим друзьям.
	Например в <a share-vk>ВКонтакте</a>, <a share-fb>Facebook</a> или <a share-twitter>Twitter</a>.

* Поддержите дальнейшую разработку - <a href ="https://goo.gl/Qhh0yL">€2 через PayPal</a>
`,
		/*------------------------------------------------------------*/
		"de-DE": `
Es gibt viele Wege uns zu helfen:

* Gib uns 5⭐ im <a storebot>Verzeichnis der Bots</a>.

* Erzähl am besten all deinen Freunde davon.
Du könntest es auf <a share-fb>Facebook</a> posten oder auf <a share-twitter>Twitter</a> twittern.

* Ansonsten auch gerne eine kleine Spende - <a href = "https://goo.gl/Qhh0yL">€2 via PayPal</a> (<i>about $2.2</i>)
`,
		/*------------------------------------------------------------*/
		"en-UK": `
You can help us a lot if you:

* Give us 5⭐ at <a storebot>directory of bots</a>.

* Tell about the app to your friends.
For example at <a share-fb>Facebook</a> or <a share-twitter>Twitter</a>.

* Support further development - <a href = "https://goo.gl/Qhh0yL">€2 via PayPal</a> (<i>about $2.2</i>)
`,
		/*------------------------------------------------------------*/
		"es-ES": `
Nos ayudarías mucho si:

* Nos pusieras 5⭐ en <a storebot>directory of bots</a>.

* Hablaras del bot a tus amigos.
Por ejemplo <a share-fb>Facebook</a> o <a share-twitter>Twitter</a>.

* Apoyaras el siguiente trabajo - <a href = "https://goo.gl/Qhh0yL">€2 vía PayPal</a> (<i>about $2.2</i>)
`,
		/*------------------------------------------------------------*/

		// TODO(FA)
		"fa-IR": `:شما به ما کمک بسیاری می کنید اگر

* ثبت بازخورد مثبت در <a storebot>دایرکتوری روبات ها</a>.

* به دوستان خود در مورد برنامه اطلاع رسانی کنید.
برای مثال در <a share-fb>فیسبوک</a> یا <a share-twitter>توئیتر</a>.

* از توسعه بیشتر برنامه پشتیبانی کنید - <a href = "https://goo.gl/Qhh0yL">€2 از طریق پی پال</a> (<i>$2.2 حدود </i>)`,

		// TODO(IT)
		"it-IT": `
		Ci aiuteresti moltissimo se:

		  * Lasci un feedback positivo 5⭐ alla <a storebot>pagina del bot</a>.

		  * Raccontare dell'app ai tuoi amici.
		    Per esempio su <a share-fb>Facebook</a> o su <a share-twitter>Twitter</a>.

		  * Supporta ulteriormente lo sviluppo del bot - <a href="https://goo.gl/Qhh0yL">2€ tramite PayPal</a> (<i>circa $2.2</i>)
		`,
	},
	MESSAGE_TEXT_COUNTERPARTY_HAS_EMPTY_BALANCE: {
		"de-DE": `Du bist quitt mit %v`,
		"en-UK": `Balance is empty for %v`,
		"es-ES": `El balance es cero para %v`,
		"fa-IR": `تراز خالی است برای %v`,
		"it-IT": `Non hai alcun credito o debito con %v`,
		"ru-RU": `Нулевой баланс для %v`,
	},
	MESSAGE_TEXT_ASK_TO_TRANSLATE: {
		"ru-RU": `Хотите чтобы наш бот разговаривал на другом языке? Вы можете <a>помочь с переводом</a>.`,
		"de-DE": `Möchtest du den Bot in einer anderen Sprache? Du kannst beim <a>Übersetzen helfen</a>.`,
		"en-UK": `Do you want our bot to talk in other language? You can <a>help with translation</a>.`,
		"es-ES": `¿Te gustaría que nuestro bot hablara en otro idioma? Puedes <a>ayudar con traducción</a>.`,

		"fa-IR": `آیا می خواهید ربات ما به زبان دیگری صحبت کند؟ شما می توانید <a>با ترجمه به ما کمک کنید</a>.`,

		"it-IT": `Vuoi che il nostro bot parli altre lingue? Aiuta con la <a>traduzione</a>.`,
	},
	MESSAGE_TEXT_ON_FEEDBACK_NEUTRAL: {
		"ru-RU": `Чтож, мы очень старались. Ваша оценка будет передана разработчикам.

Может быть вы <a submit-bug>сообщите нам что не работает</a> или подскажите <a suggest-idea>как можно улучшить</a>?`,
		/*------------------------------------------------------------*/
		"de-DE": `Gut, wir geben uns Mühe. Deine Rückmeldung wird an die Entwickler weitergeleitet.

Vielleicht willst du <a submit-bug>einen Fehler melden</a> oder <a suggest-idea>eine Verbesserung vorschlagen</a>?`,
		/*------------------------------------------------------------*/
		"en-UK": `Well, we worked hard. You feedback will be passed to developers.

Maybe you can <a submit-bug>report your issue</a> or <a suggest-idea>suggest how we can improve</a>?`,
		/*------------------------------------------------------------*/
		"es-ES": `Bueno, hemos trabajado duro. Tu opinión se pasará a los creadores.

Quizás puedas <a submit-bug>informarnos de algún problema</a> o <a suggest-idea>proponernos qué podríamos mejorar</a>?`,
		/*------------------------------------------------------------*/

		"fa-IR": `خب، ما سخت کارکردیم. بازخورد شما به توسعه دهندگان برنامه منعکس می شود.

شما می توانید <a submit-bug>مشکل خود را گزارش دهید</a> یا <a suggest-idea>پیشنهاد دهید چطور میتوانیم بهبود ایجاد کنیم</a>?`,

		"it-IT": `Bene, il nostro lavoro non e' andato in vano. Il tuo feedback sara' inoltrato agli sviluppatori.

Per caso vuoi anche <a submit-bug>segnalare un problema</a> oppure <a suggest-idea>suggerire come possiamo migliorare</a>?`,
	},
	MESSAGE_TEXT_ON_FEEDBACK_NEGATIVE: {
		"de-DE": `Das tut uns sehr leid. Vielleicht willst du uns <a submit-bug>einen Fehler melden</a> oder <a suggest-idea>eine Verbesserung vorschlagen</a>?`,
		"en-UK": `We are very sorry. Maybe you can <a submit-bug>let us know what is wrong</a> or <a suggest-idea>suggest how we can improve</a>?`,
		/*------------------------------------------------------------*/
		"es-ES": `Lo sentimos mucho. Igual podrías <a submit-bug>decirnos qué no funcciona bien</a> o <a suggest-idea>proponernos cómo podemos mejorarlo</a>?`,
		/*------------------------------------------------------------*/

		"fa-IR": `ما بسیار متاسفیم. شما می توانید <a submit-bug>به ما بگویید مشکلتان چیست</a> یا <a suggest-idea>پیشنهاد دهید چطور میتوانیم بهبود ایجاد کنیم</a>?`,

		"it-IT": `Ci dispiace molto. Potresti farci sapere<a submit-bug>cosa non ti e' piaciuto</a> oppure <a suggest-idea>suggerirci come possiamo migliorare</a>?`,
		"ru-RU": `Нам очень стыдно. Может быть вы <a submit-bug>подскажите что не так</a> или <a suggest-idea>предложите усовершенствования</a>?`,
		/*------------------------------------------------------------*/
	},
	COMMAND_TEXT_ASK_FOR_FEEDBACK: {
		"de-DE": "Bitte bewerte unsere App",
		"en-UK": "Please rate our app",
		"es-ES": "Por favor valora nuestro app",
		"fa-IR": "لطفاً به برنامه ما امتیاز دهید",
		"it-IT": "Per favore vota il nostro bot",
		"ru-RU": "Оцените наше приложение?",
	},
	COMMAND_TEXT_FEEDBACK_POSITIVE: {
		"de-DE": "Ja, es ist eine tolle App!",
		"en-UK": "Yes, it's a great app!",
		"es-ES": "¡Sí, es una app fantástica!",
		"fa-IR": "بله، این برنامه عالی است",
		"it-IT": "Si, e' un app fantastica!",
		"ru-RU": "Да, отличное приложение!",
	},
	COMMAND_TEXT_FEEDBACK_NEUTRAL: {
		"de-DE": "Nicht schlecht, aber könnte besser sein",
		"en-UK": "Not bad but can be better.",
		"es-ES": "No está mal, pero podría ser mejor.",
		"fa-IR": "بد نیست ولی می تواند بهتر باشد.",
		"it-IT": "Non male ma potrebbe esser migliore.",
		"ru-RU": "Неплохо, но можно лучше.",
	},
	COMMAND_TEXT_FEEDBACK_NEGATIVE: {
		"de-DE": "Ich mag sie nicht",
		"en-UK": "Don't like it",
		"es-ES": "No me gusta",
		"fa-IR": "از این برنامه را نمی پسندم",
		"it-IT": "Non mi piace",
		"ru-RU": "Не нравится",
	},
	COMMAND_TEXT_FEEDBACK_NOT_READY: {
		"de-DE": "Ich bin mir noch unsicher",
		"en-UK": "Not decided yet",
		"es-ES": "Estoy aún indeciso",
		"fa-IR": "هنوز تصمیم نگرفته ام.",
		"it-IT": "Sono indeciso",
		"ru-RU": "Пока не понятно",
	},
	MESSAGE_TEXT_SETTINGS: {
		"de-DE": "Was willst du ändern?",
		"en-UK": "What do you want to adjust?",
		"es-ES": "¿Qué te gustaría modificar?",
		"fa-IR": "می خواهید چه چیزی را تنظیم کنید؟",
		"it-IT": "Cosa vuoi modificare?",
		"ru-RU": "Что будем настраивать?",
	},
	MT_NO_OUTSTANDING_TRANSFERS: {
		"de-DE": "", // TODO(DE)
		"en-UK": `You are trying to create return record but there are no outstanding debts.

If you believe this is a mistale please let us know in @DebtsTrackerGroup.`,
		"es-ES": "", // TODO(ES)
		"fa-IR": "", // TODO(FA)
		"it-IT": "", // TODO(IT)
		"ru-RU": `Вы пытаетесь создать запись о возврате долга, но мы не нашли не закрытых долгов.

Если вы считаете что это ошибка пожалуйста дайте нам знать в @DebtsTrackerGroup.`,
	},
	MT_ATTEMPT_TO_CREATE_DEBT_WITH_INTEREST_AFFECTING_OUTSTANDING: {
		"de-DE": "", // TODO(DE)
		"en-UK": "You are trying to create a debt with interest that will affect outstanding transfers. Please close them first.",
		"es-ES": "", // TODO(ES)
		"fa-IR": "", // TODO(FA)
		"it-IT": "", // TODO(IT)
		"ru-RU": "Вы пытаетесь создать запись о долге с процентами которая затронет незакрытые долги. Пожалуйста закройте их сначала.",
	},
	MESSAGE_TEXT_NOT_IMPLEMENTED_YET: {
		"de-DE": "Entschuldigung, diese Funktion ist noch nicht fertig programmiert.",
		"en-UK": "Sorry, this functionality is not implemented yet.",
		"es-ES": "Lo sentimos, esta función no está activa aún.",
		"fa-IR": "متاسفم، این عملکرد هنوز پیاده سازی نشده است.",
		"it-IT": "Spiacenti ma questa funzionalita' non e' ancora attiva.",
		"ru-RU": "Извините, данный функционал ещё не реализован.",
	},
	MESSAGE_TEXT_ASK_INVITE_CHANNEL: {
		"de-DE": "Worüber möchtest du eingeladen werden?",
		"en-UK": "How do you want to get an invite?",
		"es-ES": "¿Comó prefieres recibir la invitación?",
		"fa-IR": "می خواهید چگونه دعوت شوید؟",
		"it-IT": "Come vuoi ricevere l'invito?",
		"ru-RU": "Как вы хотите получить код приглашения?",
	},
	MESSAGE_TEXT_PLEASE_ENTER_INVITE_CODE: {
		"de-DE": "Bitte gib den Bestätigungs-Code ein:",
		"en-UK": "Please enter an invite code:",
		"es-ES": "Introduce el código de la invitación",
		"fa-IR": "لطفاً یک کد دعوت وارد کنید:",
		"it-IT": "Inserisci un codice invito:",
		"ru-RU": "Пожалуйста введите код приглашения:",
	},
	MESSAGE_TEXT_USER_EMAIL_FOR_INVITE_RECEIVED: {
		"de-DE": "Wir haben eine Nachricht an %v gesendet.\n\nBitte öffne die Nachricht und klick auf den Link, um deine Mail-Adresse zu bestätigen.",
		"en-UK": "We've sent a message to %v.\n\nPlease open the email and click a link to confirm your email address.",
		"es-ES": "Hemos enviado un mensage a %v.\n\nPor favor, abre tu e-mail y haz click en el link para confirmar tu e-mail.",
		"fa-IR": "ما یک پیام ارسال کردیم به %v.\n\nلطفاً ایمیل خود را باز کرده و روی لینک کلیک کنید تا آدرس ایمیل شما تایید شود.",
		"it-IT": "Abbiamo inviato un messaggio a %v.\n\nPer favore apri l'email e clicca sul link per confermare il tuo indirizzo email",
		"ru-RU": "Мы отправили письмо на %v.\n\nПожалуйста откройте его и кликните на ссылку для подтверждения адреса.",
	},
	MESSAGE_TEXT_USER_EMAIL_FOR_INVITE_SENT_TELEGRAM: {
		"de-DE": "Wenn Telegram öffnet, drücke auf <b>Start</b>.",
		"en-UK": "Once Telegram pop ups click the <b>Start</b> button.",
		"es-ES": "Después de abrir Telegram aprieta el <b>Start</b> botón.",
		"fa-IR": "وقتی تلگرام اجرا شد برروی دکمه  <b>شروع</b> کلیک کنید.",
		"it-IT": "Una volta aperto il bot su telegram clicca su <b>Start</b>.",
		"ru-RU": "После того как откроется Telegram нажмите кнопку <b>Start</b>.",
	},
	MESSAGE_TEXT_USER_CONTACT_FOR_INVITE_RECEIVED: {
		"de-DE": "Danke, du bist in der Warteschlange für eine Einladung.\n\nEs dauert etwa zwei bis drei Tage.\n\nAber du könntest den Code noch heute bekommen, wenn du einen Link auf Facebook teilst.",
		"en-UK": "Thanks, you've been queued for an invite.\n\nCurrent awaiting time is 2-3 days.\n\nYou can get an invite code today by sharing a link to the bot on Facebook.",
		"es-ES": "Gracias, ya estás inscrito en la cola para conseguir la invitación.\n\nTiempo de espera 2-3 días.\n\nPuedes conseguirlo hoy si compartes el link de nuestro bot en Facebook.",
		"fa-IR": "سپاسگزاریم، شما در نوبت دعوت قرار گرفتید\n\nزمان انتظار شما در حال حاضر 2-3 روز می باشد.\n\n شما می توانید با به اشتراک گذاری لینک روبات در فیسبوک امروز یک کد دعوت دریافت کنید. ",
		"it-IT": "Grazie, ora sei in coda per un codice invito.\n\nTempo di attesa medio 2-3 giorni.\n\nPuoi ottenere un codice invito subito condividendo il link al bot su Facebook.",
		"ru-RU": "Спасибо, вы записаны в очередь на получение приглашения.\n\nТекущее время ожидания 2-3 дня.\n\nВы можете получить приглашение сегодня если расскажите о нашем боте на Facebook.",
	},
	MESSAGE_TEXT_PLEASE_PROVIDE_YOUR_EMAIL: {
		"de-DE": "Bitte gib deine e-Mail Adresse an:",
		"en-UK": "Please provide your email address",
		"es-ES": "Por favor, esctibe tu e-mail",
		"fa-IR": "لطفاً آدرس ایمیل خود را وارد کنید.",
		"it-IT": "Inserisci il tuo indirizzo email:",
		"ru-RU": "Пожалуйста напишите ваш email адрес:",
	},
	MESSAGE_TEXT_PLEASE_PROVIDE_YOUR_PHONE_NUMBER: {
		"de-DE": "Bitte gib deine Telefonnummer an:",
		"en-UK": "Please provide your phone number",
		"es-ES": "Por favor, esctibe tu número de teléfono",
		"fa-IR": "لطفاً شماره تلفن خود را وارد نمایید.",
		"it-IT": "Inserisci il tuo numero di telefono:",
		"ru-RU": "Пожалуйста напишите номер вашего телефона:",
	},
	MESSAGE_TEXT_WRONG_INVITE_CODE: {
		"de-DE": "Ungültiger Code: %v",
		"en-UK": "Wrong invite code: %v",
		"es-ES": "El código de invitación no es correcto: %v",
		"fa-IR": "کد دعوت اشتباه است %v",
		"it-IT": "Codice invito: %v errato",
		"ru-RU": "Неправильный код приглашения: %v",
	},
	MESSAGE_TEXT_WRONG_EMAIL: {
		"de-DE": "Ungültige e-Mail Adresse.",
		"en-UK": "Wrong email address.",
		"es-ES": "El e-mail no es correcto.",
		"fa-IR": "آدرس ایمیل اشتباه است.",
		"it-IT": "L'email inserita e' sbagliata.",
		"ru-RU": "Неправильный email адрес.",
	},
	MESSAGE_TEXT_WRONG_PHONE_NUMBER: {
		"de-DE": "Ungültige Telefonnummer.",
		"en-UK": "Wrong phone number.",
		"es-ES": "El número de telefono no es correcto.",
		"fa-IR": "شماره تلفن اشتباه است",
		"it-IT": "Il numero inserito e' sbagliato.",
		"ru-RU": "Неправильный номер телефона.",
	},
	MESSAGE_TEXT_OK_PLEASE_TRY_AGAIN: {
		"de-DE": "Ok, bitte versuche es erneut.",
		"en-UK": "Ok, please try again.",
		"es-ES": "Ok, inténtalo de nuevo.",
		"fa-IR": "بسیار خوب، لطفا مجدداً سعی کنید.",
		"it-IT": "Ok, prova di nuovo.",
		"ru-RU": "Хорошо, попробуйте ещё раз.",
	},

	COMMAND_TEXT_MISTYPE_WILL_TRY_AGAIN: {
		"de-DE": "Ich habe mich vertippt.",
		"en-UK": "I've mistyped, will try again.",
		"es-ES": "Me he equivocado, lo intentaré otra vez",
		"fa-IR": "اشتباه تایپ کردم، دوباره امتحان می کنم.",
		"it-IT": "Ho sbagliato, riprovo.",
		"ru-RU": "Я опечатался, попробую ещё раз.",
	},
	COMMAND_TEXT_TELL_ME_MORE_ABOUT_INVITES: {
		"de-DE": "Erzähl mir mehr über diese Codes",
		"en-UK": "Tell me more about the codes",
		"es-ES": "Dime más de los códigos",
		"fa-IR": "در مورد کدها بیشتر برای من توضیح دهید.",
		"it-IT": "Ulteriori informazioni riguardo il codice invito.",
		"ru-RU": "Расскажите ка мне об этих кодах",
	},
	COMMAND_TEXT_INVITE_ME_BY_EMAIL: {
		"de-DE": "Sende mir einen Code per Mail",
		"en-UK": "Send me invite code by email",
		"es-ES": "Envíame el código de invitación por e-mail",
		"fa-IR": "کد دعوت را برای من از طریق ایمیل ارسال کنید.",
		"it-IT": "Inviami il codice invito tramite email",
		"ru-RU": "Хочу код приглашения на email",
	},
	COMMAND_TEXT_INVITE_ME_BY_SMS: {
		"de-DE": "Sende mir einen Code per SMS",
		"en-UK": "Send me invite code by SMS",
		"es-ES": "Envíame el código de invitación a través de SMS",
		"fa-IR": "کد دعوت را برای من از طریق پیام کوتاه ارسال کنید.",
		"it-IT": "Inviami il codice invito tramite SMS",
		"ru-RU": "Хочу код приглашения по SMS",
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_EMAIL: {
		"de-DE": "Sende mir nochmal einen Code per Mail",
		"en-UK": "Send me new invite code by email",
		"es-ES": "Envíame un nuevo código de invitación por e-mail",
		"fa-IR": "یک کد دعوت جدیداز طریق ایمیل برای من  ارسال کنید.",
		"it-IT": "Inviami il nuovo codice invito tramite email",
		"ru-RU": "Новый код приглашения на email",
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_SMS: {
		"de-DE": "Sende mir nochmal einen Code per SMS",
		"en-UK": "Send me new invite code by SMS",
		"es-ES": "Envíame un nuevo código de invitación a través de SMS",
		"fa-IR": "یک کد دعوت جدید از طریق پیام کوتاه برای من ارسال کنید.",
		"it-IT": "Inviami il nuovo codice invito tramite SMS",
		"ru-RU": "Новый код приглашения по SMS",
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_TELEGRAM: {
		"de-DE": "Sende mir nochmal einen Code per Telegram",
		"en-UK": "Send me new invite by Telegram",
		"es-ES": "Envíame un nuevo código de invitación a través de Telegram",
		"fa-IR": "یک کد دعوت جدید از طریق تلگرام برای من ارسال کنید.",
		"it-IT": "Inviami il nuovo codice invito tramite Telegram",
		"ru-RU": "Получить приграшение в Telegram",
	},
	MESSAGE_TEXT_UNKNOWN_LANGUAGE: {
		"de-DE": "Unbekannte Sprache. Bitte wähle eine aus der Liste:",
		"en-UK": "Unknown language. Please choose 1 from the options:",
		"es-ES": "Idioma desconocido. Por favor elige 1 de las opciones:",
		"fa-IR": "زبان ناشناخته. لطفاً یکی از گزینه ها را انتخاب کنید:",
		"it-IT": "Lingua socnosciuta. Per favore scegline una tra le opzioni:",
		"ru-RU": "Незнакомый язык. Пожалуйста выберете один из предоставленных:",
	},
	MESSAGE_TEXT_UNKNOWN_COUNTERPARTY: {
		"de-DE": "Unbekannte Gegnerpartei. Bitte wähle aus der Liste oder wähle <b>neuer Kontakt</b>.",
		"en-UK": "Unknown counterparty. Please choose <b>Add new</b> if it's a new contact.",
		"es-ES": "Contacto desconocido. Por favor elige <b>Añadir</b> si es un contacto nuevo.",
		"fa-IR": "مخاطب تراکنش شناسایی نشد. <b>یک مورد جدید اضافه کنید</b> اگر این ایشان یک مخاطب جدید هستند.",
		"it-IT": "Destinatario sconosciuto. Scegli <b>Aggiugni nuovo</b> se e' un nuovo contatto.",
		"ru-RU": "Неизвестный контакт. Пожалуйста выберите <b>Добавить</b> если это новый контакт.",
	},
	MESSAGE_TEXT_UNKNOWN_COUNTERPARTY_ON_RETURN: {
		"de-DE": "Unbekannte Gegnerpartei. Bitte wähle aus der Liste.",
		"en-UK": "Unknown counterparty. Please choose from the list.",
		"es-ES": "Contacto desconocido. Por favor elige de la lista.",
		"fa-IR": "مخاطب تراکنش شناسایی نشد. لطفاً از فهرست انتخاب کنید.",
		"it-IT": "Destinatario sconosciuto. Scegli dalla lista qui sotto.",
		"ru-RU": "Неизвестный контакт. Пожалуйста выберите из списка.",
	},
	MESSAGE_TEXT_UNKNOWN_DEBT: {
		"de-DE": "Unbekannter Schuldschein. Bitte wähle aus der Liste.",
		"en-UK": "Unknown debt. Please choose from the list.",
		"es-ES": "Deuda desconocida. Por favor elige de la lista.",
		"fa-IR": "بدهی ناشناخته است. لطفا از فهرست انتخاب کنید.",
		"it-IT": "Debito sconosciuto. Scegli dalla lista qui sotto.",
		"ru-RU": "Неизвестный долг. Пожалуйста выберите из списка.",
	},
	MESSAGE_TEXT_BILL_CARD_HEADER: {
		"de-DE": `<b>Rechnung</b>: <code>%v</code> — %v`,
		"en-UK": `<b>Bill</b>: <code>%v</code> — %v`,
		"es-ES": `<b>Cuenta</b>: <code>%v</code> — %v`,
		"fa-IR": "<b>Bill</b>: <code>%v</code> — %v", // TODO(FA)
		"it-IT": "<b>Bill</b>: <code>%v</code> — %v", // TODO(IT)
		"ru-RU": `<b>Cчёт</b>: <code>%v</code> — %v`,
	},
	MESSAGE_TEXT_MEMBERS_TITLE: {
		"de-DE": "Mitglieder",
		"en-UK": "Members",
		"es-ES": "Miembros",
		"fa-IR": "اعضا",   // TODO(FA) verify
		"it-IT": "Membri", // TODO(IT)
		"ru-RU": "Участники",
	},
	MT_TEXT_MEMBERS_COUNT: {
		"en-UK": "<b>Members</b>: %d",
		"es-ES": "<b>Miembros</b>: %d",
		"fa-IR": "اعضا %d",           // TODO(FA) verify - missin <b></b>
		"it-IT": "<b>Membri</b>: %d", // TODO(IT)
		"ru-RU": "<b>Участников</b>: %d",
	},
	ALERT_TEXT_NOTHING_CHANGED: {
		"de-DE": "Nichts geändert",
		"en-UK": "Nothing changed",
		"es-ES": "Nada ha cambiado",     // TODO(ES)
		"fa-IR": "چیزی تغییر نکرده است", // TODO(FA) verify
		"it-IT": "Niente è cambiato",    // TODO(IT) verify
		"ru-RU": "Ничего не изменилось",
	},
	ALERT_TEXT_YOU_ARE_ALREADY_MEMBER_OF_THE_GROUP: {
		"de-DE": "Du bist schon Mitglied beim Teilen dieser Rechnung.",
		"en-UK": "You are already a member of this bill splitting group.",
		"es-ES": "Ya es miembro de este grupo de división de facturas.",  // TODO(ES)
		"fa-IR": "شما قبلا عضو این گروه تقسیم لایحه هستید.",              // TODO(FA)
		"it-IT": "Sei già membro di questo gruppo di divisione fatture.", // TODO(IT)
		"ru-RU": "Вы уже участник этой группы по совместной оплате счетов.",
	},
	MESSAGE_TEXT_YOUR_BILL_SPLITTING_GROUPS: {
		"de-DE": "Gruppen, mit denen du Rechnungen teilst",
		"en-UK": "Your bills splitting groups",
		"es-ES": "Ya es miembro de este grupo de división de facturas.",  // TODO(ES) verify
		"fa-IR": "شما قبلا عضو این گروه تقسیم لایحه هستید.",              // TODO(FA) verify
		"it-IT": "Sei già membro di questo gruppo di divisione fatture.", // TODO(IT) verify
		"ru-RU": "Ваши группы совметсной оплаты",
	},
	MESSAGE_TEXT_USE_ARROWS_TO_SELECT_GROUP: {
		"en-UK": "Use ⬅️ & ➡️ to select group",
		"es-ES": "Usa ⬅️ y ➡️ para seleccionar el grupo",    // TODO(ES) verify
		"fa-IR": "برای انتخاب گروه از ⬅️ & ❢️ استفاده کنید", // TODO(FA) verify
		"it-IT": "Usare ⬅️ & ➡️ per selezionare il gruppo",  // TODO(IT) verify
		"ru-RU": "Используйте ⬅️ и ➡️ чтобы выбрать группу.",
	},
	MESSAGE_TEXT_NO_GROUPS: {
		"de-DE": "Du gehörst zu keiner Gruppe, die sich Rechnungen teilt.",
		"en-UK": "You are not a participant of any bill splitting group.",
		"es-ES": "Usted no es participante de ningún grupo de división de facturas.", // TODO(ES) verify
		"fa-IR": "شما شرکت کننده در هر گروه تقسیم لایحه نیستید.",                     // TODO(FA) verify
		"it-IT": "Non sei un partecipante a qualsiasi gruppo di divisione fatture.",  // TODO(IT) verify
		"ru-RU": "Вы не состоите в группах совместной оплаты.",
	},
	MESSAGE_TEXT_USER_JOINED_GROUP: {
		"de-DE": `Hi %v, du bist der Gruppe, die sich Rechnungen teilt, beigetreten.`,
		"en-UK": `Hi %v, you joined this bill splitting group.`,
		"fa-IR": "سلام %v، شما به گروه تقسیم این لایحه پیوستید",              // TODO(FA) verify
		"it-IT": "Hi %v, sei entrato in questo gruppo di divisione fatture.", // TODO(IT) verify
		"ru-RU": `Привет %v, вы присоеденились к этой группе по совместной оплате счетов.
		`,
	},
	MESSAGE_TEXT_MEMBERS_CARD_TITLE: {
		"de-DE": "<b>Wer sich die Rechnung teilt</b> (%d)",
		"en-UK": "<b>Bills splitting members</b> (%d)",
		"fa-IR": "(%d) <b>نقض تقسیم اعضا</b>",                     // TODO(FA) verify
		"it-IT": "<b>Membri di divisione delle bollette</b> (%d)", // TODO(IT) verify
		"ru-RU": "<b>Участники совместных оплат</b> (%d)",
	},
	MESSAGE_TEXT_SPLIT_MODE: {
		"de-DE": "<b>Teilen</b>: %v",
		"en-UK": "<b>Split</b>: %v",
		"es-ES": "<b>División</b>: %v", // TODO(ES) verify
		"fa-IR": "<b>شکاف</b>: %v",     // TODO(FA) verify
		"it-IT": "<b>Diviso</b>: %v",   // TODO(IT) verify
		"ru-RU": "<b>Делить</b>: %v",
	},
	MESSAGE_TEXT_ASK_HOW_TO_SPLIT_IN_GROP: {
		"de-DE": "In welchem Verhältnis teilt ihr in dieser Gruppe eure Rechnungen?",
		"en-UK": "In what proportion do you split bills in this group?",
		"es-ES": "¿En qué proporción divide las facturas en este grupo?",     // TODO(ES) verify
		"fa-IR": "در این سهم، آیا شما در این گروه حساب ها را تقسیم می کنید؟", // TODO(FA) verify
		"it-IT": "In quale percentuale dividi le fatture in questo gruppo?",  // TODO(IT) verify
		"ru-RU": "В какой пропорции вы делите счета в этой группе?",
	},
	MESSAGE_TEXT_MEMBERS_CARD_FOOTER: {
		"de-DE": "Klick <code>Join</code>, um auch Rechnungen zu teilen.",
		"en-UK": "Click <code>Join</code> to participate in bills splitting.",
		"es-ES": "¿En qué proporción divide las facturas en este grupo?",     // TODO(ES) verify
		"fa-IR": "در این سهم، آیا شما در این گروه حساب ها را تقسیم می کنید؟", // TODO(FA) verify
		"it-IT": "In quale percentuale dividi le fatture in questo gruppo?",  // TODO(IT) verify
		"ru-RU": "Жмите <code>Присоедениться</code> чтобы учавствовать.",
	},
	MESSAGE_TEXT_BILL_CARD_MEMBER_TITLE: {
		"en-UK": "{{.N}}. {{.MemberName}}",
		"es-ES": "{{.N}}. {{.MemberName}}",
		"fa-IR": "{{.N}}. {{.MemberName}}", // TODO(FA) verify
		"it-IT": "{{.N}}. {{.MemberName}}",
		"ru-RU": "{{.N}}. {{.MemberName}}",
	},
	MESSAGE_TEXT_BILL_CARD_MEMBERS_ROW: {
		"de-DE": "<i>{{.Percent}}%</i>",
		"en-UK": "<i>{{.Percent}}%</i>",
		"es-ES": "<i>{{.Percent}}%</i>",
		"fa-IR": "<i>{{.Percent}}%</i>", // TODO(FA) verify
		"it-IT": "<i>{{.Percent}}%</i>",
		"ru-RU": "<i>{{.Percent}}%</i>",
	},
	MESSAGE_TEXT_BILL_CARD_MEMBERS_ROW_OWES: {
		"de-DE": "\n   <i>schuldet {{.Owes}}</i>",
		"en-UK": "\n   <i>owes {{.Owes}}</i>",
		"es-ES": "\n   <i>debo {{.Owes}}</i>",
		"ru-RU": "\n   <i>должен {{.Owes}}</i>",
		"fa-IR": "\n   <i>بدهکار است {{.Owes}}</i>", // TODO(FA) verify
	},
	MESSAGE_TEXT_BILL_CARD_MEMBERS_ROW_PAID: {
		"de-DE": "\n   <i>bezahlte {{.Paid}}</i>",
		"en-UK": "\n   <i>paid {{.Paid}}</i>",
		"es-ES": "\n   <i>he pagado {{.Paid}}</i>",
		"ru-RU": "\n   <i>заплатил {{.Paid}}</i>",
		"fa-IR": "\n   <i>پرداخت شده {{.Paid}}</i>", // TODO(FA) verify
	},
	MESSAGE_TEXT_BILL_CARD_MEMBERS_ROW_PART_PAID: {
		"de-DE": "\n<i>bezahlte {{.Paid}}, schuldet noch {{.Owes}}</i>",
		"en-UK": "\n<i>paid {{.Paid}}, owes {{.Owes}}</i>",
		"es-ES": "\n<i>he pagado {{.Paid}}, debo {{.Owes}}</i>",
		"ru-RU": "\n<i>заплатил {{.Paid}}, должен {{.Owes}}</i>",
		"fa-IR": "\n<i>پرداخت شده {{.Paid}}, بدهکار است {{.Owes}}</i>", // TODO(FA) verify
	},
	MESSAGE_TEXT_BILL_ASK_WHO_PAID: {
		"de-DE": "Bitte wähle, wer die Rechnung gezahlt hat:",
		"en-UK": "Please choose who paid for the bill:",
		"es-ES": "Por favor, elige quien ha pagado la cuenta:",
		"fa-IR": "لطفا انتخاب کنید که چه کسانی برای این لایحه پرداخت کرده اند:", // TODO(FA) verify
		"it-IT": "Scegliere chi ha pagato la fattura:",                          // TODO(IT) verify
		"ru-RU": "Пожалуйста выберите кто заплатил по счёту:",
	},
	MESSAGE_TEXT_STATUS: {
		"de-DE": "Status: %v",
		"en-UK": "Status: %v",
		"es-ES": "Estado: %v",
		"fa-IR": "وضعیت: %v", // TODO(FA) verify
		"it-IT": "Stato: %v", // TODO(IT) verify
		"ru-RU": "Статус: %v",
	},
	BUTTON_TEXT_ADD_MEMBER: {
		"de-DE": "Partei hinzufügen",
		"en-UK": "Add participant",
		"es-ES": "Añadir participante",
		"fa-IR": "افزودن مشارکت کننده",   // TODO(FA) verify
		"it-IT": "Aggiungi partecipante", // TODO(IT) verify
		"ru-RU": "Добавить участника",
	},
	BUTTON_TEXT_FINALIZE_BILL: {
		"de-DE": "🔓 Rechnung abschließen",
		"en-UK": "🔓 Lock the bill",
		"es-ES": "🔓 Cerrar la cuenta",
		"fa-IR": "🔓 لایحه را قفل کنید", // TODO(FA) verify
		"it-IT": "🔓 Bloccare il conto", // TODO(IT) verify
		"ru-RU": "🔓 Закрыть счёт",
	},
	BUTTON_TEXT_EDIT_BILL: {
		"de-DE": "✏️ Bearbeiten",
		"en-UK": "✏️ Edit",
		"es-ES": "✏️ Editar",
		"fa-IR": "✏️ ویرایش",   // TODO(FA) verify
		"it-IT": "✏️ Modifica", // TODO(IT) verify
		"ru-RU": "✏️ Изменить",
	},
	BUTTON_TEXT_SPLIT_MODE: {
		"de-DE": "➗ Teilen: %v",
		"en-UK": "➗ Split: %v",
		"es-ES": "➗ Dividir: %v",
		"fa-IR": "➗ تقسیم: %v", // TODO(FA) verify
		"it-IT": "➗ Split: %v", // TODO(IT) verify
		"ru-RU": "➗ Делить: %v",
	},
	MESSAGE_TEXT_SPLIT_LABEL_WITH_VALUE: {
		"de-DE": "Teilen: %v",
		"en-UK": "Split: %v",
		"es-ES": "Dividir: %v",
		"fa-IR": "تقسیم: %v", // TODO(FA) verify
		"it-IT": "Split: %v", // TODO(IT) verify
		"ru-RU": "Делить: %v",
	},
	STATUS_DRAFT: {
		"de-DE": "Entwurf",
		"en-UK": "draft",
		"es-ES": "borrador",
		"fa-IR": "پیش نویس", // TODO(FA) verify
		"it-IT": "bozza",    // TODO(IT) verify
		"ru-RU": "черновик",
	},
	SPLIT_MODE_EQUALLY: {
		"de-DE": "Gleichverteilt",
		"en-UK": "Equally",
		"es-ES": "A partes iguales",
		"ru-RU": "Поровну",
		"fa-IR": "به همان اندازه", // TODO(FA) verify
		"it-IT": "Ugualmente",     // TODO(IT) verify
	},
	SPLIT_MODE_PERCENTAGE: {
		"de-DE": "Prozentual",
		"en-UK": "Percentage",
		"es-ES": "Porcentaje",
		"fa-IR": "درصد",        // TODO(FA)
		"it-IT": "Percentuale", // TODO(IT) verify
		"ru-RU": "В процентах",
	},
	SPLIT_MODE_EXACT_AMOUNT: {
		"de-DE": "Exakte Summen",
		"en-UK": "Exact amounts",
		"es-ES": "Importes exactos",
		"fa-IR": "مقادیر دقیق",     // TODO(FA) verify
		"it-IT": "Quantità esatte", // TODO(IT) verify
		"ru-RU": "Точные суммы",
	},
	SPLIT_MODE_SHARES: {
		"de-DE": "Teilen",
		"en-UK": "Shares",
		"es-ES": "Proporciones",
		"fa-IR": "سهام",   // TODO(FA) verify
		"it-IT": "Azioni", // TODO(IT) verify
		"ru-RU": "В долях",
	},
	BUTTON_TEXT_JOIN: {
		"de-DE": "➕ Beitreten",
		"en-UK": "➕ Join",
		"es-ES": "➕ Adherirse",
		"fa-IR": "➕ عضویت", // TODO(FA) verify
		"it-IT": "➕ Join",  // TODO(IT) verify
		"ru-RU": "➕ Присоедениться",
	},
	BUTTON_TEXT_LEAVE: {
		"de-DE": "Verlassen",
		"en-UK": "Leave",
		"es-ES": "Salir",    // TODO(ES) verify
		"fa-IR": "ترک کردن", // TODO(FA) verify
		"it-IT": "Partire",  // TODO(IT) verify
		"ru-RU": "Покинуть",
	},
	BUTTON_TEXT_DUE: {
		"de-DE": "📅 Fällig: %v",
		"en-UK": "📅 Due: %v",
		"es-ES": "📅 Hasta: %v",
		"fa-IR": "📅 مورد: %v",   // TODO(FA) verify
		"it-IT": "📅 Dovuto: %v", // TODO(IT) verify
		"ru-RU": "📅 До: %v",
	},
	NOT_SET: {
		"de-DE": "nicht gesetzt",
		"en-UK": "not set",
		"es-ES": "no establecido",
		"fa-IR": "تنظیم نشده",    // TODO(FA) verify
		"it-IT": "non impostato", // TODO(IT) verify
		"ru-RU": "не задано",
	},
	BUTTON_TEXT_MANAGE_MEMBERS: {
		"de-DE": "👫 Parteien",
		"en-UK": "👫 Participants",
		"es-ES": "👫 Participantes",
		"fa-IR": "👫 شرکت کنندگان", // TODO(FA) verify
		"it-IT": "👫 Partecipanti", // TODO(IT) verify
		"ru-RU": "👫 Участники",
	},
	BUTTON_TEXT_CHANGE_BILL_PAYER: {
		"de-DE": "🔀 Bezahler ändern",
		"en-UK": "🔀 Change payer",
		"es-ES": "🔀 Cambiar el pagador",
		"fa-IR": "🔀 تغییر پرداخت کننده", // TODO(FA) verify
		"it-IT": "🔀 Cambia il pagatore", // TODO(IT) verify
		"ru-RU": "🔀 Сменить плательщика",
	},
	COMMAND_TEXT_I_PAID: {
		"de-DE": "Ich habe bezahlt",
		"en-UK": "I paid",
		"es-ES": "he pagado",
		"fa-IR": "پرداخت کردم", // TODO(FA) verify
		"it-IT": "Ho pagato",   // TODO(IT) verify
		"ru-RU": "Я заплатил",
	},
	COMMAND_TEXT_I_OWE: {
		"de-DE": "Ich schulde",
		"en-UK": "I owe",
		"es-ES": "Yo debo",
		"fa-IR": "من بدهکارم",     // TODO(FA) verify
		"it-IT": "Sono in debito", // TODO(IT) verify
		"ru-RU": "Я должен",
	},
	COMMAND_TEXT_OWED_TO_ME: {
		"de-DE": "schuldet mir",
		"en-UK": "Owed to me",
		"es-ES": "Me deben",
		"fa-IR": "به من تعلق دارد", // TODO(FA) verify
		"it-IT": "È dovuto a me",   // TODO(IT) verify
		"ru-RU": "Должны мне",
	},
	MESSAGE_TEXT_BILL_HEADER: {
		"de-DE": "Rechnung: %v",
		"en-UK": "Bill: %v",
		"es-ES": "Cuenta: %v",
		"fa-IR": "بیل :%v",  // TODO(FA) verify
		"it-IT": "Bill: %v", // TODO(IT) verify
		"ru-RU": "Cчёт: %v",
	},
	MESSAGE_TEXT_NEW_DEBT_HEADER: {
		"de-DE": "Rechnung: %v",
		"en-UK": "Bill: %v",
		"es-ES": "Cuenta: %v",
		"fa-IR": "بیل: %v",  // TODO(FA) verify
		"it-IT": "Bill: %v", // TODO(IT) verify
		"ru-RU": "Cчёт: %v",
	},
	MESSAGE_TEXT_GROUPS_ONLY_COMMAND: {
		"de-DE": "",
		"en-UK": "This command is available in group chats only for now.",
		"es-ES": "",
		"fa-IR": "", // TODO(FA)
		"it-IT": "", // TODO(IT)
		"ru-RU": "Эта команда пока что доступна только в групповых чатах",
	},
	MESSAGE_TEXT_ALREADY_HAS_CONTACT_WITH_SUCH_NAME: {
		"de-DE": "", // TODO(DE)
		"en-UK": "You already have contact with name: %v",
		"es-ES": "", // TODO(ES)
		"fa-IR": "", // TODO(FA)
		"it-IT": "", // TODO(IT)
		"ru-RU": "У вас уже есть контакт с таким именем: %v",
	},
	MESSAGE_TEXT_ALREADY_BILL_MEMBER: {
		"de-DE": "%v, du teilst diese Rechnung bereits.",
		"en-UK": "%v, you are sharing this bill already.",
		"es-ES": "%v, estás compartiendo esta cuenta ya.",
		"fa-IR": "%v، شما قبلا این لایحه را به اشتراک می گذارید",      // TODO(FA) verify
		"it-IT": "%v, stai già condividendo questo disegno di legge.", // TODO(IT) verify
		"ru-RU": "%v, вы уже входите в состав участников.",
	},
	MESSAGE_TEXT_USER_JOINED_BILL: {
		"de-DE": "%v ist dem Teilen der Rechnung beigetreten.",
		"en-UK": "%v joined to bill sharing.",
		"es-ES": "%v pagar conjuntamente.",
		"fa-IR": "%v به اشتراک گذاری لایحه پیوست.",        // TODO(FA) verify
		"it-IT": "%v unito alla condivisione di fatture.", // TODO(IT) verify
		"ru-RU": "%v присоеденился к совместной оплате.",
	},
	BUTTON_TEXT_I_PAID_FOR_THE_BILL: {
		"de-DE": "Die Rechnung wurde von mir bezahlt.",
		"en-UK": "The bill was paid by me.",
		"es-ES": "La factura fue pagada por mí.",  // TODO(ES) verify
		"fa-IR": "این لایحه توسط من پرداخت شد",    // TODO(FA) verify
		"it-IT": "Il conto è stato pagato da me.", // TODO(IT) verify
		"ru-RU": "Этот счёт оплатил(а) я.",
	},
	BUTTON_TEXT_I_OWE_FOR_THE_BILL: {
		"de-DE": "Ich muss noch was dabeigeben",
		"en-UK": "I owe for this bill",
		"es-ES": "Debo esta factura",                // TODO(ES) verify
		"fa-IR": "من برای این لایحه بدهکار هستم",    // TODO(FA) verify
		"it-IT": "Devo per questo disegno di legge", // TODO(IT) verify
		"ru-RU": "Я должен по этому счёту",
	},
	BUTTON_TEXT_I_DO_NOT_SHARE_THIS_BILL: {
		"de-DE": "Ich habe damit nichts zutun",
		"en-UK": "I don't share this bill",
		"es-ES": "No comparto esta cuenta",               // TODO(ES) verify
		"fa-IR": "من این لایحه را به اشتراک نمی گذارم",   // TODO(FA) verify
		"it-IT": "Non condivido questo disegno di legge", // TODO(IT) verify
		"ru-RU": "Я не учавствую в этой покупке",
	},
	MESSAGE_TEXT_YOU_JOINED_BILL: {
		"de-DE": "Du bist dem Teilen der Rechnung beigetreten.",
		"en-UK": "You've joined to bill sharing.",
		"es-ES": "Te has agregado para pagar conjuntamente .",
		"fa-IR": "شما به اشتراک گذاشتن لایحه پیوستید",          // TODO(FA) verify
		"it-IT": "Sei entrato a far parte della fatturazione.", // TODO(IT) verify
		"ru-RU": "Вы присоеденились к совместной оплате.",
	},
	ARTICLE_TITLE_SPLIT_BILL: {
		"de-DE": "eine Rechnung teilen",
		"en-UK": "Split bill/purchase",
		"es-ES": "Compartir la cuenta/compra",
		"fa-IR": "لایحه / خرید تقسیم شده",    // TODO(FA) verify
		"it-IT": "Bolletta Split / acquisto", // TODO(IT) verify
		"ru-RU": "Разделить счёт/покупку",
	},
	ARTICLE_SUBTITLE_SPLIT_BILL: {
		"de-DE": "Wert: %v\nTeile deine Kosten mit Freunden und verfolge deren Rückzahlungen.",
		"en-UK": "Amount: %v\nShares expenses with friends & track paybacks",
		"es-ES": "Importe: %v\nCompartir los gastos entre amigos y seguir las devoluciones",  // TODO(ES): Have to be shorter
		"fa-IR": "مقدار: %v" + "\n" + "هزینه ها را با دوستان و بازپرداخت پیگیری می کند",      // TODO(FA) verify
		"it-IT": "Importo: %v\nDisponi i costi con gli amici e le retribuzioni delle tracce", // TODO(IT) verify
		"ru-RU": "Сумма: %v\nПоделить траты между друзьями и отследить возвраты",
	},

	ARTICLE_NEW_DEBT_TITLE: {
		"de-DE": "Neuer Schuldschein",
		"en-UK": "New debt",
		"es-ES": "Nueva deuda",
		"fa-IR": "بدهی جدید",    // TODO(FA): Verify
		"it-IT": "Nuovo debito", // TODO(IT): Verify
		"ru-RU": "Новый долг",
	},
	ARTICLE_NEW_DEBT_SUBTITLE: {
		"de-DE": "Wert: %v\nZur Fälligkeit wird eine Benachrichtigung geschickt, falls so eingestellt",
		"en-UK": "Amount: %v\nSends notifications on due date if set",
		"es-ES": "Importe: %v\nEnviar las notificaciones el día de vencimiento",
		"fa-IR": "مقدار: %v" + "\n" + "اگر تنظیم شود، اطلاعیه ها را در تاریخ تعیین شده ارسال می کند", // TODO(FA):  verify
		"it-IT": "Importo: %v\nSend le notifiche alla data di scadenza se impostato",                 // TODO(IT) verify
		"ru-RU": "Сумма: %v\nЗапись долга и рассылка оповещений в день возврата.",
	},
	SPLITUS_PLEASE_JOIN_IF_NOT_ON_THE_LIST: {
		"de-DE": `Bitte tritt zuerst bei, falls dein Name nicht auf der Liste ist.`,
		"en-UK": `Please join if your name is not on the list above.`,
		"fa-IR": `اگر نام شما در لیست بالا نیست، لطفا پیوست شوید.`,                   // TODO(FA) verify
		"it-IT": `Si prega di unirti se il tuo nome non è nell'elenco di cui sopra.`, // TODO(IT) verify
		"ru-RU": `Пожалуйста присоеденяйтесь если ваше не в списке.`,                 // TODO(RU)
	},
	SPLITUS_TEXT_HI_IN_GROUP: {
		"de-DE": `Ich bin <b>Splitus</b>. Danke fürs hinzufügen!`,
		"en-UK": `I'm <b>Splitus</b>. Thanks for adding me!`,
		"es-ES": `Soy <b>Splitus</b>. ¡Gracias por agregarme!`,      // TODO(ES) verify
		"fa-IR": `من <b>Splitus</b> با تشکر برای اضافه کردن من!`,    // TODO(FA) verify
		"it-IT": `Sono <b>Splitus</b>. Grazie per averci aggiunto!`, // TODO(IT) verify
		"ru-RU": `Меня зовут <b>Сплитус</b>. Спасибо что добавили!`,
	},
	COLLECTUS_TEXT_HI_IN_GROUP: {
		"en-UK": `I'm <b>Collectus.</b> Thanks for adding me!`,
		"es-ES": `Soy <b>Collectus.</b> ¡Gracias por agregarme!`,      // TODO(ES) verify
		"fa-IR": `من <b>Collectus</b> با تشکر برای اضافه کردن من!`,    // TODO(FA) verify
		"it-IT": `Sono <b>Collectus.</b> Grazie per averci aggiunto!`, // TODO(IT) verify
		"ru-RU": `Меня зовут <b>Коллектус.</b> Спасибо что добавили!`,
	},
	MT_GROUP_LABEL: {
		"en-UK": `<b>Group</b>: %v`,
		"es-ES": `<b>Group</b>: %v`, // TODO(ES)
		"fa-IR": `<b>Group</b>: %v`, // TODO(FA)
		"it-IT": `<b>Group</b>: %v`, // TODO(IT)
		"ru-RU": `<b>Группа</b>: %v`,
	},
	MT_SPONSORS_HEADER: {
		"en-UK": `<b>Sponsors</b>:`,
		"es-ES": `<b>Patrocinadores</b>:`, // TODO(ES)
		"fa-IR": `<b>حامیان</b>:`,         // TODO(FA) verify
		"it-IT": `<b>Sponsors</b>:`,       // TODO(IT)
		"ru-RU": `<b>Спонсоры</b>:`,
	},
	MT_DEBTORS_HEADER: {
		"en-UK": `<b>Debtors</b>:`,
		"es-ES": `<b>Deudores</b>:`, // TODO(ES) verify
		"fa-IR": `<b>Debtors</b>:`,  // TODO(FA) verify
		"it-IT": `<b>بدهکار</b>:`,   // TODO(IT) verify
		"ru-RU": `<b>Должники</b>:`,
	},
	BT_DEFAULT_CURRENCY: {
		"en-UK": `Currency: %v`,
		"es-ES": `Moneda: %v`,   // TODO(ES) verify
		"fa-IR": `واحد پول: %v`, // TODO(FA) verify
		"it-IT": `Moneta: %v`,   // TODO(IT) verify
		"ru-RU": `Валюта: %v`,
	},
	MESSAGE_TEXT_ASK_LANG: {
		"de-DE": `Welche Sprache wird hier gesprochen?`,
		"en-UK": `What language should I use in this group?`,
		"es-ES": `¿Qué idioma debería usar en este grupo?`,      // TODO(ES) verify
		"fa-IR": `کدام زبان باید در این گروه استفاده کنم؟`,      // TODO(FA) verify
		"it-IT": `Che lingua devo utilizzare in questo gruppo?`, // TODO(IT) verify
		"ru-RU": `Какой язык я должен использовать в этой группе?`,
	},
	MESSAGE_TEXT_HI_IN_GROUP_LANG_SET: {
		"en-UK": `Great, I'll be using English.`,
		"es-ES": `Genial, usaré español.`, // TODO(ES) verify
		"de-DE": `Kein Problem, dann schreibe ich auf Deutsch.`,
		"fa-IR": `عالی، من از فارسی استفاده خواهم کرد.`, // TODO(FA) verify
		"it-IT": `Ottimo, userò l'italiano.`,            // TODO(IT) verify
		"ru-RU": `Отлично, я буду использовать русский`,
	},
	SPLITUS_TEXT_ABOUT_ME_AND_CO: {
		"de-DE": `Ich kann helfen, <b>Rechnungen zu teilen</b>. Mein Freund @DebtsTrackerBot passt darauf auf, dass alle Schulden zurückgezahlt werden.`,
		"en-UK": `I help to <b>split bills</b>. My friend @DebtsTrackerBot is tracking paybacks & debts.`,
		"es-ES": `Ayudo a <b>dividir billetes</b>. Mi amigo @DebtsTrackerBot está rastreando pagos y deudas.`,                 // TODO(ES) verify
		"fa-IR": `من به <b>تقسیم صورتحساب </b> کمک میکنم دوست منDebtsTrackerBot ردیابی بازپرداخت و بدهی است.`,                 // TODO(FA) verify
		"it-IT": `Aiuto a <b>dividere le bollette</b>. Il mio amico @DebtsTrackerBot sta monitorando i pagamenti e i debiti.`, // TODO(IT) verify
		"ru-RU": `Я помогаю делить счета. Мой друг @DebtsTrackerRuBot отслеживает платежи и долги.`,
	},
	COLLECTUS_TEXT_ABOUT_ME_AND_CO: {
		"de-DE": `Ich kann dabei helfen <b>Geld zu sammeln</b>. Zum Beispiel für ein Geburtstagsgeschenk. 🎉

Mein Freund @DebtsTrackerBot überwacht Schulden und deren Rückzahlungen.

Wenn du Geld für irgendwas sammeln willst, kann @SplitusBot dir dabei helfen.`,

		"en-UK": `I help to <b>collect money</b> for a good cause. For example for a birthday present. 🎉

My buddy @DebtsTrackerBot is tracking debts & paybacks.

And if you do collective purchases and want to split bills @SplitusBot is here to help.`,
		"es-ES": ``, // TODO(ES)
		"fa-IR": ``, // TODO(FA)
		"it-IT": ``, // TODO(IT)
		"ru-RU": `Я помогаю <b>собирать деньги</b> на что нибудь. Например для подарка на день рождение. 🎉

Мой друг @DebtsTrackerRuBot отслеживает долги и платежи.

А если хотите вести учёт совместных покупок и удобно делить счета вам поможет @SplitusBot.`,
	},
	SPLITUS_TEXT_HI: {
		"de-DE": `Ich bin <b>Splitus</b>.`,
		"en-UK": `I'm <b>Splitus</b>.`,
		"es-ES": ``, // TODO(ES)
		"fa-IR": ``, // TODO(FA)
		"it-IT": ``, // TODO(IT)
		"ru-RU": `Меня зовут <b>Сплитус</b>.`,
	},
	COLLECTUS_TEXT_HI: {
		"de-DE": `Ich bin <b>Collectus</b>.`,
		"en-UK": `I'm <b>Collectus</b>.`,
		"es-ES": ``, // TODO(ES)
		"fa-IR": ``, // TODO(FA)
		"it-IT": ``, // TODO(IT)
		"ru-RU": `Меня зовут <b>Коллектус</b>.`,
	},
	SPLITUS_TG_COMMANDS: {
		"de-DE": ``,
		"en-UK": `<b>Bot commands:</b>
	/groups - List of groups
	/bills - List of outstanding bills
	/menu - Main menu
	/settings - Settings
	/help - Learn how to use bot, report issues, ask questions`,
		"es-ES": ``, // TODO(ES)
		"fa-IR": ``, // TODO(FA)
		"it-IT": ``, // TODO(IT)
		"ru-RU": `<b>Команды для бота:</b>
	/groups - Список групп
	/bills - Список незакрытых платежей
	/menu - Главное меню
	/settings - Настройки
	/help - Узнать как использовать, сообщить о проблеме, задать вопрос`,
	},
	COLLECTUS_TG_COMMANDS: {
		"en-UK": `<b>Bot commands:</b>

	/groups - List of groups
	/fundraisings - List of active fundraisings
	/help - Learn how to use bot, report issues, ask questions`,
		"es-ES": ``, // TODO(ES)
		"fa-IR": ``, // TODO(FA)
		"it-IT": ``, // TODO(IT)
		"ru-RU": `<b>Команды для бота:</b>
	/groups - Список групп
	/fundraisings - Список активных сборов
	/help - Узнать как использовать, сообщить о проблеме, задать вопрос`,
	},
	MESSAGE_TEXT_SEND_HELP_COMMAND_FOR_HELP: { // This is the same for all languages.
		"de-DE": `Sende /help für Hilfe über den Umgang mit diesen Bot.`,
		"en-UK": `Send /help for details on how to use this bot.`,
		"es-ES": ``,
		"fa-IR": ``,
		"it-IT": ``,
		"ru-RU": `Отправьте /help для справки по использованию бота.`,
	},
	MESSAGE_TEXT_HI: { // This is the same for all languages.
		"de-DE": `¡Hola! Hi! Привет! سلام! Hallo!`,
		"en-UK": `¡Hola! Hi! Привет! سلام! Hallo!`,
		"es-ES": `¡Hola! Hi! Привет! سلام! Hallo!`,
		"fa-IR": `¡Hola! Hi! Привет! سلام! Hallo!`,
		"it-IT": `¡Hola! Hi! Привет! سلام! Hallo!`,
		"ru-RU": `¡Hola! Hi! Привет! سلام! Hallo!`,
	},
	MESSAGE_TEXT_HI_USERNAME: { // This is the same for all languages.
		"de-DE": `Hi %v!`,
		"en-UK": `Hi %v!`,
		"es-ES": `¡Hola %v!`,
		"fa-IR": ``,
		"it-IT": ``, // TODO(IT)
		"ru-RU": `Привет %v!`,
	},
	MESSAGE_TEXT_BACK_TO_MAIN_MENU: {
		"de-DE": `Du kannst zurück zum Haupt /menu`,
		"en-UK": `You can go back to main /menu`,
		"es-ES": `Puedes volver al inicio /menú`,
		"fa-IR": `شما میتوانید به /منو ی اصلی مراجعه کنید.`,
		"it-IT": `Puoi tornare al menu' principale tramite /menu`,
		"ru-RU": `Можно вернуться назад в главное /меню`,
	},
	MESSAGE_TEXT_YOUR_SELECTED_PREFERRED_LANGUAGE: { // This is the same for all languages.
		"de-DE": `Bevorzugte Sprache: %s`,
		"en-UK": `Preferred bot language: %s`,
		"es-ES": `Idioma favorito del bot: %s`,
		"fa-IR": `زبان برنامه: %s`,
		"it-IT": `Lingua del bot preferita: %s`,
		"ru-RU": `Выбранный язык бота: %s`,
	},
	MESSAGE_TEXT_ONBOARDING_ASK_TO_CHOOSE_LANGUAGE: {
		"de-DE": `Welche ist deine bevorzugte Sprache?`,
		"en-UK": `What is your preferred language?`,
		"es-ES": `¿cuál es tu idioma preferida?`,
		"fa-IR": `شما چه زبانی را ترجیح می دهید؟`,
		"it-IT": `Qual'e' la tua lingua madre?`,
		"ru-RU": `На каком языке вы хотели бы общаться?`,
	},
	MESSAGE_TEXT_CHOOSE_UI_LANGUAGE: {
		"de-DE": "Auf welcher Sprache würdest du dich gerne mit mir unterhalten?",
		"en-UK": "Which language you would like to talk to me?",
		"es-ES": "¿En cuál idioma preferirías hablar conmigo?",
		"fa-IR": "دوست دارید به چه زبانی با من صحبت کنید؟",
		"it-IT": "In quale lingua preferisci parlarmi?",
		"ru-RU": "На каком языке вы хотели бы общаться со мной?",
	},
	MESSAGE_TEXT_LOCALE_CHANGED: {
		"de-DE": "Du hast die Sprache geändert auf %v",
		"en-UK": "You've switched language to %v",
		"es-ES": "Has cambiado el idioma a %v",
		"fa-IR": "شما زبان را به %v تغییر دادید. ",
		"it-IT": "Hai cambiato lingua in %v",
		"ru-RU": "Вы поменяли язык на %v",
	},
	MESSAGE_TEXT_WHATS_NEXT: {
		"de-DE": "Was als nächstes?",
		"en-UK": "What's next?",
		"es-ES": "¿Algo más?",
		"fa-IR": "اقدام بعدی چیست؟",
		"it-IT": "Cosa posso fare ora?",
		"ru-RU": "Что будем делать дальше?",
	},
	MESSAGE_TEXT_REFERRERS_TITLE: {
		"en-UK": "Our friends:",
		"ru-RU": "Наши друзья:",
	},
	COMMAND_TEXT_ADD_MY_TG_CHANNEL: {
		"en-UK": "Add my channel",
		"ru-RU": "Добавить мой канал",
	},
	MESSAGE_TEXT_DEBTUS_COMMANDS: {
		"en-UK": `<b>Commands for the bot</b>
🏠 /menu - show main menu
🔙 /return - return previously recorded debt
📥 /got - record money you received from others
📤 /gave - record money you gave to others
📚 /history - latest transactions
🏁 /balance - display current balance
⚙ /settings - adjust your preferences`,

		"ru-RU": `<b>Команды бота</b>
🏠 /menu - показать основное меню
🔙 /return - записать возврат долга
📥 /got - записать о взятии в долг
📤 /gave - записать о том что дали взаймы
📚 /history - последние транзакции
🏁 /balance - показать текущий баланс
⚙ /settings - настройки`,
	},
	MESSAGE_TEXT_WHATS_NEXT_HINT: {
		"de-DE": `
Wenn du dir was von jemanden geliehen hast, wähle /anleihen.
Wenn du was an jemanden verliehen hast, wähle /verleihen.

Oder benutzt das Menü unten.`,
		"en-UK": `
If you borrowed from someone to record it use /got.
If you lent to someone to record it use /gave.

Or use menu at the bottom.`,

		"es-ES": `
		Si alguien te ha prestado usa el comando  /recibido.
		Si has prestado a alguien usa el comando /dado.

O usa el menú de abajo en la pantalla.`,

		"fa-IR": `
اگر از کسی قرض گرفته اید برای ثبت آن از /قرض_گرفتن استفاده کنید.
اگر به کسی قرض داده اید برای ثبت آن از /قرض_دادن استفاده کنید.

یا از منوی پایین استفاده نمایید.`,

		"it-IT": `
Se qualcuno ti ha prestato qualcosa per memorizzarlo usa /got.
Se hai prestato qualcosa a qualcuno per memorizzarlo usa /gave.

O usa il menu' qui sotto.`,

		"ru-RU": `
	Если вы дали в долг воспользуйтесь командой /дал.
	Если вы одолжили что-то - командой /взял.

	Или воспользуйтесь меню внизу экрана.`,
	},
	MESSAGE_TEXT_HISTORY_HEADER: {
		"de-DE": "Verlauf",
		"en-UK": "History",
		"es-ES": "Cronología",
		"fa-IR": "سوابق",
		"it-IT": "Cronologia",
		"ru-RU": "История",
	},
	MESSAGE_TEXT_HISTORY_NO_RECORDS: {
		"de-DE": "Du hast noch nichts gespeichert.",
		"en-UK": "You don't have any records yet.",
		"es-ES": "Aún no tienes ninguna notificación.",
		"fa-IR": "شما هنوز هیچ ثبت سابقه ای ندارید.",
		"it-IT": "Non hai nulla memorizzato.",
		"ru-RU": "У вас пока нет ни одной записи.",
	},
	MESSAGE_TEXT_HISTORY_LIST: {

		"de-DE": `<b>%v</b> <i>(bis %d):</i>
─────────────
%v`,
		"en-UK": `<b>%v</b> <i>(last %d):</i>
─────────────
%v`,
		"es-ES": `<b>%v</b> <i>(últimos %d):</i>
─────────────
%v`,

		"fa-IR": `<b>%v</b> <i>(آخرین %d):</i>
─────────────
%v`,

		"it-IT": `<b>%v</b> <i>(ultimi %d):</i>

─────────────
%v`,
		"ru-RU": `<b>%v</b> <i>(%d последних)</i>
─────────────
%v`,
	},
	MESSAGE_TEXT_BALANCE_IS_ZERO: {
		"de-DE": "Keine unbeglichenen Schulden.",
		"en-UK": "You have no records on current debts.",
		"es-ES": "No hay ninguna notificación de deudas actuales.",
		"fa-IR": "شما در خصوص بدهی های اخیر ثبت سابقه ای ندارید.",
		"it-IT": "Non hai nulla memorizzato nel debito corrente.",
		"ru-RU": "Нет записей о текущих долгах.",
	},
	MESSAGE_TEXT_BALANCE_CURRENCY_TOTAL_INTRO: {
		"de-DE": "Insgesamt",
		"en-UK": "Total",
		"es-ES": "Total",
		"fa-IR": "مجموع",
		"it-IT": "Totale",
		"ru-RU": "Всего",
	},
	BT_OTHER_CURRENCY: {
		"en-UK": "Another currency",
		"es-ES": "Otra moneda",     // TODO(es) verify
		"fa-IR": "ارز دیگر",        // TODO(fa) verify
		"it-IT": "Un'altra valuta", // TODO(it) verify
		"ru-RU": "Другая валюта",
	},
	MESSAGE_TEXT_PRIMARY_CURRENCY_IS_SET_TO: {
		"de-DE": "OK, von nun an ist '%v' deine Hauptwährung.",
		"en-UK": "OK, from now on I will use '%v' as a primary currency.",
		"es-ES": "OK, ahora voy a usar '%v' como moneda principal. ",
		"fa-IR": "بسیار خوب، از الان من از '%v' بعنوان واحد پول اولیه استفاده می کنم",
		"it-IT": "OK, da ora in poi usero' '%v' come moneta principale.",
		"ru-RU": "OK, теперь я буду использовать '%v' как основную валюту.",
	},
	MESSAGE_TEXT_BALANCE_SINGLE_CURRENCY_COUNTERPARTY_DEBT_TO_USER: {
		"de-DE": "%v - schuldet dir %v",
		"en-UK": "%v - owes you %v",
		"es-ES": "%v - te debe %v",
		"fa-IR": "%v - %v به شما بدهکار است ",
		"it-IT": "%v - ti deve %v.",
		"ru-RU": "%v - долг вам %v",
	},
	MESSAGE_TEXT_BALANCE_CURRENCY_ROW_DEBT_TO_USER: {
		"de-DE": "Schuldet dir %v",
		"en-UK": "Owes to you %v",
		"es-ES": "Te debe %v",
		"fa-IR": "%v به شما بدهکار است ",
		"it-IT": "Hai un credito di %v",
		"ru-RU": "Вам должны %v",
	},
	MESSAGE_TEXT_ON_RETURN_USER_DOES_NOT_OWE_ANYTHING_TO_COUNTERPARTY_ANYMORE: {
		"de-DE": "Hurra, du bist jetzt quitt mit <b>%v</b>.",
		"en-UK": "Congratulations! You don't owe anything more to <b>%v</b>.",
		"es-ES": "Bravo! Has saldado tu deuda con <b>%v</b>.",
		"fa-IR": "تبریک! شما دیگر چیزی به <b>%v</b> بدهکار نیستید .",
		"it-IT": "Bravo! Hai saldato il tuo debito con <b>%v</b>.",
		"ru-RU": "Поздравляем! У вас не осталось долгов перед <b>%v</b>.",
	},
	MESSAGE_TEXT_ON_RETURN_COUNTERPARTY_DOES_NOT_OWE_ANYTHING_TO_USER_ANYMORE: {
		"de-DE": "Du bist jetzt mit <b>%v</b> quitt.",
		"en-UK": "<b>%v</b> does not owe anything more to you.",
		"es-ES": "<b>%v</b> nadie te debe nada ya.",
		"fa-IR": "<b>%v</b> دیگر چیزی به شما بدهکار نیست",
		"it-IT": "<b>%v</b> ha saldato il suo debito verso di te.",
		"ru-RU": "У <b>%v</b> больше не осталось долгов перед вами.",
	},
	MESSAGE_TEXT_BALANCE_CURRENCY_ROW_DEBT_BY_USER: {
		"de-DE": "Du schuldest %v",
		"en-UK": "You owe %v",
		"es-ES": "Tú debes %v",
		"fa-IR": "شما %v بدهکار هستید ",
		"it-IT": "Hai un debito di %v",
		"ru-RU": "Вы должны %v",
	},
	MESSAGE_TEXT_BALANCE_SINGLE_CURRENCY_COUNTERPARTY_DEBT_BY_USER: {
		"de-DE": "%v - du schuldest %v",
		"en-UK": "%v - you owe %v",
		"es-ES": "%v - tú debes %v",
		"fa-IR": "%v - شما %v بدهکار هستید ",
		"it-IT": "%v - tu gli/le devi %v",
		"ru-RU": "%v - вы должны %v",
	},
	MESSAGE_TEXT_ASK_PRIMARY_CURRENCY: {
		"de-DE": "Was ist deine Hauptwährung?",
		"en-UK": "What is your primary currency?",
		"es-ES": "¿Cuál es tu moneda principal?",
		"fa-IR": "واحد پولی اولیه شما چیست؟",
		"it-IT": "Qual'e' la tua valuta principale?",
		"ru-RU": "Какая валюта для вас основная?",
	},
	MESSAGE_TEXT_ASK_PRIMARY_CURRENCY_FOR_GROUP: {
		"en-UK": "What is a primary currency for this group?",
		"es-ES": "¿Cuál es tu moneda principal?",     //TODO(ES)
		"fa-IR": "واحد پولی اولیه شما چیست؟",         //TODO(FA)
		"it-IT": "Qual'e' la tua valuta principale?", //TODO(IT)
		"ru-RU": "Какая валюта основная для этой группы?",
	},
	MESSAGE_TEXT_FAILED_TO_DELETE_USER: {
		"de-DE": "Konnte den Benutzer nicht löschen: %v",
		"en-UK": "Failed to delete user: %v",
		"es-ES": "Error durante la cancelación del usuario: %v",
		"fa-IR": "حذف کاربر ناموفق بود: %v",
		"it-IT": "Errore durante la cancellazione dell'utente: %v",
		"ru-RU": "Не удалось удалить данные пользователя: %v",
	},
	MESSAGE_TEXT_USER_DELETED: {
		"de-DE": "Die Benutzerdaten wurden gelöscht.",
		"en-UK": "User's data has been deleted",
		"es-ES": "Los datos del usuario han sido eliminados",
		"fa-IR": "اطلاعات کاربر حذف شد.",
		"it-IT": "I dati dell'utente sono stati cancellati",
		"ru-RU": "Данные пользователя удалены",
	},
	MESSAGE_TEXT_PLEASE_WAIT_WHILE_WE_GENERATE_INVITE_CODE: {
		"de-DE": "Einen Moment - der Sicherheitscode wird generiert...",
		"en-UK": "Please wait a moment while we are generating a security access code...",
		"es-ES": "Por favor, espera un momento mientras generamos un código de acceso seguro…",
		"fa-IR": "لطفاً کمی صبر کنید تا ما یک کد دسترسی امنیتی  ایجاد کنیم.",
		"it-IT": "Aspetta un attimo mentre sto generando un codice di accesso sicuro...",
		"ru-RU": "Пожалуйста подождите пока мы генерируем секретный код доступа...",
	},
	MESSAGE_TEXT_RETURN_ASK_TO_CHOOSE_COUNTERPARTY: {
		"de-DE": "Bitte wähle, wer die Schuld beglichen hat oder wem du sie zurückgezahlt hast.",
		"en-UK": "Please choose who returned the debt or to who you returned it.",
		"es-ES": "Por favor, elige quien ha devuelto o a quien ha sido devuelta la deuda ",
		"fa-IR": "لطفاً انتخاب کنید چه کسی بدهی اش را به شما پرداخت کرده یا شما بدهیتان را به چه کسی بازپرداخت نموده اید.",
		"it-IT": "Scegli con chi hai sanato un credito o un debito.",
		"ru-RU": "Выберете кому вы вернули долг или кто вернул его вам.",
	},
	MESSAGE_TEXT_CHOOSE_DEBT_THAT_HAS_BEEN_RETURNED: {
		"de-DE": "Bitte wähle, ob die Schuld vollständig oder teilweise beglichen wurde.",
		"en-UK": "Please choose a debt that has been returned fully or partially.",
		"es-ES": "Por favor, elige una deuda que ha sido devuelta total o parcialmente. ",
		"fa-IR": "لطفاً انتخاب کنید تمام یا بخشی از کدام بدهی پرداخت شده است.",
		"it-IT": "Scegli un debito che e' stato restituito completamente o parzialmente.",
		"ru-RU": "Выберите долг который был возвращён целиком или частично.",
	},
	MESSAGE_TEXT_NO_DEBTS_TO_RETURN: {
		"de-DE": "You have no records for debts that can be returned.",
		"en-UK": "",
		"es-ES": "",
		"fa-IR": "",
		"it-IT": "",
		"ru-RU": "У вас нет записей о догах для возврата.",
	},
	MESSAGE_TEXT_PLEASE_ACKNOWLEDGE_TRANSFER: {
		"de-DE": "Bitte stimme dem zu oder lehne es ab.",
		"en-UK": "Please confirm or decline this transfer.",
		"es-ES": "Por favor, confirma o rechaza la transacción.",
		"fa-IR": "لطفاً این تراکنش را تایید یا رد نمایید.",
		"it-IT": "Conferma o rifiuta questo debito/credito.",
		"ru-RU": "Пожалуйста подтвердите или отклоните эту транзакцию.",
	},
	MESSAGE_TEXT_ALREADY_ACCEPTED_TRANSFER: {
		"de-DE": "Du hast dem bereits zugestimmt.",
		"en-UK": "This transfer has been accepted already.",
		"es-ES": "Esta transacción ya ha sido aceptada.",
		"fa-IR": "این تراکنش قبلا قبول شده است.",
		"it-IT": "Questo debito/credito e' gia' stato accettato.",
		"ru-RU": "Эта транзакция уже подтверждена.",
	},
	MESSAGE_TEXT_ALREADY_DECLINED_TRANSFER: {
		"de-DE": "Du hast dem bereits widersprochen.",
		"en-UK": "This transfer has been declined already.",
		"es-ES": "Esta transacción ya ha sido rechazada.",
		"fa-IR": "این تراکنش قبلاً رد شده است.",
		"it-IT": "Questo debito/credito e' gia' stato rifiutato.",
		"ru-RU": "Эта транзакция уже отклонена.",
	},
	MESSAGE_TEXT_RECEIPT_LINK: {
		"de-DE": "Details hier: %v",
		"en-UK": "Details here: %v",
		"es-ES": "Detalles aquí: %v",
		"it-IT": "Maggiori dettagli qui: %v",
		"fa-IR": "جزئیات: %v",
		"ru-RU": "Подробнее тут: %v",
	},
	MESSAGE_TEXT_ASK_PHONE_NUMBER_OF_COUNTERPARTY: {
		"de-DE": "Bitte gib die Telefonnummer von <b>%v</b> an:",
		"en-UK": "Please provide phone number for <b>%v</b>",
		"es-ES": "Por favor escribe el número de teléfono de <b>%v</b>",
		"it-IT": "Per favore fornisci il numero di telefono di <b>%v</b>",
		"fa-IR": "لطفا شماره تلفن ایشان را وارد کنید <b>%v</b>",
		"ru-RU": "Пожалуйста напишите номер телефона <b>%v</b>.",
	},
	MESSAGE_TEXT_USE_CONTACT_TO_SEND_PHONE_NUMBER: {
		"de-DE": "Wenn die Telefonnummer in deinem Adressbuch ist, kannst du den <b>%v Button benutzen</b>, um einen Kontakt zu senden.",
		"en-UK": "If phone number is in your address book you can <b>use %v button</b> to send the contact.",
		"es-ES": "Si el número está en tu agenda puedes <b>usar %v el botón</b> para enviar el contacto.",
		"fa-IR": "اگر شماره تلفن در فهرست مخاطبین شما وجود دارد شما می توانید <b> با استفاده از این %v دکمه</b> تماس را ارسال نمایید.",
		"it-IT": "Se il numero e' nella tua rubrica, puoi <b> usare il pulsante %v</b> per inviare il contatto.",
		"ru-RU": "Если номер телефона есть в записной книжке <b>воспользуйтесь кнопкой %v</b> (скрепка) чтобы отправить контакт.",
	},
	MESSAGE_TEXT_ABOUT_PHONE_NUMBER_FORMAT: {
		"en-UK": "The number should be in international standard:\n\t* Starts with '+' following by country code\n\t* Consist of numbers only\nExample: <b>+1</b><code>999012345678</code>",
		"es-ES": "El número debe tener formato internacional estándar:\n\t* Empezar con '+' seguido del código del país\n\t* formado solo por números\nEjemplo: <b>+1</b><code>999012345678</code>",
		"de-DE": "Die Telefonnummer sollte dem internationalen Standard entsprechen:\n\t* Beginnend mit '+' gefolgt vom Ländercode (Deutschland +49)\n\t* Consist of numbers only\nExample: <b>+49</b><code>157123456</code>",
		"fa-IR": "شماره باید به صورت استاندارد بین المللی باشد\n\t* با '+' شروع شده و بدنبال آن کد کشور وارد شود\n\t* تنها شامل اعداد باشد\nمثال: <b>+1</b><code>999012345678</code>",
		"it-IT": "Il numero deve essere in formato internazionale:\n\t* Inizia con '+' seguito dal codice del paese (Italia +39)\n\t* \nEsempio: <b>+39</b><code>34612345678</code>",
		"ru-RU": "Номер должен быть в международном формате:\n\t* Начинаться со знака '+' и кода страны\n\t* Состоять только из цифр\nПример: <b>+7</b><code>999012345678</code>",
	},
	MESSAGE_TEXT_THIS_NUMBER_WILL_BE_USED_TO_SEND_RECEIPT: {
		"de-DE": "Wir werden eine SMS an diese Nummer schicken:",
		"en-UK": "Will send an SMS to this number:",
		"es-ES": "Enviaremos una SMS a este número:",
		"fa-IR": "یک پیام کوتاه به این شماره ارسال خواهد شد:",
		"it-IT": "Invieremo un SMS a questo numero:",
		"ru-RU": "На этот номер мы отправим SMS:",
	},
	MESSAGE_TEXT_COUNTERPARTY_OWES_YOU_SINGLE_DEBT: {
		"de-DE": `<b>%v</b> schuldet dir <b>%v</b>.`,
		"en-UK": `<b>%v</b> owes to you <b>%v</b>.`,
		"es-ES": `<b>%v</b> has prestado <b>%v</b>.`,
		"fa-IR": `<b>%v</b> به شما بدهکار بوده <b>%v</b>.`,
		"it-IT": `<b>%v</b> e' in debito di <b>%v</b>.`,
		"ru-RU": `<b>%v</b> одалживал(а) у вас <b>%v</b>.`,
	},
	MESSAGE_TEXT_YOU_OWE_TO_COUNTERPARTY_SINGLE_DEBT: {
		"de-DE": "Du schuldest <b>%v</b> <b>%v</b>.",
		"en-UK": "You owe to <b>%v</b> <b>%v</b>.",
		"es-ES": "Te ha prestado <b>%v</b> <b>%v</b>.",
		"fa-IR": "شما بدهکار هستید به <b>%v</b> <b>%v</b>.",
		"it-IT": `Tu devi dare a <b>%v</b> <b>%v</b>.`,
		"ru-RU": `<b>%v</b> одалживал(а) вам <b>%v</b>.`,
	},
	MESSAGE_TEXT_IS_IT_RETURNED_IN_FULL: {

		"de-DE": `Wurde diese Schuld vollständig beglichen?

		<i>Falls nur teilweise, kann der Teilbetrag direkt eingegeben werden.</i>`,
		"en-UK": `Has this debt been returned in full?

		<i>If partially you can enter amount right away.</i>`,

		"es-ES": `¿Ha sido totalmente devuelta esta deuda?

		<i>si ha sido devuelta parcialmente puedes introducir el importe.</i>`,

		"fa-IR": `آیا این بدهی بصورت کامل بازپرداخت شده است؟

		<i>اگر بخشی از بدهی پرداخت شده است شما میتوانید مبلغ را وارد کنید.</i>`,

		"it-IT": `Il debito e' stato saldato?

		<i>Se la risposta e' NO puoi inserire l'ammontare da sottrarre, direttamente qui sotto.</i>`,
		"ru-RU": `Возвращено полностью?

		<i>Если частично можете сразу написать сумму.</i>`,
	},
	MESSAGE_TEXT_PLEASE_HELP_MAKE_IT_BETTER: {
		"de-DE": `Das Programm ist <b>kostenlos</b>. Bitte <a href="https://debtstracker.io/en/help-us">hilf</a> es besser zu machen!`,
		"en-UK": `This program is <b>free to use</b>. Please <a href="https://debtstracker.io/en/help-us">help</a> to make it better!`,
		"es-ES": `Este programa es <b>gratis</b>. Por favor <a href="https://debtstracker.io/en/help-us">ayúdanos</a> a mejorarlo!`,
		"it-IT": `Questo programma e' <b> completamente gratis</b>. Per favore <a href="https://debtstracker.io/">aiuta</a> a migliorarlo!`,
		"fa-IR": `این برنامه <b>رایگان می باشد</b>.لطفاً <a href="https://debtstracker.io/">به ما کمک کنید</a>تا آنرا بهبود دهیم!`,
		"ru-RU": `Эта программа <b>бесплатна</b>. <a href="https://debtstracker.io/ru/help-us">Помогите</a> сделать её лучше!`,
	},
	BUTTON_TEXT_YOU_OWE_AMOUNT_TO_SOMEONE: {
		"de-DE": "%v | du schuldest: %v",
		"en-UK": "%v | you owe: %v",
		"es-ES": "%v | tú debes: %v",
		"fa-IR": "%v | شما بدهکارید: %v",
		"it-IT": "%v | tu devi: %v",
		"ru-RU": "%v | вы должны: %v",
	},
	BUTTON_TEXT_SOMEONE_OWES_TO_YOU_AMOUNT: {
		"de-DE": "%v | schuldet dir: %v",
		"en-UK": "%v | owes to you: %v",
		"es-ES": "%v | te debe: %v",
		"fa-IR": "%v | به شما بدهکار است: %v",
		"it-IT": "%v | ti deve: %v",
		"ru-RU": "%v | долг вам: %v",
	},
	BUTTON_TEXT_DEBT_RETURNED_FULLY: {
		"de-DE": "Ja, vollständig",
		"en-UK": "Yes, fully",
		"es-ES": "Sí, completamente",
		"fa-IR": "بله، به صورت کامل",
		"it-IT": "Si, completamente",
		"ru-RU": "Да, целиком",
	},
	BUTTON_TEXT_DEBT_RETURNED_PARTIALLY: {
		"de-DE": "Nein, nur teilweise",
		"en-UK": "No, just partially",
		"es-ES": "No, solo parcialmente",
		"fa-IR": "خیر، تنها قسمتی",
		"it-IT": "No, parzialmente",
		"ru-RU": "Нет, только часть",
	},
	MESSAGE_TEXT_ATTEMPT_TO_USE_OWN_INVITE: {
		"de-DE": "Du solltest dich nicht selber einladen ;)",
		"en-UK": "You should not use your own invite ;)",
		"es-ES": "No deberías invitarte a ti mismo ;)",
		"fa-IR": "نباید از دعوت خود استفاده کنید ;)",
		"it-IT": "Non dovresti usare il tuo codice invito con te stesso :)",
		"ru-RU": "Хорошая попытка пригласить самого себя ;)",
	},
	MESSAGE_TEXT_WELCOME_ONBOARDING_INVITE_ACCEPTED: {
		"de-DE": "Willkommen. Schön, dass du der Einladung gefolgt bist!",
		"en-UK": "Welcome and thanks for accepting the invite!",
		"es-ES": "Bienvenido y gracias por aceptar la invitación",
		"fa-IR": "خوش آمدید و سپاسگزاریم که دعوت را پذیرفتید!",
		"it-IT": "Benvenuto e grazie per aver accettato l'invito!",
		"ru-RU": "Спасибо за то что воспользовались приглашением!",
	},
	MESSAGE_TEXT_FOR_COUNTERPARTY_ONLY: {
		"de-DE": "Das darf nur %v.",
		"en-UK": "This action for %v only.",
		"es-ES": "Esta acción está disponible solo para%v.",
		"fa-IR": "این عمل تنها برای %v می باشد.",
		"it-IT": "Questa azione e' disponibile solo per %v.",
		"ru-RU": "Это действие доступно только для %v.",
	},
	BUTTON_TEXT_SEE_RECEIPT_DETAILS: {
		"de-DE": "Quittungsdetails anzeigen",
		"en-UK": "Show receipt details",
		"es-ES": "Mostrar detalles",
		"fa-IR": "جزئیات رسید را نشان بده",
		"it-IT": "Mostra i dettagli del debito/credito",
		"ru-RU": "Показать детали",
	},
	MESSAGE_TEXT_YOU_SELECTED_INVITE_BY_EMAIL: {
		"de-DE": "Du hast gewählt, einen Freund per Mail einzuladen.",
		"en-UK": "You've selected to invite friend by email.",
		"es-ES": "Has decidido invitar a un amigo por e-mail.",
		"fa-IR": "شما انتخاب کردید که یک دوست را بوسیله ایمیل دعوت کنید.",
		"it-IT": "Hai scelto di invitare l'amico tramite email.",
		"ru-RU": "Вы решили пригласить друга через email.",
	},
	MESSAGE_TEXT_YOU_SELECTED_INVITE_BY_SMS: {
		"de-DE": "Du hast gewählt, einen Freund per SMS einzuladen.",
		"en-UK": "You've selected to invite friend by SMS.",
		"es-ES": "Has decidido invitar a un amigo por SMS.",
		"fa-IR": "شما انتخاب کردید که یک دوست را بوسیله پیام کوتاه دعوت کنید",
		"it-IT": "Hai scelto di invitare l'amico tramite SMS.",
		"ru-RU": "Вы решили пригласить друга через SMS.",
	},
	MESSAGE_TEXT_ABOUT_INVITES: {
		"de-DE": "Wie möchtest du den Code weitergeben?",

		"en-UK": `How do you want to pass the invite code?`,

		"es-ES": `¿Cómo quieres enviarle el código?`,

		"fa-IR": `آیا میخواهید کد دعوت را ارسال کنید؟`,

		"it-IT": `Come vuoi inviargli il codice invito?`,

		"ru-RU": `Как вы хотите передать код приглашение?`,
	},
	MESSAGE_TEXT_USER_BLOCKED_TRANSFER_NOTIFICATIONS_BY: {
		"de-DE": "%v hat Erinnerungen über folgendes Anliegen blockiert: %v",
		"en-UK": "%v blocked reminders about transactions by: %v",
		"es-ES": "%v ha bloqueado las notificaciones de las transacciones por: %v",
		"fa-IR": "%v یادآور تراکنش مسدود شده است بوسیله ی: %v",
		"it-IT": "%v bloccato promemoria riguardo la transazione da: %v.",
		"ru-RU": "%v заблокировал получение оповешений о транзакиях через: %v.",
	},
	COMMAND_TEXT_WAIT_A_SECOND: {
		"de-DE": "Warte eine Sekunde...",
		"en-UK": "Wait a second...",
		"es-ES": "Un segundo...",
		"fa-IR": "یک ثانیه صبر کنید ...",
		"it-IT": "Solo un attimo...",
		"ru-RU": "Секундочку...",
	},
	HTML_USING_TELEGRAM: {
		"de-DE": "benutze Telegram messenger",
		"en-UK": "using Telegram messenger",
		"es-ES": "Usando Telegram",
		"fa-IR": "استفاده از پیام رسان تلگرام",
		"it-IT": "usa Telegram",
		"ru-RU": "используя Telegram мессенджер",
	},
	COMMAND_TEXT_ACCEPT: {
		"de-DE": "Akzeptieren",
		"en-UK": "Accept",
		"es-ES": "Aceptar",
		"fa-IR": "قبول",
		"it-IT": "Accetta",
		"ru-RU": "Согласиться",
	},
	//BUTTON_TEXT_ACCEPT_TRANSFER_USING_TELEGRAM:{
	//	"en-UK": "Accept ",
	//	"es-ES": "Confirmar ",

	//	"fa-IR": "قبول ",

	//  "it-IT": "Accetta",
	//	"ru-RU": "Подтвердить ",

	//},
	//BUTTON_TEXT_DECLINE_TRANSFER_USING_TELEGRAM:{
	//	"en-UK": "Decline (using Telegram messenger)",
	//	"es-ES": "Rechazar (usandoTelegram)",

	//	"fa-IR": "رد ( استفاده از پیام رسان تلگرام)",

	//  "it-IT": "Rifiuta (usando Telegram)",
	//	"ru-RU": "Отказаться (используя Telegram)",

	//},
	COMMAND_TEXT_DECLINE: {
		"de-DE": "Ablehnen",
		"en-UK": "Decline",
		"es-ES": "Rechazar",
		"fa-IR": "رد",
		"it-IT": "Rifiuta",
		"ru-RU": "Отклонить",
	},
	COMMAND_TEXT_ACCEPT_INVITE: {
		"de-DE": "Akzeptiere Einladung",
		"en-UK": "Accept invite",
		"es-ES": "Aceptar la invitación",
		"fa-IR": "قبول دعوت",
		"it-IT": "Accetta l'invito",
		"ru-RU": "Принять приглашение",
	},
	COMMAND_TEXT_VIEW_RECEIPT_DETAILS: {
		"de-DE": "Quittungsdetails anzeigen",
		"en-UK": "See receipt details",
		"es-ES": "Ver el recibo",
		"fa-IR": "دیدن جزئیات رسید",
		"it-IT": "Vedi dettagli",
		"ru-RU": "Посмотреть квитанцию",
	},
	COMMAND_TEXT_OTHER_WAYS_TO_SEND_INVITE: {
		"de-DE": "Andere Wege, eine Einladung zu senden",
		"en-UK": "Other ways to send an invite",
		"es-ES": "Otras maneras para enviar la invitación",
		"fa-IR": "سایر راههای ارسال دعوت",
		"it-IT": "Altri modi per inviare un invito",
		"ru-RU": "Другие способы отправить приглашение",
	},
	COMMAND_TEXT_SEND_MY_PHONE_NUMBER: {
		"de-DE": "meine Telefonnummer senden",
		"en-UK": "Send my phone number",
		"es-ES": "Enviar mi número",
		"fa-IR": "شماره تلفن مرا ارسال کنید",
		"it-IT": "Invia il mio numero",
		"ru-RU": "Отправить мой номер",
	},
	COMMAND_TEXT_SEND_BY_EMAIL: {
		"de-DE": "per Mail",
		"en-UK": "By Email",
		"es-ES": "Vía e-mail",
		"fa-IR": "بوسیله ی ایمیل",
		"it-IT": "Tramite email",
		"ru-RU": "Через Email",
	},
	COMMAND_TEXT_SEND_BY_SMS: {
		"de-DE": "per SMS",
		"en-UK": "By SMS",
		"es-ES": "Vía SMS",
		"it-IT": "Tramite SMS",
		"fa-IR": "بوسیله پیام کوتاه",
		"ru-RU": "Через SMS",
	},
	COMMAND_TEXT_INVITE_BY_TELEGRAM: {
		"de-DE": "Einladen per Telegram",
		"en-UK": "Invite By Telegram",
		"es-ES": "Invitar vía Telegram",
		"it-IT": "Tramite Telegram",
		"fa-IR": "دعوت با تلگرام",
		"ru-RU": "Пригласить через Telegram",
	},
	MESSAGE_TEXT_INVITE_CREATED: {

		"de-DE": `Wir haben deinen Freund einen Code geschickt. (#%v)

Sobald dein Freund die Einladung akzeptiert hat, könnt ihr das Geld, was ihr euch teit, mit Leichtigkeit managen.`,

		"en-UK": `We've sent an invite code to your friend. (#%v)

Once your friend accepts invitation you'll share balance & transfers between you to make sure you both are on the same page with minimum efforts.`,

		"es-ES": `Hemos enviado el código de la invitación a tu amigo. (#%v)

Cuando tu amigo accepte la invitación vaís a tener transacciones y balance en común (solo entre vosotros). Todo eso os ayuda minimizar los esfuerzos para controlar la cuenta.`,

		"fa-IR": `ما برای دوست شما یک  پیام دعوت ارسال کردیم. (#%v)

وقتی دوست شما دعوت را بپذیرد شما تراز و مبادلات بین خود را به اشتراک می گذارید تا با کمترین تلاش از درک مشترک میان خود اطمینان حاصل کنید. `,

		"it-IT": `Abbiamo inviato il codice invito al tuo amico. (#%v)

Una volta che il tuo amico accetta l'invito potrete condividere i bilanci ed i trasferimenti con il minimo sforzo.`,
		"ru-RU": `Мы отправили код приглашения на указынный вами адрес. (#%v)

Когда ваш друг потдвердит приглашение у вас будут общий баланс и транзакции (только между вами). Это поможет вам минимизировать усилия по ведению учёта.`,
	},
	MESSAGE_TEXT_INVITE_BY_EMAIL: {
		"de-DE": "Bitte gib die eMail Adresse deines Freundes an, wohin wir den Code schicken:",
		"en-UK": "Please enter emaill address of your friend where we should send an invite code.",
		"es-ES": "Por favor, introduce el e-maill de tu amigo al cual enviaremos el código de la invitación.",
		"fa-IR": "لطفاً آدرس ایمیل دوست خود را وارد کنید تا ما از آن طریق کد دعوت را ارسال نماییم.",
		"it-IT": "Inserisci l'email dell'amico al quale inviare il codice invito.",
		"ru-RU": "Введите email вашего друга на который мы отправим код приглашения.",
	},
	MESSAGE_TEXT_INVITE_ASK_EMAIL_FOR_RECEIPT: {
		"de-DE": "Bitte gib die eMail Adresse deines Freundes (%v), wohin wir die Quittung schicken:", // TODO(DE) - verify
		"en-UK": "Please enter emaill address of your friend (%v) where we should send the receipt.",
		"es-ES": "Introduce el e-maill de tu amigo (%v) al cual enviaremos el recibo.",
		"fa-IR": "لطفا آدرس ایمیل دوست خود را وارد کنید (%v) تا ما از آن طریق کد دعوت را ارسال نماییم.",
		"it-IT": "Inserisci l'email del tuo amico (%v) alla quale potremo inviare il debito/credito",
		"ru-RU": "Введите email вашего друга (%v) на который мы отправим квитанцию подтверждения.",
	},
	MESSAGE_TEXT_INVITE_BY_SMS: {
		"de-DE": "Bitte gib die Telefonnummer deines Freundes an oder teile einen Kontakt, wohin wir den Code schicken:",
		"en-UK": "Please enter number of yoru friend where we'll send an invite.",
		"es-ES": "Por favor, introduce el número del teléfono de tu amigo al cual enviaremos el código de la invitación.",
		"fa-IR": "لطفا شماره دوستان را که می خواهید برای او کد دعوت بفرستیم از لیست مخاطبین به اشتراک گذاشته یا آن را وارد کنید.",
		"it-IT": "Condividi il contatto o inserisci il numero di telefono del tuo amico al quale invieremo il codice invito.",
		"ru-RU": "Введите номер телефона вашего друга на который мы отправим код приглашения.",
	},
	MESSAGE_TEXT_INVITE_BY_TELEGRAM: {
		"de-DE": "Bitte wähle einen Kontakt, welchen du den Code schicken willst:",
		"en-UK": "Please share a contact of your friend you wish to send an invite code:",
		"es-ES": "Por favor, comparte el contacto de tu amigo al cual quieres enviar el código de la invitación.",
		"fa-IR": "لطفا اطلاعات تماس دوستتان را که میخواهید برایشان کد دعوت ارسال شود به اشتراک بگذارید.",
		"it-IT": "Condividi il contatto di un amico al quale desideri inviare il codice invito.",
		"ru-RU": "Вставьте в чат контакт вашего друга которому вы хотите отправить приглашение.",
	},
	MESSAGE_TEXT_INVALID_EMAIL: {
		"de-DE": "Ungültige eMail Adresse. Versuche es erneut oder gehe ins Haupt /menu",
		"en-UK": "Invalid email. Check and try it again? /menu",
		"es-ES": "Email incorrecto. ¿Comprobarlo e intentalo de nuevo? /menú",
		"fa-IR": "ایمیل غیر معتبر می باشد. آیا بررسی نموده، دوباره سعی می کنید؟ /منو",
		"it-IT": "Email scorretta. Controlla e riprova. /menu",
		"ru-RU": "Неверный email. Проверить и попробовать ещё раз? /menu",
	},
	MESSAGE_TEXT_INVALID_YEAR: {
		"de-DE": "Ungültiges Jahr. Der Jahresangabe sollte 2 oder 4 Ziffern sein (<i>z.B. 2016 or 16</i>).",
		"en-UK": "Invalid year. The year part should be 2 or 4 numbers (<i>e.g. 2016 or 16</i>).",
		"es-ES": "Año incorrecto. El año tiene que constar de 2 o 4 números (<i>ejemplo 2016 o 16</i>).",
		"fa-IR": "سال غیرمعتبر می باشد. سال باید به صورت 2 یا 4 رقمی وارد شود (<i>برای مثال 16 یا 2016</i>).",
		"it-IT": "Anno scorretto. L'anno dev;essere composta da 2 o 4 numeri (<i>esempio 2017 oppure 17</i>)",
		"ru-RU": "Неверный год. Год должен быть 2 или 4 цифры (<i>например 2016 или 16)</i>).",
	},
	MESSAGE_TEXT_INVALID_MONTH: {
		"de-DE": "Ungültiger Monat. Der Monatsangabe sollte eine Ganzzahl zwischen 1 und 12 sein.",
		"en-UK": "Invalid month. Month should be an integer from 1 to 12.",
		"es-ES": "El mes es incorrecto. El mes hay que introducirlo del 1 al 12.",
		"fa-IR": "ماه غیر معتبر می باشد. ماه باید به صورت عددی صحیح بین 1 تا 12 باشد.",
		"it-IT": "Mese scorretto. Il mese dovrebbe essere un numero da 1 a 12.",
		"ru-RU": "Неверный месяц. Месяц должен быть задан целым числом от 1 до 12.",
	},
	MESSAGE_TEXT_INVALID_DAY: {
		"de-DE": "Ungültiger Tag. Der Tagesangabe sollte eine Ganzzahl zwischen 1 und 31 sein.",
		"en-UK": "Invalid day. The day should be an integer from 1 to 31.",
		"es-ES": "El día es incorrecto. El día hay que introducirlo del 1 al 31.",
		"fa-IR": "روز غیر معتبر می باشد. روز باید عددی صحیح بین 1 تا 31 باشد.",
		"it-IT": "Giorno scorretto. Il giorno dovrebbe essere un numero da 1 a 31.",
		"ru-RU": "Неверный день. День должен быть задан целым числом от 1 до 31.",
	},
	MESSAGE_TEXT_INVALID_DATE: {
		"de-DE": "Ungültiges Datumsformat. Zum Beispiel für den 20. February 2019 sende: 20.02.2019 oder 20.02.19",
		"en-UK": "Invalid date format. For example for 20 February 2019 please submit: 20.02.2019 or 20.02.19",
		"es-ES": "El formato de la fecha no es correcto. Por ejemplo para el día 20 de Febrero de 2019 introduce: 20.02.2019 o 20.02.19",
		"fa-IR": "فرمت تاریخ غیر معتبر می باشد. برای مثال برای 20 فوریه 2019 لطفا اینگونه وارد کنید: 20.02.2019 یا 20.02.19",
		"it-IT": "Formato data sbagliato. Esempio: per il 20 Febbraio 2019 inserisci: 20.02.2019 oppure 20.02.19",
		"ru-RU": "Неверный формат даты. Например для 20 февраля 2019 года надо ввести: 20.02.2019 или 20.02.19",
	},
	MESSAGE_TEXT_INVALID_PHONE_NUMBER: {
		"de-DE": "Ungültige Telefonnummer. Versuche es erneut oder gehe ins Haupt /menu",
		"en-UK": "Invalid phone number. Check and try it again? /menu",
		"es-ES": "El número del teléfono no es correcto. ¿Comprobarlo y intentarlo de nuevo? /menú",
		"fa-IR": "شماره تلفن غیر معتبر می باشد. آیا بررسی نموده، مجدداً سعی می کنید؟ /منو",
		"it-IT": "Numero di telefono invalido. Controlla e riprova. /menu",
		"ru-RU": "Неверный номер. Проверить и попробовать ещё раз? /menu",
	},
	MESSAGE_TEXT_PHONE_NUMBER_IS_NOT_SMS_CAPABLE: {
		"de-DE": "Diese Telefonnummer kann von uns keine SMS empfangen. Versuche eine andere oder gehe ins Haupt /menu",
		"en-UK": "This phone number not able to receive SMS. Try another number? /menu",
		"es-ES": "Este número de teléfono no acepta SMS. ¿Intentar otro número? /menú",
		"fa-IR": "این شماره تلفن قادر به دریافت پیام کوتاه نمی باشد. آیا شماره دیگری را امتحان میکنید؟ /منو",
		"it-IT": "Questo numero di telefono non e' abilitato a ricevere SMS. Vuoi provare un altro numero? /menu",
		"ru-RU": "Данный номер не принимает SMS. Попробовать другой номер? /menu",
	},
	MESSAGE_TEXT_NO_CONTACT_RECEIVED: {
		"de-DE": "Wir haben keine Kontakte empfangen. Versuche es erneut oder gehe ins Haupt /menu",
		"en-UK": "We have not received any contacts. INSTRUCTION HOW TO DO IT. /menu",
		"es-ES": "No hemos recibido ningún contacto. LA INSTRUCCIÓN COMO HACERLO. /menú",
		"fa-IR": "ما هیچ اطلاعات تماسی دریافت نکردیم. دستورالعمل چگونگی انجام این کار. /منو",
		"it-IT": "Non abbiamo ricevuto nesusn contatto. ISTRUZIONI SU COME FARE. /menu",
		"ru-RU": "Мы не получили контакта. Тут инструкция как это сделать. /menu",
	},
	MESSAGE_TEXT_YOU_HAVE_NO_CONTACTS: {
		"de-DE": "Du hast noch keine Kontakte hinzugefügt.",
		"en-UK": "You have not created any contacts yet.",
		"es-ES": "Todavía no has creado ningún contacto.", //TODO:es - verify
		"fa-IR": "هنوز هیچ مخاطبی را ایجاد نکرده اید",     //TODO:fa - verify
		"it-IT": "Non hai ancora creato alcun contatto.",  //TODO:it - verify
		"ru-RU": "Вы ещё не создали контактов",
	},
	MESSAGE_TEXT_CONTACT_NAME_IS_NUMBER: {
		"de-DE": "Du kannst nicht nur Zahlen für einen Kontaktnamen eingeben. Bitte benutz ein paar Buchstaben - du kannst mir vertrauen.",
		"en-UK": "You've entered just digits for a contact name. Please use some text characters.",
		"es-ES": "Has introducido solo números para el nombre del contacto. Por favor usa algunas letras.",
		"fa-IR": "شما تنها اعداد را برای نام مخاطب وارد کرده اید. لطفا کاراکتر های متنی وارد کنید.",
		"it-IT": "Hai inserito solamente numeri per un nome contatto. Usa anche alcune lettere.",
		"ru-RU": "Вы ввели только цифры в качестве имени контакта. Пожалуйста используйте текстовые символы.",
	},
	MESSAGE_TEXT_CURRENCY_NAME_IS_NUMBER: {
		"de-DE": "Bei der Währung erwarte ich eigentlich keine Zahlen. Nimm ein paar Buchstaben hinzu, um Verwirrung zu vermeiden.",
		"en-UK": "You've entered just digits for currency. Please use some text characters.",
		"es-ES": "Has introducido solo números para la moneda. Por favor usa algunas letras.",
		"fa-IR": "شما تنها اعداد را برای واحد پولی وارد کرده اید. لطفا کاراکترهای متنی وارد کنید.",
		"it-IT": "Hai inserito solamente numeri per la valuta. Usa anche alcune lettere.",
		"ru-RU": "Вы ввели только цифры в качестве номинала. Пожалуйста используйте текстовые символы.",
	},
	MESSAGE_TEXT_HISTORY_ROW_TO_USER_WITH_NAME: {
		"de-DE": "%v - %s ⇒ dir: %s",
		"en-UK": "%v - %s ⇒ to you: %s",
		"es-ES": "%v - %s ⇒ a ti: %s",
		"fa-IR": "%v - %s ⇒ به شما: %s",
		"it-IT": "%v - %s ⇒ a te: %s",
		"ru-RU": "%v - %s ⇒ Вам : %s",
	},
	MESSAGE_TEXT_HISTORY_ROW_FROM_USER_WITH_NAME: {
		"de-DE": "%v - Du ⇒ %s : %s",
		"en-UK": "%v - You ⇒ %s : %s",
		"es-ES": "%v - Tú ⇒ %s : %s",
		"fa-IR": "%v - شما ⇒ %s : %s",
		"it-IT": "%v - Tu ⇒ %s : %s",
		"ru-RU": "%v - Вы ⇒ %s : %s",
	},
	MESSAGE_TEXT_LETS_SEND_SMS: {
		"de-DE": "Lass uns eine SMS senden",
		"en-UK": "Let's send SMS",
		"es-ES": "Vamos a enviar un SMS",
		"fa-IR": "پیام کوتاه ارسال کنید",
		"it-IT": "Inviamo un SMS",
		"ru-RU": "Давайте отправим SMS",
	},
	MESSAGE_TEXT_SMS_QUEUING_FOR_SENDING: {
		"de-DE": "Setze SMS in Sendewarteschlange für %v...",
		"en-UK": "Queuing SMS for sending to number %v...",
		"es-ES": "El SMS se está poniendo en la cola para enviar al número %v...",
		"fa-IR": "پیام کوتاه برای ارسال به شماره مقابل در حال قرارگیری در نوبت ارسال می باشد %v...",
		"it-IT": "SMS in coda per l'invio al numero %v...",
		"ru-RU": "SMS ставится в очередь на отправку на номер %v...",
	},
	MESSAGE_TEXT_SMS_QUEUED_FOR_SENDING: {
		"de-DE": "SMS in Sendewarteschlange für %v",
		"en-UK": "SMS is queued for sending to number %v",
		"es-ES": "El SMS está en la cola para enviar al número %v",
		"fa-IR": "پیام کوتاه برای شماره مقابل در نوبت ارسال قرار گرفت %v",
		"it-IT": "SMS inserito in coda per l'invio al numero %v",
		"ru-RU": "SMS поставлена в очередь на отправку на номер %v",
	},
	MESSAGE_TEXT_BALANCE_HEADER: {
		"de-DE": "Ausstehend",
		"en-UK": "Balance",
		"es-ES": "Balance",
		"fa-IR": "تراز",
		"it-IT": "Bilancio",
		"ru-RU": "Баланс",
	},
	MESSAGE_TEXT_RECEIPT_AVAILABLE_CHANNELS: {
		"de-DE": "Entschuldigung, im Moment funktionieren nur ein paar Kanäle für den Versand von Quittungen:",
		"en-UK": "Sorry, at the moment just this channels are available for sending a receipt:",
		"es-ES": "Disculpa, de momento solo estos canales están disponibles para enviar el recibo:",
		"fa-IR": "متاسفانه، در حال حاضر تنها این کانالها برای ارسال رسید در دسترس می باشند.",
		"it-IT": "Spiacenti ma al momento solo questi canali sono disponibili per inviare debiti/crediti:",
		"ru-RU": "Извините, в данный момент доступны только эти каналы для отправки квитанции:",
	},
	MESSAGE_TEXT_RECEIPT_SENT_THROW_TELEGRAM: {
		"de-DE": "📤 Quittung per Telegram verschickt.",
		"en-UK": "📤 Receipt sent through Telegram.",
		"es-ES": "📤 El recibo ha sido enviado vía Telegram.",
		"fa-IR": "📤 رسید از طریق تلگرام ارسال شد.",
		"it-IT": "📤 Notifica inviata tramite Telegram.",
		"ru-RU": "📤 Квитанция отправлена через телеграм.",
	},
	MESSAGE_TEXT_RECEIPT_NOT_SENT_AS_COUNTERPARTY_HAS_DISABLED_TG_BOT: {
		"de-DE": "Quittung konnte nicht per Telegram gesendet werden, da %v den Chat mit dem Bot gelöscht hat.",
		"en-UK": "Receipt NOT sent through Telegram as %v has deleted chat with the bot.",
		"es-ES": "El recibo NO ha sido enviado vía Telegram porque %v ha eliminado el chat del bot.",
		"fa-IR": "از آنجایی که %v چت انجام شده با روبات را حذف کرده است رسید از طریق تلگرام ارسال نشد.",
		"it-IT": "Notifica NON inviata tramite Telegram a %v perche' ha cancellato la chat con il bot",
		"ru-RU": "Квитанция НЕ отправлена через телеграм так как %v удалил чат с ботом.",
	},
	MESSAGE_TEXT_RECEIPT_SENT_THROW_EMAIL: {
		"de-DE": "Quittung wurde per Mail gesendet. (id: %v)",
		"en-UK": "Receipt sent through email. (id: %v)",
		"es-ES": "El recibo ha sido enviado vía e-mail. (id: %v)",
		"fa-IR": "رسید از طریق ایمیل ارسال شد. (id: %v)",
		"it-IT": "Notifica inviata tramite email (id: %v)",
		"ru-RU": "Квитанция отправлена через email. (id: %v)",
	},
	MESSAGE_TEXT_RECEIPT_SENT_THROW_SMS: {
		"de-DE": "Quittung wurde per SMS gesendet.",
		"en-UK": "Receipt sent through SMS",
		"es-ES": "El recibo ha sido enviado vía SMS.",
		"fa-IR": "رسید از طریق پیام کوتاه ارسال شد.",
		"it-IT": "Notifica inviata tramite SMS",
		"ru-RU": "Квитанция отправлена через SMS.",
	},
	MESSAGE_TEXT_SWITCH_TO_PM_TO_VIEW_RECEIPT: {
		"de-DE": "Schalte in den Privatmodus, um die Quittungsdetails zu sehen",
		"en-UK": "Switch to private mode to see receipt details.",
		"es-ES": "Pasar al modo privado para ver el recibo.",
		"fa-IR": "انتقال به حالت خصوصی جهت رویت جزئیات رسید.",
		"it-IT": "Passa alla modalita' privata per vedere i dettagli dei tuoi crediti/debiti.",
		"ru-RU": "Переключитьсь на чат с ботом чтобы посмотреть квитанцию",
	},
	MESSAGE_TEXT_RECEIPT_ATTEMPT_TO_VIEW_OWN: {
		"en-UK": "An attempt to view own receipt.",
		"ru-RU": "Попытка посмотреть свою собственную квитанцию.",
	},
	MESSAGE_TEXT_RECEIPT_VIEWED_BY_COUNTERPARTY: {
		"de-DE": "👓 Quittung wurde von Gegenpartei gesichtet",
		"en-UK": "👓 Receipt viewed",
		"es-ES": "👓 El recibo ha sido visto",
		"fa-IR": "👓 رسید رویت شد.",
		"it-IT": "👓 Debiti visti",
		"ru-RU": "👓 Квитанция просмотрена",
	},
	MESSAGE_TEXT_VIEW_MY_NUMBER_IN_INTERNATIONAL_FORMAT: {
		"de-DE": "Du kannst deine eigene Nummer in dem Format anzeigen, welches wir erwarten.",
		"en-UK": "You can view your own phone number in the format we are expecting.",
		"es-ES": "Puedes ver tu número de teléfono en el formato previsto.",
		"fa-IR": "شما می توانید شماره تلفن خود را با فرمتی که ما انتظار داریم ببینید.",
		"it-IT": "Puoi visionare il tuo numero di telefono nel formato previsto.",
		"ru-RU": "Вы можете посмотреть свой номер телефона в ожидаемом нами формате.",
	},
	COMMAND_TEXT_VIEW_MY_NUMBER_IN_INTERNATIONAL_FORMAT: {
		"de-DE": "Zeige meine Nummer im erwarteten Format.",
		"en-UK": "View my number in the expectd format",
		"es-ES": "Mostrar mi número de teléfono en el formato previsto",
		"fa-IR": "رویت شماره خودم با فرمت مورد انتظار",
		"it-IT": "Mostra il mio numero nel formato previsto",
		"ru-RU": "Посмотреть мой номер в ожидаемоем формате",
	},
	INLINE_BUTTON_SHOW_FULL_HISTORY: {
		"de-DE": "Lade den ganzen Verlauf",
		"en-UK": "Show full history",
		"es-ES": "Mostrar cronología completa",
		"fa-IR": "نمایش کامل سوابق",
		"it-IT": "Mostra cronologia completa",
		"ru-RU": "Показать всю историю",
	},
	MESSAGE_TEXT_INCORRECT_VALUE_NOT_A_NUMBER: {
		"de-DE": "Das ist keine Nummer",
		"en-UK": "it is not a number",
		"es-ES": "No es un número",
		"fa-IR": "این یک شماره نیست",
		"it-IT": "Non e' un numero",
		"ru-RU": "Это не цифра",
	},
	MESSAGE_TEXT_INCORRECT_VALUE_IS_NEGATIVE: {
		"de-DE": "Die Nummer sollte positiv sein (<i>größer als 0</i>)",
		"en-UK": "The number should be positive (<i>greater than 0</i>)",
		"es-ES": "El número tiene que ser positivo (<i>más de  0</i>)",
		"fa-IR": "شماره باید مثبت باشد (<i>بزرگتر از 0</i>)",
		"it-IT": "Il numero dovrebbe essere positivo (<i>maggiore di 0</i>)",
		"ru-RU": "Цифра должна быть положительной (<i>больше нуля</i>)",
	},
	MESSAGE_TEXT_ASK_HOW_MUCH_HAS_BEEN_RETURNED: {
		"de-DE": "Wie viel wurde beglichen?",
		"en-UK": "How much have been returned?",
		"es-ES": "¿Cuánto/s te han devuelto?",
		"fa-IR": "چه مقدار بازپرداخت شده است؟",
		"it-IT": "Quanto ti e' stato restituito?",
		"ru-RU": "Сколько было возвращено?",
	},
	MESSAGE_TEXT_RETURN_IS_TOO_BIG: {
		"de-DE": "Sie haben entschieden, %v zurückzugeben, aber der ausstehende Betrag ist nur %v.\n\nBitte geben Sie Werte gleich %v oder weniger ein.", // TODO(DE) verify
		"en-UK": "You decided to return %v but outstanding amount is just %v.\n\nPlease enter values equal to %v or less.",
		"es-ES": "Decidiste devolver %v pero la cantidad pendiente es solo %v.\n\nPor favor ingrese valores iguales a %v o menos.", // TODO(ES) verify
		"fa-IR": "شما تصمیم گرفتید %v را بازگردانید اما مقدار قابل توجهی فقط %v است.\n\nلطفا مقادیر برابر %v یا کمتر را وارد کنید", // TODO(FA) verify
		"it-IT": "Hai deciso di restituire %v ma la quantità in sospeso è solo %v.\n\nInserisci valori pari o uguali a %v o meno.", // TODO(IT) verify
		"ru-RU": "Вы решили вернуть %v, но непогашенная сумма равна %v. \n\n Пожалуйста, введите значение равное %v или меньше.",
	},
	MESSAGE_TEXT_HELP_ROOT: {
		"de-DE": "Was hast du für eine Frage? Wenn irgendwas unklar ist, frag ruhig hier @%v",
		"en-UK": "What is your question? If anything is missed here, feel free to ask in our @%v",
		"es-ES": "¿Cuál es tu pregunta? Si algo se pierde aquí, siéntase libre de preguntar en nuestro @%v",   // TODO(es) verify
		"fa-IR": "سوالت چیست؟ اگر چیزی در اینجا از دست رفته است، لطفا در @%v ما بپرسید",                       // TODO(fa) verify
		"it-IT": "Qual è la tua domanda? Se qualche cosa è mancato qui, non esitate a chiedere al nostro @%v", // TODO(it) verify
		"ru-RU": "Какой у вас вопрос? Если здесь нет ответа пожалуйста спросите в нашей группе @%v",
	},
	MESSAGE_TEXT_HELP_BACK_TO_ROOT: {
		"de-DE": "Zurück zur FAQ Liste",
		"en-UK": "Back to FAQ list",
		"es-ES": "Volver a la lista de preguntas frecuentes", // TODO(es) verify
		"fa-IR": "بازگشت به لیست سوالات متداول",              // TODO(fa) verify
		"it-IT": "Torna all'elenco delle FAQ",                // TODO(it) verify
		"ru-RU": "Назад к списку вопросов",
	},
	HELP_HOW_TO_CREATE_BILL_Q: {
		"de-DE": "Wie erstellt man Rechnungen?",
		"en-UK": "How to create new bill?",
		"es-ES": "¿Cómo crear una nueva factura?", // TODO(es) verify
		"fa-IR": "چگونه برای ایجاد لایحه جدید؟",   // TODO(fa) verify
		"it-IT": "Come creare un nuovo conto?",    // TODO(it) verify
		"ru-RU": "Как создать новый счёт?",
	},
	HELP_HOW_TO_CREATE_BILL_A: {
		"de-DE": `<b>Wie man eine Rechnung erstellt</b>
<pre>Rechnung — Kosten, die unter zwei oder mehr Personen geteilt werden.</pre>

Deswegen kannst du am besten direkt <b>im Telegram Chat eine Rechnung erstellen, in zwei Schritten</b>:
<i>Benutze die Funktion "Teile Rechnung mit Telegram Benutzer", um es schnell zu machen:</i>

	1. <a href="https://t.me/{{.BotCode}}?startgroup=from-help">Füg mich einer Gruppe hinzu</a> oder öffne einen Chat mit deinem Freund.

	2. Schreibe <code>@splitusbot {amount} {bill_name}</code> und wähle ein passendes Ergebnis. Zum Beispiel:
<pre>		@splitusbot 45.5 pizza</pre>
	   Und dann kann jeder, der die Rechnung teilen will mit <code>Join</code> einsteigen.

<b>Alternativ</b> kannst du die Rechnung auch direkt in @{{.BotCode}} erstellen. Aber dann musst du die Personen, mit denen du die Rechnung teilst, einzeln hinzufügen.`,
		"en-UK": `<b>How to create a new bill</b>
<pre>Bill — shared expense between two or more people.</pre>

That is why the best is to <b>create bill in Telegram chat just in 2 steps</b>:
<i>use "Split bill with Telegram user(s)" to do it quickly</i>

	1. <a href="https://t.me/{{.BotCode}}?startgroup=from-help">Add me to Telegram group</a> or open chat with a friend.

	2. Type <code>@splitusbot {amount} {bill_name}</code> and select result from menu. For example:
<pre>		@splitusbot 45.5 pizza</pre>
	   Than any member of the group can share the bill by pressing <code>Join</code> button.

<b>Alternatively</b> you can create a bill right in the @{{.BotCode}}. But then you would need manually to add participants.`,
		"es-ES": "", // TODO(ES)
		"fa-IR": "", // TODO(FA)
		"it-IT": "", // TODO(IT)
		"ru-RU": `<b>How to create a new bill</b>
<pre>Bill — shared expense between two or more people.</pre>

That is why the best is to <b>create bill in Telegram chat just in 2 steps</b>:
<i>use "Split bill with Telegram user(s)" to do it quickly</i>

	1. <a href="https://t.me/{{.BotCode}}?startgroup=from-help">Add me to Telegram group</a> or open chat with a friend.

	2. Type <code>@splitusbot {amount} {bill_name}</code> and select result from menu. For example:
<pre>		@splitusbot 45.5 pizza</pre>
	   Than any member of the group can share the bill by pressing <code>Join</code> button.

<b>Alternatively</b> you can create a bill right in the @{{.BotCode}}. But then you would need manually to add participants.`,
	},
	MESSAGE_TEXT_HELP: {
		"de-DE": "Bitte melde jedes Problem und jeden Wunsch auf unserer Webseite.",
		"en-UK": "Please report any issue or submit a feature request at our website.",
		"es-ES": "Puedes informarnos de algún problema o proponernos cualquier mejora en nuestra web.",
		"fa-IR": "لطفاً در وب سایت ما هرگونه مسئله ای را اعلام فرموده یا درخواست خود را مطرح نمایید.",
		"it-IT": "Segnala un problema o proponi un miglioramento sul nostro sito web.",
		"ru-RU": "Вы можете сообщить о проблеме или сделать предложения по улучшению программы на нашем сайте.",
	},
	COMMAND_TEXT_OPEN_USER_REPORT: {
		"de-DE": "Support Seite",
		"en-UK": "Support page",
		"es-ES": "La página de ayuda",
		"fa-IR": "صفحه پشتیبانی",
		"it-IT": "Pagina d'aiuto",
		"ru-RU": "Cтраница поддержки ",
	},
	COMMAND_TEXT_REPORT_A_BUG: {
		"de-DE": "einen Fehler melden",
		"en-UK": "Report a bug",
		"es-ES": "Informar de un error",
		"fa-IR": "گزارش یک باگ",
		"it-IT": "Segnala un bug",
		"ru-RU": "Сообщить об ошибке",
	},
	COMMAND_TEXT_SUBMIT_AN_IDEA: {
		"de-DE": "eine Idee äußern",
		"en-UK": "Add an idea",
		"es-ES": "Proponer una idea",
		"fa-IR": "یک ایده اضافه کنید.",
		"it-IT": "Proponi un idea",
		"ru-RU": "Предложить идею",
	},
	MESSAGE_TEXT_WELCOME: {
		"de-DE": `Hallo, ich bin Collectius - dein persönlicher Buchhalter.

Ich kann dir sagen, wer wem was schuldet und wann die Schuld jeweils fällig ist.

Was würdest du gerne machen?`,
		"en-UK": `Hi, I'm Collectius - your personal accountant & collector.

I can record who is owing to whom and remind when the return is due.

What would you like to do?`,

		"es-ES": `Hola, me llamo Collectius, soy tu contable y asesor personal.

Puedo apuntar quién debe a quién y recordarte la fecha de devolución.

¿Qué te apetecería hacer?`,

		"fa-IR": `سلام، من کالکتیوس هستم - حسابدار شخصی و مامور وصول شما

من میتوانم اینکه چه کسی به چه کسی بدهکار است را ثبت کرده و زمان بازپرداخت را یادآوری کنم.

دوست دارید چکار کنید؟`,

		"it-IT": `Ciao, sono Collectius - il tuo contabile ed esattore.

Posso annotare chi deve soldi a chi e ricordarti la data di scadenza.

Cosa vorresti fare ora?`,

		"ru-RU": `Привет, я Коллектиус - Ваш персональный счетовод и коллектор.

Могу записать кто кому чего должен и, и при необходимости, напомнить когда должок пора возвращать.

Чего изволит новый хозяин?`,
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE: {
		"de-DE": "Ich hätte gerne einen Code",
		"en-UK": "I want to get an invite",
		"es-ES": "Me gustaría obtener una invitación",
		"fa-IR": "می خواهم یک دعوت دریافت کنم.",
		"it-IT": "Voglio ottenere un invito",
		"ru-RU": "Хочу получить приглашение",
	},
	COMMAND_TEXT_I_HAVE_INVITE: {
		"de-DE": "Ich habe einen Code",
		"en-UK": "I have the invitation code",
		"es-ES": "Tengo el código de la invitación",
		"fa-IR": "من کد دعوت را دارم.",
		"it-IT": "Ho il codice invito",
		"ru-RU": "У меня есть код приглашения",
	},
	COMMAND_TEXT_I_HAVE_NOT_GOT_EMAIL: {
		"de-DE": "Ich habe noch keine Mails bekommen",
		"en-UK": "I have not got any emails",
		"es-ES": "No he recibido ningún e-mail",
		"fa-IR": "من ایمیلی دریافت نکردم",
		"it-IT": "Non ho ricevuto nessun email",
		"ru-RU": "Я не получил письма на email",
	},
	MESSAGE_TEXT_ONBOARDING_TELL_ABOUT_INVITES: {

		"de-DE": `<b>%v</b>,

Im Moment ist der Bot leider nur durch Einladungen von Freunden zugänglich.

Wenn du keinen Code hast, lass deine Kontaktdaten da und wir senden dir einen Code sobald du dran bist.

Wir senden 10 Codes am Tag an die, die am längsten warten und einen zufällig.`,

		"en-UK": `<b>%v</b>,

At the moment our bot is available just by invitation from friends.

If you have no code you can leave your contact details and we'll send you an invite as soon as your queue is due.

We send 10 invites per day to those in the head of the queue and 1 randomly.`,

		"es-ES": `<b>%v</b>,

De momento nuestro bot está disponible solo por invitación de amigos.

Si no tienes el código puedes dejarnos tu contacto y te lo enviaremos cuando sea tu turno en la cola .

Enviamos 10 invitaciones por día a los primeros de la cola + 1 de modo casual.`,

		"fa-IR": `<b>%v</b>,

درحال حاضر ربات ما تنها با دریافت دعوت از دوستان در دسترس می باشد.

اگر شما کدی در اختیار ندارید می توانید اطلاعات تماس خود را برای من وارد نموده و من به محض اینکه نوبت شما فرارسید یک دعوتنامه برایتان ارسال می کنم.

ما روزانه 10 دعوتنامه برای نفرات اول لیست انتظار و همچنین یک دعوتنامه تصادفی ارسال میکنیم.`,

		"it-IT": `<b>%v</b>,

Al momento il nostro bot e' disponibile solo tramite invito da amici.

Se non hai un codice puoi lasciarci il tuo contatto e ti manderemo un invito non appena sara' il tuo turno.

Inviamo 10 inviti al giorno ai primi 10 della lista d'attesa ed 1 in modo casuale.`,

		"ru-RU": `<b>%v</b>,

	На данный момент наш бот доступен только тем кто получил приглашение от друзей.

	Если у вас нет кода можете оставить свои контакты и мы вышлем вам приглашение как только наступит ваша очередь.

	Мы высылаем по 10 кодов в день первоочередникам + 1 случайным образом.`,
	},
	EMAIL_INVITE_SUBJ: {
		"de-DE": "Eine Einladung von {{.FromName}} - Code: {{.InviteCode}}",
		"en-UK": "An invite from {{.FromName}} - code: {{.InviteCode}}",
		"es-ES": "La invitación de {{.FromName}} - el código: {{.InviteCode}}",
		"fa-IR": "دعوت از طرف {{.FromName}} - کد: {{.InviteCode}}",
		"it-IT": "Hai ricevuto un codice invito da {{.FromName}} - codice: {{.InviteCode}}",
		"ru-RU": "Приглашение от {{.FromName}} - код: {{.InviteCode}}",
	},
	SMS_INVITE_TEXT: {

		"de-DE": `Hey {{.ToName}}, {{.FromName}} lädt dich ein die neue Schuldentracker App auszuprobieren - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Dein persönlicher Code lautet: {{.InviteCode}}`,

		"en-UK": `Hi {{.ToName}}, {{.FromName}} is inviting you to try debts tracking app - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Your personal invitation code is: {{.InviteCode}}`,

		"es-ES": `Hola {{.ToName}}, {{.FromName}} te ha invitado a probar la aplicación para controlar tus deudas - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

El código de tu invitación es: {{.InviteCode}}`,

		"fa-IR": `سلام{{.ToName}}, {{.FromName}} شما را دعوت کرده تا برنامه ردیابی بدهی ها را امتحان کنید.- https://debtstracker.io/invite#id={{.InviteCode}}&ربات تلگرام={{.TgBot}}&{{.Utm}}

کد دعوت شما: {{.InviteCode}}`,

		"it-IT": `Ciao {{.ToName}}, {{.FromName}} ti ha invitato a provare 'debts tracking app' - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Il tuo codice di invito personale e': {{.InviteCode}}`,

		"ru-RU": `Привет {{.ToName}}, {{.FromName}} рекомендует приложение для учёта долгов - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Код приглашения: {{.InviteCode}}`,
	},
	EMAIL_INVITE_TEXT: {
		"de-DE": `Hey {{.ToName}}, 
		
{{.FromName}} lädt dich ein die neue Schuldentracker App auszuprobieren - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Dein persönlicher Code lautet: {{.InviteCode}}`,

		"en-UK": `Hi {{.ToName}},

{{.FromName}} is inviting you to use debts tracking app - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

You invitation code is: {{.InviteCode}}`,

		"es-ES": `Hola {{.ToName}},

{{.FromName}} te ha invitado a probar la aplicación para controlar tus deudas - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

El código de tu invitación es: {{.InviteCode}}`,

		"fa-IR": `سلام{{.ToName}},

{{.FromName}} شما را دعوت کرده تا از برنامه ردیابی بدهی ها استفاده کنید.- https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

کد دعوت شما: {{.InviteCode}}`,

		"it-IT": `Ciao {{.ToName}},

{{.FromName}} ti ha invitato a provare 'debts tracking app' - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Il tuo codice di invito personale e': {{.InviteCode}}`,

		"ru-RU": `Привет {{.ToName}},

	{{.FromName}} приглашает тебя попробовать приложение для учёта долгов - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

	Ваш код приглашения: {{.InviteCode}}`,
	},
	EMAIL_INVITE_HTML: {
		"de-DE": `<p>Hey {{.ToName}}, </p>
		
<p>{{.FromName}} lädt dich ein <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">die neue Schuldentracker App auszuprobieren</a>.</p>

<p>Dein persönlicher Code lautet: <b>{{.InviteCode}}</b></p>`,
		"en-UK": `<p>Hi {{.ToName}}, </p>

<p>{{.FromName}} is inviting you to try <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">debts tracking app</a>.</p>

<p>Your invitation code is: <b>{{.InviteCode}}</b></p>`,

		"es-ES": `<p>Hola {{.ToName}}, </p>

<p>{{.FromName}} te ha invitado a <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">probar la app para controlar tus deudas</a>.</p>

<p>El código de tu invitación es: <b>{{.InviteCode}}</b></p>`,

		"fa-IR": `<p>سلام{{.ToName}},</p>

<p>{{.FromName}} п شما را دعوت کرده به <a href="https://debtstracker.io/invite#id={{.InviteCode}}&ربات تلگرام={{.TgBot}}&{{.Utm}}"> امتحان برنامه ردیابی بدهی ها</a>.</p>

<p>You invitation code is: <b>{{.InviteCode}}</b></p>`,

		"it-IT": `<p>Ciao {{.ToName}},</p>

<p>{{.FromName}} ti ha invitato a provare <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">debts tracking app</a>.</p>

<p>Il tuo codice di invito personale e': <b>{{.InviteCode}}</b></p>`,

		"ru-RU": `<p>Привет {{.ToName}}, </P

	<p>{{.FromName}} приглашает тебя <a href = "https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">попробовать приложение для учёта долгов</a>.</p>

	<p>Ваш код приглашения: <b>{{.InviteCode}}</b></p>`,
	},
	EMAIL_RECEIPT_SUBJ: {
		"de-DE": "Schuldschein - {{.FromName}}",
		"en-UK": "Debt record - {{.FromName}}",
		"es-ES": "La notificación de la deuda - {{.FromName}}",
		"fa-IR": "سوابق بدهی - {{.FromName}}",
		"it-IT": "Debito - {{.FromName}}",
		"ru-RU": "Запись о долге - {{.FromName}}",
	},
	EMAIL_RECEIPT_BODY_TEXT: {
		"de-DE": "{{.FromName}} hat einen Schuldschein erstellt: {{.ReceiptURL}}",
		"en-UK": "{{.FromName}} created a debt record: {{.ReceiptURL}}",
		"es-ES": "{{.FromName}} ha creado una notificación de la deuda: {{.ReceiptURL}}",
		"fa-IR": "{{.FromName}} یک سابقه بدهی ایجاد کرده است: {{.ReceiptURL}}",
		"it-IT": "{{.FromName}} ha creato un debito: {{.ReceiptURL}} ",
		"ru-RU": "{{.FromName}} создал(а) запись о долге: {{.ReceiptURL}}",
	},
	MESSENGER_RECEIPT_TEXT: {
		"de-DE": "Ich habe online einen Schuldschein erstellt, für mehr Details siehe {{.ReceiptURL}}",
		"en-UK": "I've created a debt record regards you, see details at {{.ReceiptURL}}",
		"es-ES": "He creado una notificación de la deuda, las detalles están aquí {{.ReceiptURL}}",
		"fa-IR": "من یک سابقه بدهی برای شما ایجاد کرده ام، لطفا جزئیات را ملاحظه فرمایید در {{.ReceiptURL}}",
		"it-IT": "Ho creato un debito a tuo nome, controlla i dettagli qui - {{.ReceiptURL}}",
		"ru-RU": "Я создал(а) запись о долге, подробности тут - {{.ReceiptURL}}",
	},
	EMAIL_RECEIPT_BODY_HTML: {
		"de-DE": "{{.FromName}} hat einen Schuldschein erstellt: {{.ReceiptURL}}",
		"en-UK": "{{.FromName}} created a debt record: {{.ReceiptURL}}",
		"es-ES": "{{.FromName}} ha creado una notificación de la deuda: {{.ReceiptURL}}",
		"fa-IR": "{{.FromName}} یک سابقه بدهی ایجاد کرده است: {{.ReceiptURL}}",
		"it-IT": "{{.FromName}} ha creato un debito: {{.ReceiptURL}}",
		"ru-RU": "{{.FromName}} создал(а) запись о долге: {{.ReceiptURL}}",
	},
	INLINE_RECEIPT_TITLE: {
		"de-DE": "Empfänger: %v",
		"en-UK": "Receipt: %v",
		"es-ES": "El recibo: %v",
		"fa-IR": "رسید: %v",
		"it-IT": "Debito/credito: %v",
		"ru-RU": "Квитанция: %v",
	},
	INLINE_RECEIPT_DESCRIPTION: {
		"de-DE": "Klicken sie hier, um die Quittung zu sehen",
		"en-UK": "Click here to send the receipt",
		"es-ES": "Haz click aquí para enviar el recibo",
		"fa-IR": "برای ارسال رسید اینجا کلیک کنید.",
		"it-IT": "Clicca qui per inviare un debito/credito",
		"ru-RU": "Нажмите здесь чтобы отправить квитанцию",
	},
	INLINE_RECEIPT_CHOOSE_LANGUAGE: {
		"de-DE": "<b>Bitte wählen Sie eine Sprache, um weitere Details zu sehen, die den Schuldschein betreffen</b>, der erstellt wurde von {{.Creator}}.",
		"en-UK": "<b>Please choose language to see details of the debt</b> that has been recorded by {{.Creator}}.",
		"es-ES": "<b>Elige el idioma para ver los detalles de la deuda</b> que ha sido creada por {{.Creator}}.",
		"fa-IR": "<b> لطفا برای رویت جزئیات بدهی که توسط </b>  {{.Creator}} ثبت شده است زبان را انتخاب کنید.",
		"it-IT": "<b>Scegli la lingua per vedere i dettagli del debito</b> registrato da {{.Creator}}.",
		"ru-RU": "<b>Выберите язык чтобы посмотреть подробности записи о долге</b> которую создал(а) {{.Creator}}.",
	},
	INLINE_RECEIPT_FOOTER: {
		//-------------------------------------------------------
		"de-DE": `{{.SiteLink}} — eine App, die dir hilft Schulden zu überwachen:

  - Du weißt immer, wie viel du allen schuldest

  - Keine Fälligkeit wird verpasst
    <i>(erinnert dich und die Gläubiger)</i>`,
		//-------------------------------------------------------
		"en-UK": `{{.SiteLink}} — an app for debts tracking will help you to:

  - Always know your bottom line

  - Return debts on time
    <i>(reminders to you & your debtors)</i>`,
		//-------------------------------------------------------
		"es-ES": `{{.SiteLink}} — la app para controlar tus deudas te ayuda a:

  - Saber siempre quién debe a quién

  - Devolver la deuda a tiempo
    <i>(recordatorio a ti y a tus deudores)</i>`,
		//-------------------------------------------------------
		"fa-IR": `{{.SiteLink}} — یک برنامه پیگیری بدهی است که به شما کمک می کند تا:

  - همیشه از سود و زیان خود مطلع باشید.

  - بدهی ها به موقع پرداخت شوند.
    <i>(با ارسال یادآوری به  شما و بدهکاران )</i>`,
		//-------------------------------------------------------
		"it-IT": `{{.SiteLink}} — un app per i debiti che ti consento di:

  - Sapere sempre chi deve soldi a chi

  - Restituire i soldi in tempo
    <i>(lo ricorda a te ed al tuo debitore)</i>`,
		//-------------------------------------------------------
		"ru-RU": `{{.SiteLink}} — программа для учёта долгов поможет:

	  - Всегда знать кто кому сколько должен

	  - Незабыть вовремя отдать или востребовать долг
	    <i>(напоминания вам и вашим должникам)</i>`,
		//-------------------------------------------------------
	},
	INLINE_RECEIPT_GENERATING_MESSAGE: {
		"de-DE": `<b>{{.Creator}} erstellte online einen Schuldschein</b> der dich betrifft.

>> Generating receipt`, // TODO(DE)
		//-------------------------------------------------------
		"en-UK": `<b>{{.Creator}} recorded a debt</b> associated with you.

  >> <i>Generating receipt...</i>`,
		//-------------------------------------------------------
		"es-ES": `<b>{{.Creator}} ha creado una deuda</b> asociada a ti.

  >> <i>Generating receipt...</i>`, // TODO(ES)
		//-------------------------------------------------------
		"fa-IR": `<b>{{.Creator}} یک بدهی </b> مرتبط با شما ثبت نموده است.

  >> <i>Generating receipt...</i>`, // TODO(FA)
		//-------------------------------------------------------
		"it-IT": `<b>{{.Creator}} ha registrato un debito</b> associato a te.

  >> <i>Generating receipt...</i>`, // TODO(IT)
		//-------------------------------------------------------
		"ru-RU": `<b>{{.Creator}} создал(а) запись о долге</b> касающегося Вас.

  >> <i>Generating receipt...</i>`,
		//-------------------------------------------------------
	},
	//	INLINE_RECEIPT_MESSAGE: {
	//		//-------------------------------------------------------
	//		"en-UK": `<b>{{.Creator}} recorded a debt</b> associated with you.
	//
	//`,
	//		//-------------------------------------------------------
	//		"es-ES": `<b>{{.Creator}} ha creado una deuda</b> asociada a ti.
	//
	//`,
	//		//-------------------------------------------------------
	//		"fa-IR": `<b>{{.Creator}} یک بدهی </b> مرتبط با شما ثبت نموده است.
	//
	//`,
	//		//-------------------------------------------------------
	//		"it-IT": `<b>{{.Creator}} ha registrato un debito</b> associato a te.
	//
	//`,
	//		//-------------------------------------------------------
	//		"ru-RU": `<b>{{.Creator}} создал(а) запись о долге</b> касающегося Вас.
	//
	//`,
	//		//-------------------------------------------------------
	//	},
	INLINE_RECEIPT_MESSAGE: {
		//-------------------------------------------------------
		"en-UK": `<b>{{.Creator}} recorded a debt</b> associated with you.

  >> <a href="{{.ReceiptUrl}}">Click here to view receipt</a>`,
		//-------------------------------------------------------
		"es-ES": `<b>{{.Creator}} ha creado una deuda</b> asociada a ti.

  >> <a href="{{.ReceiptUrl}}">Click here to view receipt</a>`,
		//-------------------------------------------------------
		"fa-IR": `<b>{{.Creator}} یک بدهی </b> مرتبط با شما ثبت نموده است.

  >> <a href="{{.ReceiptUrl}}">Click here to view receipt</a>`,
		//-------------------------------------------------------
		"it-IT": `<b>{{.Creator}} ha registrato un debito</b> associato a te.

  >> <a href="{{.ReceiptUrl}}">Click here to view receipt</a>`,
		//-------------------------------------------------------
		"ru-RU": `<b>{{.Creator}} создал(а) запись о долге</b> касающегося Вас.

  >> <a href="{{.ReceiptUrl}}">Посмотреть квитанцию</a>`,
		//-------------------------------------------------------
	},
	INLINE_INVITE_TITLE: {
		"de-DE": "Einladung zu %v",
		"en-UK": "Invitation to %v",
		"es-ES": "Invitación a %v",
		"fa-IR": "دعوت به %v",
		"it-IT": "Invito per %v",
		"ru-RU": "Приглашение в %v",
	},
	INLINE_INVITE_DESCRIPTION: {
		"de-DE": "Klick hier, um eine Einladung zu versenden",
		"en-UK": "Click here to send an invite",
		"es-ES": "Haz click para enviar la invitación",
		"fa-IR": "برای ارسال یک دعوتنامه اینجا کلیک کنید.",
		"it-IT": "Clicca qui per spedire un invito",
		"ru-RU": "Нажмите здесь для отправки приглашения",
	},
	INLINE_INVITE_MESSAGE: {
		"de-DE": "%v hat dich eingeladen %v auszuprobieren",
		"en-UK": "%v invited you to try %v",
		"es-ES": "%v te ha invitado a probar %v",
		"fa-IR": "%v شمارا دعوت کرده است به امتحان %v",
		"it-IT": "%v ti ha invitato a provare %v",
		"ru-RU": "%v пригласил вас попробовать %v",
	},
	SMS_RECEIPT_YOU_GOT: {
		"de-DE": "Du hast dir %v von %v geliehen.",
		"en-UK": "You've got %v from %v.",
		"es-ES": "Has recibido %v de %v.",
		"fa-IR": "شما دریافت کرده اید %v از %v.",
		"it-IT": "Hai ricevuto %v da %v",
		"ru-RU": "Вы получили %v от %v.",
	},
	SMS_RECEIPT_YOU_GAVE: {
		"de-DE": "Du hast %v an %v verliehen.",
		"en-UK": "You've given %v to %v.",
		"es-ES": "Has dado %v a %v.",
		"fa-IR": "شما پرداخت کرده اید %v به %v.",
		"it-IT": "Hai dato %v a %v",
		"ru-RU": "Вы дали %v - взял(а) %v.",
	},
	SMS_CLICK_TO_CONFIRM_OR_DECLINE: {
		"de-DE": "Klicke %v um zu akzeptieren oder abzulehnen.",
		"en-UK": "Click %v to confirm or decline.",
		"es-ES": "Haz click %v para confirmar o rechazar.",
		"fa-IR": "کلیک کنید %v تا رد خود را تایید نمایید",
		"it-IT": "Clicca %v per confermare o rifiutare",
		"ru-RU": "Нажмите %v чтобы подтвердить или отвергнуть.",
	},
	HTML_DATE: {
		"de-DE": "Datum",
		"en-UK": "Date",
		"es-ES": "Fecha",
		"fa-IR": "تاریخ",
		"it-IT": "Data",
		"ru-RU": "Дата",
	},
	HTML_RECEIPT: {
		"de-DE": "Empfänger",
		"en-UK": "Receipt",
		"es-ES": "Recibo",
		"fa-IR": "رسید",
		"it-IT": "Scontrino", //To upgrade, not the best translation from Russian
		"ru-RU": "Квитанция",
	},
	HTML_AMOUNT: {
		"de-DE": "Betrag",
		"en-UK": "Amount",
		"es-ES": "Importe",
		"fa-IR": "مقدار",
		"it-IT": "Totale",
		"ru-RU": "Сумма",
	},
	HTML_FROM: {
		"de-DE": "Von",
		"en-UK": "From",
		"es-ES": "De",
		"fa-IR": "از",
		"it-IT": "Da",
		"ru-RU": "Дал",
	},
	HTML_TO: {
		"de-DE": "An",
		"en-UK": "To",
		"es-ES": "Para",
		"fa-IR": "به",
		"it-IT": "Per",
		"ru-RU": "Получил",
	},
	NO_NAME: {
		"de-DE": "unbekannt",
		"en-UK": "no name",
		"es-ES": "sin nombre", // TODO(es) verify
		"fa-IR": "بدون نام",   // TODO(fa) verify
		"it-IT": "senza nome", // TODO(it) verify
		"ru-RU": "без имени",
	},
	TELEGRAM_RECEIPT: {
		"de-DE": "{{.FromName}} hat einen Schuldschein erstellt ({{.TransferCurrency}})",
		"en-UK": "{{.FromName}} created a debt record ({{.TransferCurrency}})",
		"es-ES": "{{.FromName}} ha creado una deuda ({{.TransferCurrency}})",
		"fa-IR": "{{.FromName}} ایجاد یک سابقه بدهی ({{.TransferCurrency}})",
		"it-IT": "{{.FromName}} ha registrato un debito ({{.TransferCurrency}})",
		"ru-RU": "{{.FromName}} создал запись о долге ({{.TransferCurrency}})",
	},
	MESSAGE_TEXT_PLEASE_CHOOSE_FROM_OPTIONS_PROVIDED: {
		"de-DE": "Bitte wähle aus den angezeigten Optionen.",
		"en-UK": "Please choose from provided options.",
		"es-ES": "Por favor, elige una de las siguientes opciones.",
		"fa-IR": "لطفاً از گزینه های ارائه شده انتخاب نمایید.",
		"it-IT": "Scegli tra le opzioni fornite.",
		"ru-RU": "Пожалуйста выберете из предоставленных опций.",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_NOTE_OR_COMMENT: {
		"de-DE": "<b>Möchtest du eine Bemerkung oder Notiz hinzufügen?</b>\n%v Deine Notizen kannst nur du sehen.\n%v Eine Bemerkung wird quasi auf dem Schuldschein und der Quittung vermerkt und ist insofern für beide sichtbar.",
		"en-UK": "<b>Do you want to add a note or comment?</b>\n%v Memos are private records for yoru own reference.\n%v Comments are available to everyone who has permission to view this transaction.",
		"es-ES": "<b>¿Quieres añadir una nota o comentario?</b>\n%v Las notas se graban de manera privada para tu propia información.\n%v Los comentarios son visibles para todos los autorizados a ver esta transacción.",
		"fa-IR": "<b>آیا میخواهید یادداشت یا شرحی اضافه کنید؟</b>\n%v یادداشت ها نوشته های خصوصی برای مراجعه خود شما هستند.\n%v شرح در دسترس تمام کسانی که مجاز رویت این تراکنش هستند میباشد.",
		"it-IT": "<b>Vuoi aggiungere una nota o un commento?</b> \n%v I memo sono record privati per il riferimento di yoru.\n%v I commenti sono disponibili a tutti coloro che hanno l'autorizzazione a visualizzare questa transazione.",
		"ru-RU": "<b>Хотите добавить заметку или комментарий?</b>\n%v Заметки хранятся для вашего личного пользования.\n%v Комментарий виден всем кому разрешён просмотр этой транзакции.",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_NOTE: {
		"de-DE": "Bitte schreibe eine Notiz:",
		"en-UK": "Please write a note:",
		"es-ES": "Por favor, escribe una nota:",
		"fa-IR": "لطفاً یک یادداشت بنویسید:",
		"it-IT": "Per favore scrivi un appunto:",
		"ru-RU": "Напишите заметку:",
	},
	COMMAND_TEXT_MORE_ABOUT_INTEREST_COMMAND: {
		"de-DE": "Mehr über Prozentsätze", // TODO(DE)
		"en-UK": "More about interest",
		"es-ES": "Más sobre interés",   // TODO(ES)
		"fa-IR": "بیشتر در مورد علاقه", // TODO(FA)
		"it-IT": "Dimmi di più",        // TODO(IT)
		"ru-RU": "Подробнее о процентах",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_INTEREST_SHORT: {
		"de-DE": `<b>Prozent und Kommentar</b>

To set interest rate & period send a message in following format:
		<pre>percent/percent_period/min_period/grace_period:note</pre>`, // TODO(DE)
		"en-UK": `<b>Interest & notes</b>

To set interest rate & period send a message in following format:
		<pre>percent/percent_period/min_period/grace_period:note</pre>`,
		"es-ES": `<b>Porcentaje y comentario</b>

To set interest rate & period send a message in following format:
		<pre>percent/percent_period/min_period/grace_period:note</pre>`, // TODO(ES)
		"fa-IR": `<b>درصد و نظر</b>

To set interest rate & period send a message in following format:
		<pre>percent/percent_period/min_period/grace_period:note</pre>`, // TODO(FA)
		"it-IT": `<b>Percentuale e commento</b>

To set interest rate & period send a message in following format:
		<pre>percent/percent_period/min_period/grace_period:note</pre>`, // TODO(IT
		"ru-RU": `<b>Процент и комментарий</b>

Чтобы задать процент по долгу отправьте сообщение в следующем формате:
	<pre>процент/процентный_период/минимальный_период/грэйс_период:комментарий</pre>`,
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_INTEREST_LONG: {
		"de-DE": ``, // TODO(DE)
		"en-UK": `<b>Interest & notes</b>

To set interest rate & period send a message in following format:

		<pre>percent/percent_period/min_period/grace_period:note</pre>

Where (<i>first 2 params are required</i>):
 * <code>percent</code> - up to 2 digits after comma.
 * <code>min_period</code> - number of days for interest period.
 * <code>min_perdio</code> - minimum number of days for interest calculation. Is 1 by default and can't be less.'.
 * <code>grace_period</code> - interest-less period. At the moment you can't set grace & min period at the same time'.
 * <code>note</code> - any explanatory text that will be visible to you and your counterparty.

Interest is calculated daily (<i>every 24 hours</i>) using <a href="https://www.investopedia.com/terms/s/simple_interest.asp#utm_source=DebtsTrackerBot&utm_medium=telegram&utm_campaign=new_debt_wizard">simple percent</a> formula.

<b>Examples</b>:

		<code>2/7/5</code> - 2% per week, minimum for 5 days
		<code>15/360</code> - 15%/year, (<i>minimum for 1 day</i>)
		<code>3/30/0/10</code> - 3% per month with 10 days grace period

❗ The % functionality is in BETA testing stage, please let us know in @DebtsTrackerGroup if anything works not as you would expect.`, // TODO - replace link!
		"es-ES": ``, // TODO(ES)
		"fa-IR": ``, // TODO(FA)
		"it-IT": ``, // TODO(IT
		"ru-RU": `<b>Процент и комментарий</b>

Чтобы задать процент по долгу отправьте сообщение в следующем формате:

	<pre>процент/процентный_период/минимальный_период/грэйс_период:комментарий</pre>

Где (<i>первые 2 параметра обязательны</i>):
 * <code>процент</code> - возможно до 2х знаков после запятой.
 * <code>процентный_период</code> - количество дней за которые начисляется процент указанный предыдущим числом.
 * <code>минимальный_период</code> - минимальное количество дней за которые сразу начисляются проценты. По умолчанию 1 и не может быть меньше.
 * <code>грэйс_период</code> - безпроцентный период. Пока нельзя задавать и минимальный процент и безпроцентный период одновременно.
 * <code>комментарий</code> - любой поясняющий текст видимый вам и другому участнику сделки.

Начисление ежедневно (<i>каждые 24 часа</i>) по формуле <a href="https://ru.wikipedia.org/wiki/%D0%9F%D1%80%D0%BE%D1%81%D1%82%D1%8B%D0%B5_%D0%BF%D1%80%D0%BE%D1%86%D0%B5%D0%BD%D1%82%D1%8B">простого процента</a>.

<b>Примеры</b>:

		<code>2/7/5</code> - под 2% в неделю, минимум на 5 дней
		<code>15/360</code> - под 15% годовых
		<code>3/30/0/10</code> - под 3% в месяц с безпроцентным периодом в 10 дней

❗ Функционал % ещё тестируется, пожалуйста сообщите в @DebtsTrackerGroup если что-то пошло не так.`,
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_COMMENT_ONLY: {
		"de-DE": "Falls du eine Bemerkung auf den Schuldschein schreiben willst, schick mir jetzt den Text.",
		"en-UK": `If you want to add a comment just send a text now.`,
		"es-ES": `si quieres añadir un comentario simplemente envia un texto.`,
		"fa-IR": `شما می توانید یک شرح اضافه کنید. تنها کافیست یک متن ارسال کنید.`,
		"it-IT": `Se vuoi aggiungere un commento invia del testo ora.`,
		"ru-RU": `Если хотите добавить комментарий просто отправьте текст.`,
	},
	MESSAGE_TEXT_VISIBLE_TO_YOU_AND_COUNTERPARTY: {
		"de-DE": "sichtbar für dich & %v",
		"en-UK": "visible to you & %v",
		"es-ES": "visible solo para ti & %v",
		"fa-IR": "قابل مشاهده برای شما & %v",
		"it-IT": "visibile solo a te e %v",
		"ru-RU": "виден вам и %v",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_COMMENT: {
		"de-DE": "Schreibe nun eine Bemerkung auf den Schuldschein:",
		"en-UK": "Please write the comment:",
		"es-ES": "Por favor, escribe un comentario:",
		"fa-IR": "لطفاً شرح را ثبت کنید:",
		"it-IT": "Per favore scrivi un commento:",
		"ru-RU": "Напишите комментарий:",
	},
	MESSAGE_TEXT_TRANSFER_NOTE_ADDED_ASK_FOR_COMMENT: {
		"de-DE": "Deine Notiz wurde hinzugefügt, möchtest du noch eine Bemerkung auf den Schuldschein schreiben?",
		"en-UK": "Memo have been added. Do you want to write a comment?",
		"es-ES": "La nota está añadida. ¿Quieres escribir un comentario?",
		"fa-IR": "یادداشت اضافه شد. آیا میخواهید یک شرح ثبت کنید؟",
		"it-IT": "Promemoria aggiunto. Vuoi scrivere un commento?",
		"ru-RU": "Заметка добавлена. Хотите написать комментарий?",
	},
	MESSAGE_TEXT_TRANSFER_COMMENT_ADDED_ASK_FOR_NOTE: {
		"de-DE": "Die Bemerkung wurde vermerkt. Möchtest du noch eine Notiz für dich hinzufügen?",
		"en-UK": "Comment have been added. Do you want to write a note?",
		"es-ES": "El comentario está añadido. ¿Quieres escribir una nota?",
		"fa-IR": "شرح موردنظر شما ثبت شد. آیا می خواهید یک یادداشت بنویسید؟",
		"it-IT": "Commento aggiunto. Vuoi scrivere un appunto?",
		"ru-RU": "Комментарий добавлен. Хотите написать заметку?",
	},
	COMMAND_TEXT_NO_COMMENT_OR_NOTE_FOR_TRANSFER: {
		"de-DE": "Keine Notizen oder Bemerkungen",
		"en-UK": "Without notes or comments",
		"es-ES": "Sin notas ni comentarios",
		"fa-IR": "بدون یادداشت یا شرح",
		"it-IT": "Senza appunti o commenti",
		"ru-RU": "Без заметок и комментариев",
	},
	COMMAND_TEXT_NO_COMMENT_FOR_TRANSFER: {
		"de-DE": "Keine Bemerkungen",
		"en-UK": "No comments",
		"es-ES": "Sin comentarios",
		"fa-IR": "بدون شرح",
		"it-IT": "Nessun commento",
		"ru-RU": "Без комментариев",
	},
	COMMAND_TEXT_NO_NOTE_FOR_TRANSFER: {
		"de-DE": "Keine Notizen",
		"en-UK": "Without notes",
		"es-ES": "Sin notas",
		"fa-IR": "بدون یادداشت",
		"it-IT": "Senza appunti/note",
		"ru-RU": "Без заметок",
	},
	COMMAND_TEXT_ADD_NOTE_TO_TRANSFER: {
		"de-DE": "Notiz hinzufügen (privat)",
		"en-UK": "Add a note (private)",
		"es-ES": "Añadir una nota (privada)",
		"fa-IR": "یک یادداشت اضافه کنید (خصوصی)",
		"it-IT": "Aggiungi una nota (privata)",
		"ru-RU": "Добавить заметку",
	},
	COMMAND_TEXT_ADD_COMMENT_TO_TRANSFER: {
		"de-DE": "Bemerkung hinzufügen (öffentlich)",
		"en-UK": "Add a comment (public)",
		"es-ES": "Añadir un comentario (público)",
		"fa-IR": "یک شرح اضافه کنید (عمومی)",
		"it-IT": "Aggiungi un commento (pubblico)",
		"ru-RU": "Добавить комментарий",
	},
	DUE_IN_NOW: {
		"de-DE": "jetzt",
		"en-UK": "now",
		"es-ES": "ahora",
		"fa-IR": "حالا",
		"it-IT": "adesso",
		"ru-RU": "прямо сейчас",
	},
	DUE_IN_A_MINUTE: {
		"de-DE": "in einer Minute",
		"en-UK": "in a minute",
		"es-ES": "en un minuto",
		"fa-IR": "ظرف یک دقیقه",
		"it-IT": "in un minuto",
		"ru-RU": "через минуту",
	},
	DUE_IN_X_MINUTES: {
		"de-DE": "in %v Minuten",
		"en-UK": "in %v minutes",
		"es-ES": "en %v minutos",
		"fa-IR": "در %v دقیقه",
		"it-IT": "in %v minuti/o",
		"ru-RU": "через %v минут(ы)",
	},
	DUE_IN_AN_HOUR: {
		"de-DE": "in einer Stunde",
		"en-UK": "in an hour",
		"es-ES": "en  una hora",
		"fa-IR": "ظرف یک ساعت",
		"it-IT": "in un ora",
		"ru-RU": "через час",
	},
	DUE_IN_X_HOURS: {
		"de-DE": "in %v Stunde",
		"en-UK": "in %v hours",
		"es-ES": "en %v horas",
		"fa-IR": "در %v ساعت",
		"it-IT": "in %v ore/a",
		"ru-RU": "через %v часа/часов",
	},
	DUE_IN_X_DAYS: {
		"de-DE": "in %v Tagen",
		"en-UK": "in %v days",
		"es-ES": "en %v días",
		"fa-IR": "در %v روز",
		"it-IT": "in %v giorni/o",
		"ru-RU": "через %v дня/дней",
	},
	//-------------------------------------------------------------------------------------------------------------------
	WS_ALEX_T: {
		"de-DE": "Alexander Trakhimenok",
		"en-UK": "Alexander Trakhimenok",
		"es-ES": "Alexander Trakhimenok",
		"fa-IR": "الکساندر تراخیمِنوک",
		"it-IT": "Alexander Trakhimenok",
		"ru-RU": "Александр Трахимёнок",
	},

	WS_INDEX_TITLE: {
		"ru-RU": "DebtsTracker.io - программа для учёта личных долгов и активов",
		"en-UK": "DebtsTracker.io - an IOU app to track your personal debts & assets",
		"es-ES": "DebtsTracker.io es una aplicación para el control de sus deudas personales",
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
		"en-UK": "Demostración",
		"es-ES": "Demo en vivo",
		"fa-IR": "نسخه ی نمایشی زنده",
		"pl-PL": "Demo na żywo",
		"pt-PT": "Demonstração ao vivo",
		"de-DE": "Live-Demo",
		"fr-FR": "Démo en direct",
		"it-IT": "Demo online",
		"ko-KO": "실시간 데모",
		"ja-JP": "ライブデモ",
		"zh-CN": "现场演示",
	},
	WS_INDEX_TG_BOT_H2: {
		"ru-RU": "Бот для Telegram",
		"en-UK": "Chat bot for Telegram messenger",
		"es-ES": "Chat bot para Telegram",
		"fa-IR": "ربات چت برای پیام رسان تلگرام",
		"pl-PL": "Chat bot do telegramu posłańca",
		"pt-PT": "bot de bate-papo para Telegram messenger",
		"de-DE": "Chat-Bot für Telegram",
		"fr-FR": "bot Chat for Telegram messenger",
		"it-IT": "Bot Chat per Telegram",
		"ko-KO": "전보 메신저 채팅 봇",
		"ja-JP": "電報メッセンジャーのためのチャットボット",
		"zh-CN": "聊天机器人的电报使者",
	},
	WS_INDEX_TG_BOT_OPEN: {
		"ru-RU": "Открыть в Телеграмме &#x1F680;",
		"en-UK": "Open in Telegram &#x1F680;",
		"es-ES": "Abrir en Telegram &#x1F680;",
		"fa-IR": "بازکردن در تلگرام &#x1F680;",
		"pl-PL": "Otwórz w telegramu &#x1F680;",
		"pt-PT": "Open in Telegram &#x1F680;",
		"de-DE": "Öffnen in Telegram &#x1F680;",
		"fr-FR": "Open in Telegram &#x1F680;",
		"it-IT": "Apri su Telegram &#x1F680;",
		"ko-KO": "전보 &#x1F680; 에서 열기;",
		"ja-JP": "電報 &#x1F680; で開きます。",
		"zh-CN": "打开在电报 &#x1F680;",
	},

	WS_INDEX_TG_BOT_P: {
		"ru-RU": "В настоящий момент наша программа доступна в мессенджере <a href='https://telegram.org/'>Телеграм</a>.",
		"en-UK": "At the moment our program is available just on <a href='https://telegram.org/'>Telegram messenger</a>",
		"es-ES": "De momento nuestro programa está disponible sólo en <a href='https://telegram.org/'>Telegrama mensajero </a>",
		"fa-IR": "درحال حاضر برنامه ما فقط در دسترس است در <a href='https://telegram.org/'>Телеграм</a>تلگرام",
		"pl-PL": "W tej chwili nasz program jest dostępny tylko na <a href='https://telegram.org/'>Telegram messenger</a>",
		"pt-PT": "No momento em que o nosso programa está disponível apenas na <a href='https://telegram.org/'>Telegram messenger</a>",
		"de-DE": "Im Moment ist unser Programm nur auf <a href='https://telegram.org/'>Telegram verfügbar</a>",
		"fr-FR": "Au moment de notre programme est disponible seulement sur <a href='https://telegram.org/'>Telegram messager</a>",
		"it-IT": "Al momento il nostro programma è disponibile solo su <a href='https://telegram.org/'>Telegram</a>",
		"ko-KO": "지금이 순간 우리의 프로그램은 단지에 <a href='https://telegram.org/'>Telegram</a> 의 <b> 전보 </b>을 메신저 를 볼 수 있습니다",
		"ja-JP": "現時点では私たちのプログラムは、ちょうど上の<a href='https://telegram.org/'>Telegram</a>電報のメッセンジャーで提供されています",
		"zh-CN": "目前我们的计划是只提供在<a href='https://telegram.org/'>Telegram</a>电报的使者",
	},
	WS_MOTTO: {
		"ru-RU": "Платежи по долгам целиком и вовремя!",
		"en-UK": "Know your bottom line & never miss a debt payment!",
		"es-ES": "¡Controle sus pagos y deudas!",
		"fa-IR": "از سود و زیان خود مطلع باشید و هرگز پرداخت بدهی ای را از قلم نیندازید",
		"pl-PL": "Znaj swoją równowagę i nigdy nie przegapisz zapłatę długu!",
		"pt-PT": "Conheça o seu equilíbrio e nunca perca um pagamento da dívida!",
		"de-DE": "Wissen, wem man wie viel schuldet!",
		"fr-FR": "Apprenez à connaître votre solde et ne jamais manquer un paiement de la dette!",
		"it-IT": "Tieni sott'occhio il tuo bilancio e non dimenticarti mai di un debito!",
		"ko-KO": "균형을 알고 및 채무 지불을 놓칠 수 없어!",
		"ja-JP": "あなたのバランスを知っている＆債務の支払いを見逃すことはありません！",
		"zh-CN": "了解天平＆不会错过任何一个债务付款！",
	},
	WS_SHORT_DESC: {
		"ru-RU": "DebtsTracker.io - мобильное приложение и сервис напоминаний для учёта и своевременной выплаты долгов. Отсылает автоматические уведомления вашим должникам по email и SMS.",
		"en-UK": "DebtsTracker.io is a mobile IOU app & a reminder service that helps to track your debts, credits & assets. Sends automated email & SMS reminders to your debtors.",
		"es-ES": "DebtsTracker.io es un servicio de aplicaciones móviles y recordatorios que ayuda a realizar un seguimiento de sus deudas y créditos. Envía notificaciones por correo electrónico y SMS automatizados a sus deudores.",
		"fa-IR": "DebtsTracker.io یک برنامه موبایل و سرویس یادآور می باشد که به شما کمک می کند تا بدهی ها و اعتبارات خود را ردیابی نمایید. همچنین ایمیل و پیام کوتاه یادآوری به بدهکاران ارسال می کند.",
		"pl-PL": "DebtsTracker.io to aplikacje mobilne i przypomnienia usługa, która pozwala na śledzenie swoich długów i kredytów. Wysyła automatycznych powiadomień e-mail i SMS do swoich dłużników.",
		"pt-PT": "DebtsTracker.io é um serviço de aplicativos móveis e lembrete de que ajuda a controlar seus débitos e créditos. Envia e-mail e SMS notificações automáticas aos seus devedores.",
		"de-DE": "DebtsTracker.io ist eine mobile App, die beim Verwalten von persönlichen Schulden hilft - egal ob Sie Geld verleihen oder welches leihen. Sendet automatisierte E-Mail und SMS-Benachrichtigungen an Ihre Schuldner und Gläubiger.",
		"fr-FR": "DebtsTracker.io est une des applications mobiles et rappel service qui permet de suivre vos dettes et crédits. Envoie automatisés email & SMS reminders à vos débiteurs.",
		"it-IT": "DebtsTracker.io è un servizio di applicazioni mobili che ricordare e aiuta a monitorare i debiti e crediti. Invia notifiche e-mail e SMS automatici ai i vostri debitori.",
		"ko-KO": "DebtsTracker.io 은 채무 및 크레딧을 추적하는 데 도움이 모바일 앱 및 알림 서비스입니다. 당신의 채무자에 자동화 된 이메일 및 SMS 알림을 보냅니다.",
		"ja-JP": "DebtsTracker.io は、あなたの借金＆クレジットを追跡するのに役立ちますモバイルアプリ＆リマインダーサービスです。あなたの債務者に自動メール＆SMS通知を送信します。",
		"zh-CN": "DebtsTracker.io 是一个移动应用和提醒服务，帮助跟踪你的债务和信用。发送自动电子邮件和短信通知到您的债务人。",
	},

	WS_HELP_US_TITLE: {
		"de-DE": "Wie kann man beim DebtsTracker.io Projekt helfen kann",
		"en-UK": "How you can help to DebtsTracker.io project",
		"es-ES": "Como puedes ayudar a DebtsTracker.io project",
		"fa-IR": "چگونه می توانید به پروژه  DebtsTracker.io کمک کنید.",
		"it-IT": "Come potete aiutare il progetto DebtsTracker.io", // TODO(IT): Google translated
		"ru-RU": "Как вы можете помочь проекту DebtsTracker.io",
	},
	WS_ADS_TITLE: {
		"de-DE": "Werbung @ DebtsTracker.IO",
		"en-UK": "Ads @ DebtsTracker.IO",
		"es-ES": "Anuncio @ DebtsTracker.IO",
		"fa-IR": "تبلیغات @ DebtsTracker.IO",
		"it-IT": "Pubblicita' @ DebtsTracker.IO",
		"ru-RU": "Реклама на DebtsTracker.IO",
	},
	WS_ADS_CONTENT: {
		"de-DE": `Um Werbung in unserer App zu schalten, schick uns eine Mail an <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
		"en-UK": `To place ads in our app please send us an email to <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
		"es-ES": `Para publicar un anuncio en nuestra app escríbenos un e-mail a <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
		"fa-IR": `برای قراردادن تبلیغات در برنامه ما، درخواست خود را به این آدرس ایمیل فرمایید <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
		"it-IT": `Per inserire della pubblicita nella nostra app inviaci un email a <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
		"ru-RU": `Чтобы разместить рекламу в нашем приложении пишите нам на <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
	},
	FB_MAKE_RECORD_HEADER: {
		"de-DE": "Schuldschein",
		"en-UK": "Record debt",
		"es-ES": "Registrar la deuda",
		"fa-IR": "ثبت بدهی",
		"it-IT": "Registra il debito",
		"ru-RU": "Записать долг",
	},
	FB_MAKE_RECORD_HEADLINE: {
		"de-DE": "Scroll nach links, um weitere Optionen zu sehen.",
		"en-UK": "Scroll left to see other options.",
		"es-ES": "Desliza a la izquierda para ver otras opciones",
		"fa-IR": "برای دیدن سایر گزینه ها به سمت چپ اسکرول نمایید.",
		"it-IT": "Scrolla a sinistra per vedere altre opzioni",
		"ru-RU": "Пролистайте карточки влево чтобы увидеть другие опции.",
	},
	FB_HOW_ARE_THINGS_HEADER: {
		"de-DE": "Wie geht es dir?",
		"en-UK": "How are you doing?",
		"es-ES": "¿Cómo va todo?",
		"fa-IR": "حال شما چطوره؟",
		"it-IT": "Come te la passi?",
		"ru-RU": "Как идут дела?",
	},
	SNEATBOT_MSG_TXT_START: {
		"en-UK": `
<b>From bot's creator</b>: Hi %s!

@SneatBot helps to manage your day-to-day family life. Or you can create a space to manage your group/team/community.

I've spend lot's of time to make this bot useful, quick & reliable.I hope you'll like it.

You can learn about new features of the bot in @SneatApp channel where <a href="https://t.me/StarGiveaways_EN">we giveaway 500 🌟 EVERY month</a>.
`,
		"ru-RU": `
<b>От создателя бота:</b> Привет %s!

@SneatBot помогает организовать вашу семейную жизнь.
Так же можно создать пространство для управления группой/командой/сообществом.

Вы можете узнавать о новых возможностях бота в канале @SneatApp_ru где <a href="https://t.me/StarGiveaways_RU">мы разыгрываем 500 🌟 КАЖДЫЙ месяц</a>.
`,
	},
}
