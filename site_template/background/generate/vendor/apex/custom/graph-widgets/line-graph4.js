var options = {
	chart: {
		height: 240,
		type: 'area',
		toolbar: {
			show: false,
		},
	},
	colors: ['#007ae1', '#f3a33c'],
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
    name: 'Errors',
    data: [2, 13, 22, 18, 21, 25, 11]
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
	document.querySelector("#lineGraph4"),
	options
);

chart.render();