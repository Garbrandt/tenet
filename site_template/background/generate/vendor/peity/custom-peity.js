// Updating Chart
$(function(){
	var updatingChart = $(".updating-chart").peity("line", {
		width: '80%',
		height: 75,
		stroke: "#007ae1",
		fill: false,
		strokeWidth: 5,
	})
	setInterval(function() {
		var random = Math.round(Math.random() * 5)
		var values = updatingChart.text().split(",")
		values.shift()
		values.push(random)
		updatingChart
		.text(values.join(","))
		.change()
	}, 1000);
})


// Bar Chart
$(function(){
	$(".bar-one").peity("bar", {
		width: 90,
		height: 40,
		fill: ["#007ae1"],
	})
});
$(function(){
	$(".bar-two").peity("bar", {
		width: 90,
		height: 40,
		fill: ["#ff3e3e"],
	})
});
$(function(){
	$(".bar-three").peity("bar", {
		width: 90,
		height: 40,
		fill: function(_, i, all) {
			var g = parseInt((i / all.length) * 255)
			return "rgb(15, " + g + ", 220)"
		}
	})
});

// Line Chart
$(function(){
	$(".line-one").peity("line", {
		width: 90,
		height: 16,
		stroke: ["#ffbf05"],
		fill: false,
		strokeWidth: 3,
	})
});
$(function(){
	$(".line-two").peity("line", {
		width: 90,
		height: 16,
		stroke: ["#007ae1"],
		fill: false,
		strokeWidth: 3,
	})
});
$(function(){
	$(".line-three").peity("line", {
		width: 90,
		height: 16,
		stroke: ["#00bb42"],
		fill: false,
		strokeWidth: 3,
	})
});
$(function(){
	$(".line-four").peity("line", {
		width: 90,
		height: 16,
		stroke: ["#ffbf05"],
		fill: false,
		strokeWidth: 3,
	})
});
$(function(){
	$(".line-five").peity("line", {
		width: 90,
		height: 16,
		stroke: ["#ff3e3e"],
		fill: false,
		strokeWidth: 3,
	})
});

$(function(){
	$(".line-six").peity("line", {
		width: 90,
		height: 18,
		stroke: ["#007ae1"],
		fill: false,
		strokeWidth: 3,
	})
});
$(function(){
	$(".line-seven").peity("line", {
		width: 150,
		height: 25,
		stroke: ["#00bb42"],
		fill: false,
		strokeWidth: 4,
	})
});
