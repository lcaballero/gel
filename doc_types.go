package gel

func Html5(head, body View) View {
	return Fragment{
		Text("<!doctype html>"),
		Html(
			head,
			body,
		),
	}
}
