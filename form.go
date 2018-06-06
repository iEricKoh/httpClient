package httpClient

import "net/url"

type FormBuilder struct {
	Form *Form
}

func (f *FormBuilder) BuildForm() url.Values {
	formData := url.Values{}

	for key, value := range *f.Form {
		str, ok := value.(string)
		if ok {
			formData.Set(key, str)
		}
	}

	return formData
}
