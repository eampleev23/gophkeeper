package handlers

import (
	"encoding/json"
	"github.com/eampleev23/gophkeeper/internal/models"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

/*
На вход хэндлер ожидает json такого формата(пользователь должен быть авторизован):
{
    "meta-value": "стихотворение Пушкина Мороз и солнце день чудесный",
    "text": "Мороз и солнце; день чудесный!
Еще ты дремлешь, друг прелестный —
Пора, красавица, проснись:
Открой сомкнуты негой взоры
Навстречу северной Авроры,
Звездою севера явись!

Вечор, ты помнишь, вьюга злилась,
На мутном небе мгла носилась;
Луна, как бледное пятно,
Сквозь тучи мрачные желтела,
И ты печальная сидела —
А нынче… погляди в окно:

Под голубыми небесами
Великолепными коврами,
Блестя на солнце, снег лежит;
Прозрачный лес один чернеет,
И ель сквозь иней зеленеет,
И речка подо льдом блестит.

Вся комната янтарным блеском
Озарена. Веселым треском
Трещит затопленная печь.
Приятно думать у лежанки.
Но знаешь: не велеть ли в санки
Кобылку бурую запречь?

Скользя по утреннему снегу,
Друг милый, предадимся бегу
Нетерпеливого коня
И навестим поля пустые,
Леса, недавно столь густые,
И берег, милый для меня.",
}
*/

func (h *Handlers) AddText(w http.ResponseWriter, r *http.Request) {
	h.l.ZL.Info("Add Text handler has started..")
	// Проверяем формат запроса
	contentType := r.Header.Get("Content-Type")
	supportsJSON := strings.Contains(contentType, "application/json")
	if !supportsJSON {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Проверяем, не авторизован ли пользователь, отправивший запрос.
	ownerID, isAuth, err := h.GetUserID(r)
	if err != nil {
		h.l.ZL.Error("GetUserID fail")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !isAuth {
		h.l.ZL.Info("Пользователь не авторизован, в ответе отказано")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	h.l.ZL.Info("Пользователь авторизован, можем двигаться дальше")
	// Получаем данные в случае корректного запроса.

	var inputModel models.TextDataItem
	// Декодер работает потоково, кажется это правильнее + короче, чем анмаршал.
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&inputModel); err != nil {
		h.l.ZL.Info("cannot decode request JSON body", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// записываем значение ид автора запроса
	inputModel.OwnerID = ownerID
	outputModel, err := h.serv.InsertTextDataItem(r.Context(), inputModel)
	if err != nil {
		h.l.ZL.Error("h.serv.InsertDataItem fail..", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Если мы здесь, то текст успешно добавлен.
	h.l.ZL.Info("Success creating new text data item", zap.Any("outputModel", outputModel))
	w.WriteHeader(http.StatusOK)
	return
}
