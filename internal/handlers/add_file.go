package handlers

import (
	"github.com/eampleev23/gophkeeper/internal/models"
	"go.uber.org/zap"
	"io"
	"net/http"
	"os"
)

func (h *Handlers) AddFile(w http.ResponseWriter, r *http.Request) {
	h.l.ZL.Info("Add File handler has started..")

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

	// Maximum upload of 10 MB files
	r.ParseMultipartForm(10 << 20)
	// Get handler for filename, size and headers
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		h.l.ZL.Error("Error Retrieving the File", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// получаем мета-данные
	metaValue := r.Header.Get("Meta-Value")

	h.l.ZL.Info("got incoming HTTP request",
		zap.String("Uploaded File:", handler.Filename),
		zap.Int64("File Size:", handler.Size),
		zap.Any("MIME Header:", handler.Header),
		zap.String("Meta-Value:", metaValue),
	)

	// Create file
	servPath := "../../serv_file_store/"
	dst, err := os.Create(servPath + handler.Filename)
	defer dst.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.l.ZL.Info("Successfully Uploaded File")

	var inputModel models.FileDataItem
	inputModel.OwnerID = ownerID
	inputModel.MetaValue = metaValue
	inputModel.ServerPath = servPath + handler.Filename

	outputModel, err := h.serv.InsertFileDataItem(r.Context(), inputModel)
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
