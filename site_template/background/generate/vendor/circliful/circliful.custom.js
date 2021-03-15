$( document ).ready(function() {

	$("#todaysTarget").circliful({
		animation: 1,
		animationStep: 5,
		foregroundBorderWidth: 14,
		backgroundBorderWidth: 14,
		percent: 70,
		textStyle: 'font-size: 12px;',
		fontColor: '#2e323c',
		foregroundColor: '#f53a40',
		backgroundColor: '#d6d8e4',
	});

	$("#todaysTarget1").circliful({
		animation: 1,
		animationStep: 5,
		foregroundBorderWidth: 14,
		backgroundBorderWidth: 14,
		percent: 70,
		textStyle: 'font-size: 12px;',
		fontColor: '#2e323c',
		foregroundColor: '#008e18',
		backgroundColor: '#d6d8e4',
	});

	$("#newCustomers").circliful({
		animation: 1,
		animationStep: 5,
		foregroundBorderWidth: 14,
		backgroundBorderWidth: 14,
		percent: 85,
		textStyle: 'font-size: 12px;',
		fontColor: '#2e323c',
		foregroundColor: '#00bb42',
		backgroundColor: '#d6d8e4',
	});	

	$("#overallSales").circliful({
		animation: 1,
		animationStep: 5,
		foregroundBorderWidth: 16,
		backgroundBorderWidth: 10,
		percent: 92,
		textStyle: 'font-size: 12px;',
		fontColor: '#2e323c',
		foregroundColor: '#f53a40',
		backgroundColor: '#d6d8e4',
		multiPercentage: 1,
		percentages: [10, 20, 30],
	});


	$("#overallExpenses").circliful({
		animation: 1,
		animationStep: 5,
		foregroundBorderWidth: 16,
		backgroundBorderWidth: 10,
		percent: 78,
		fontColor: '#2e323c',
		foregroundColor: '#2fcc7e',
		backgroundColor: '#d6d8e4',
		multiPercentage: 1,
		percentages: [10, 20, 30]
	});
	$("#overallIncome").circliful({
		animation: 1,
		animationStep: 5,
		foregroundBorderWidth: 16,
		backgroundBorderWidth: 10,
		percent: 80,
		fontColor: '#2e323c',
		foregroundColor: '#ff991a',
		backgroundColor: '#d6d8e4',
		multiPercentage: 1,
		percentages: [10, 20, 30]
	});





	// With Icons
	$("#overallRevenue").circliful({
		animationStep: 5,
		foregroundBorderWidth: 7,
		backgroundBorderWidth: 7,
		percent: 80,
		fontColor: '#2e323c',
		foregroundColor: '#f53a40',
		backgroundColor: '#d6d8e4',
	});


	// With Icons
	$("#overallRevenue1").circliful({
		animationStep: 5,
		foregroundBorderWidth: 10,
		backgroundBorderWidth: 10,
		percent: 80,
		fontColor: '#2e323c',
		foregroundColor: '#00bb42',
		backgroundColor: '#d6d8e4',
	});

	// With Icons
	$("#overallRevenue2").circliful({
		animationStep: 5,
		foregroundBorderWidth: 10,
		backgroundBorderWidth: 10,
		percent: 80,
		fontColor: '#2e323c',
		foregroundColor: '#008e18',
		backgroundColor: '#d6d8e4',
	});


	$("#projectPlanning").circliful({
		animationStep: 5,
		foregroundBorderWidth: 12,
		backgroundBorderWidth: 7,
		percent: 100,
		fontColor: '#2e323c',
		foregroundColor: '#da9d46',
		backgroundColor: '#d6d8e4',
		icon: '\ea1b',
		iconColor: '#da9d46',
		iconPosition: 'middle',
		textBelow: true,
		animation: 1,
		animationStep: 1,
		start: 2,
		showPercent: 1,		
	});
	

	$("#projectDesign").circliful({
		animationStep: 5,
		foregroundBorderWidth: 12,
		backgroundBorderWidth: 7,
		percent: 100,
		fontColor: '#2e323c',
		foregroundColor: '#da9d46',
		backgroundColor: '#d6d8e4',
		icon: '\ea40',
		iconColor: '#da9d46',
		iconPosition: 'middle',
		textBelow: true,
		animation: 1,
		animationStep: 1,
		start: 2,
		showPercent: 1,
	});


});

