/*
 * PhantomJS Page Rendering script
 *
 * Source: http://phantomjs.org/api/webpage/method/render-buffer.html
 */
var system = require('system');
var args = system.args;

function log() {
    system.stderr.writeLine(["INFO: "].slice.apply(arguments).join(" "));
}

if (args.length < 2) {
    log('Usage: phantomjs render.js url querySelector clickSelector resolution');
    phantom.exit();
    exit();
}

var url = args[1];
var selector = 'html';
if (args.length >= 3) {
    selector = args[2];
}
var click = "";
if (args.length >= 4) {
    click = args[3];
}
var res = '';
if (args.length >= 5) {
    res = args[4];
}
log('Rendering ' + url + ', query=' + selector + ', click=' + click + ', res=' + res);
try {
    var page = require('webpage').create();
    log('Opening ' + url);
    if (res == 'ipad') {
        page.viewportSize = {width: 768, height: 1024}
    } else if (res == 'desktop') {
        page.viewportSize = {width: 1280, height: 768}
    }
    page.open(url, function start(status) {
        try {
            log('Status ', status);
            page.evaluate(function(click) {
                if (click) {
                    document.querySelector(click).click();
                }
            }, click);
            log('Clicked, waiting to render');
            setTimeout(function() {
                try {
                log('Detecting clip');
                var clip = page.evaluate(function(s) {
                    return document.querySelector(s).getBoundingClientRect();
                }, selector);
                log('Clip rectangle: ' + JSON.stringify(clip));
                if (clip != null) {
                    log('Updating page rectangle');
                    page.clipRect = {
                        top: clip.top,
                        left: clip.left,
                        width: clip.width,
                        height: clip.height
                    };
                }
                log('Page rect: ' , JSON.stringify(page.clipRect));
                page.render('/tmp/screenshot.jpg', {format: 'jpg', quality: 85});
                log('Page rendered ');
                } catch(e) {
                    log('Error rendering page');
                    log(e);
                } finally {
                    phantom.exit();
                }
            }, 600);
        } catch (e) {
            log('Error rendering page');
            log(e);
            phantom.exit();
        } 
    });
} catch (e) {
    log('Unexpected error');
    log(e);
    phantom.exit();
}
