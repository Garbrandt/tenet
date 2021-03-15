var options = {
	chart: {
		height: 203,
		type: 'radialBar',
		toolbar: {
			show: false,
		},
	},
	plotOptions: {
		radialBar: {
			track: {
		    background: '#ced1e0',
			},
			dataLabels: {
				name: {
					fontSize: '12px',
					fontColor: 'black',
          fontFamily: 'Open Sans, sans-serif',
				},
				value: {
					fontSize: '21px',
					fontFamily: 'Open Sans, sans-serif',
				},
				total: {
					show: true,
					label: 'Tasks',
					formatter: function (w) {
						// By default this function returns the average of all series. The below is just an example to show the use of custom formatter function
						return '21'
					}
				}
			}
		},
	},
	series: [85, 60, 45],
	labels: ['New', 'Completed', 'Pending'],
	colors: ['#007ae1', '#ff3e3e', '#00bb42'],
}

var chart = new ApexCharts(
	document.querySelector("#radialTasks"),
	options
);
chart.render();