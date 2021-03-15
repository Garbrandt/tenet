$(document).ready(function(){
  countdown();
  setInterval(countdown, 1000);
  function countdown () {
  var now = moment(), // get the current moment
    // May 28, 2013 @ 12:00AM
    then = moment([2020, 10, 15]),
    // get the difference from now to then in ms
    ms = then.diff(now, 'milliseconds', true);
    // If you need years, uncomment this line and make sure you add it to the concatonated phrase
    /*
    years = Math.floor(moment.duration(ms).asYears());
    then = then.subtract('years', years);
    */
    // update the duration in ms
    ms = then.diff(now, 'milliseconds', true);
    // get the duration as months and round down
    // months = Math.floor(moment.duration(ms).asMonths());
 
    // // subtract months from the original moment (not sure why I had to offset by 1 day)
    // then = then.subtract('months', months).subtract('days', 1);
    // update the duration in ms
    ms = then.diff(now, 'milliseconds', true);
    days = Math.floor(moment.duration(ms).asDays());
 
    then = then.subtract('days', days);
    // update the duration in ms
    ms = then.diff(now, 'milliseconds', true);
    hours = Math.floor(moment.duration(ms).asHours());
 
    then = then.subtract('hours', hours);
    // update the duration in ms
    ms = then.diff(now, 'milliseconds', true);
    minutes = Math.floor(moment.duration(ms).asMinutes());
 
    then = then.subtract('minutes', minutes);
    // update the duration in ms
    ms = then.diff(now, 'milliseconds', true);
    seconds = Math.floor(moment.duration(ms).asSeconds());
    
    // concatonate the variables
    diff = '<li class="num">' + days + ' <span class="text"> Days</span></li><li class="num">' + hours + ' <span class="text"> Hours</span></li><li class="num">' + minutes + ' <span class="text"> Minutes</span></li><li class="num">' + seconds + ' <span class="text"> Seconds</span></li>';
    $('#countdown').html(diff);
  }
 
});