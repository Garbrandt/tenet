/*!
 * jQuery Vertical News Ticker Plugin
 *
 * http://www.jugbit.com/jquery-vticker-vertical-news-ticker/
 * http://github.com/kasp3r/vTicker
 *
 * Copyright 2013 Tadas Juozapaitis
 * Released under the MIT license:
 *   http://www.opensource.org/licenses/mit-license.php
 */
(function($) {
    $.fn.vTicker = function(options) {
        var defaults = {
            speed: 700,
            pause: 4000,
            showItems: 3,
            animation: '',
            mousePause: true,
            isPaused: false,
            direction: 'up',
            height: 0
        };

        var options = $.extend(defaults, options);

        moveUp = function(obj2, height, options) {
            if (options.isPaused)
                return;

            var obj = obj2.children('ul');

            var clone = obj.children('li:first').clone(true);

            if (options.height > 0) {
                height = obj.children('li:first').height();
            }

            obj.animate({ top: '-=' + height + 'px' }, options.speed, function() {
                $(this).children('li:first').remove();
                $(this).css('top', '0px');
            });

            if (options.animation == 'fade') {
                obj.children('li:first').fadeOut(options.speed);
                if (options.height == 0) {
                    obj.children('li:eq(' + options.showItems + ')').hide().fadeIn(options.speed).show();
                }
            }

            clone.appendTo(obj);
        };

        moveDown = function(obj2, height, options) {
            if (options.isPaused)
                return;

            var obj = obj2.children('ul');

            var clone = obj.children('li:last').clone(true);

            if (options.height > 0) {
                height = obj.children('li:first').height();
            }

            obj.css('top', '-' + height + 'px')
                .prepend(clone);

            obj.animate({ top: 0 }, options.speed, function() {
                $(this).children('li:last').remove();
            });

            if (options.animation == 'fade') {
                if (options.height == 0) {
                    obj.children('li:eq(' + options.showItems + ')').fadeOut(options.speed);
                }
                obj.children('li:first').hide().fadeIn(options.speed).show();
            }
        };

        return this.each(function() {
            var obj = $(this);
            var maxHeight = 0;

            obj.css({ overflow: 'hidden', position: 'relative' })
                .children('ul').css({ position: 'absolute', margin: 0, padding: 0 })
                .children('li').css({ margin: 0, padding: 0 });

            if (options.height == 0) {
                obj.children('ul').children('li').each(function() {
                    if ($(this).height() > maxHeight) {
                        maxHeight = $(this).height();
                    }
                });

                obj.children('ul').children('li').each(function() {
                    $(this).height(maxHeight);
                });

                obj.height(maxHeight * options.showItems);
            } else {
                obj.height(options.height);
            }

            var interval = setInterval(function() {
                if (options.direction == 'up') {
                    moveUp(obj, maxHeight, options);
                } else {
                    moveDown(obj, maxHeight, options);
                }
            }, options.pause);

            if (options.mousePause) {
                obj.bind("mouseenter", function() {
                    options.isPaused = true;
                }).bind("mouseleave", function() {
                    options.isPaused = false;
                });
            }
        });
    };
})(jQuery);