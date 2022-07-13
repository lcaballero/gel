package gel

func Html5(head, body View) View {
	return Frag(
		Text("<!doctype html>"),
		Html(
			head,
			body,
		),
	)
}
