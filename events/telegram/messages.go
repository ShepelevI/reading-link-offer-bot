package telegram

const msgHelp = `Я могу сохранять и хранить ваши страницы.
Я также могу предлагать вам их прочитать.
Для того, чтобы сохранить страницу, просто пришлите мне ссылку на нее.

Чтобы получить случайную страницу из вашего списка, отправьте мне команду /rnd.

Будьте осторожны! После закрытия прочтенной страницы
она будет удалена из вашего списка ссылок для чтения!`

const msgHello = "Привет! \n\n" + msgHelp

const (
	msgUnknownCommand = "Неизвестная команда"
	msgNoSavedPages   = "У вас нет ни одной сохраненной страницы"
	msgSaved          = "Сохранено"
	msgAlreadyExists  = "Эта страница уже есть в вашем списке"
)
