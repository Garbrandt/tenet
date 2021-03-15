// Morris Donut
Morris.Donut({
	element: 'donutColors',
	data: [
		{value: 30, label: 'foo'},
		{value: 15, label: 'bar'},
		{value: 10, label: 'baz'},
		{value: 5, label: 'A really really long label'}
	],
	backgroundColor: '#ffffff',
	labelColor: '#2e323c',
	colors:['#007ae1', '#ff3e3e', '#00bb42', '#ffbf05'],
	resize: true,
	hideHover: "auto",
	gridLineColor: "#ced1e0",
	formatter: function (x) { return x + "%"}
});