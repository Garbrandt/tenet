var options = {
	chart: {
		height: 240,
		type: 'area',
		toolbar: {
			show: false,
		},
	},
	colors: ['#00bb42', '#00bb42'],
	dataLabels: {
		enabled: false,
	},
	legend: {
  	show: false,
  },
	stroke: {
		show: true,
		curve: 'smooth',
		width: 3,
		lineCap: 'square'
	},
	series: [{
		name: 'Orders',
		data: [5000, 40000, 50000, 90000, 100000, 85000, 95000]
	}],
	xaxis: {
		axisBorder: {
			show: false
		},
		categories: ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"],
		axisTicks: {
			show: true
		},
		crosshairs: {
			show: true
		},
		labels: {
			offsetX: 0,
			offsetY: 5,
		}
	},
	yaxis: {
		labels: {
			formatter: function(value, index) {
				return (value / 1000) + 'K'
			},
			offsetX: -15,
			offsetY: 20,
		}
	},
	grid: {
		borderColor: '#ced1e0',
		strokeDashArray: 5,
		xaxis: {
			lines: {
				show: true
			}
		},   
		yaxis: {
			lines: {
				show: false,
			}	
		},
		padding: {
			top: 0,
			right: 0,
			bottom: 0,
			left: 0
		}, 
	},
}

var chart = new ApexCharts(
	document.querySelector("#lineGraph3"),
	options
);

chart.render();