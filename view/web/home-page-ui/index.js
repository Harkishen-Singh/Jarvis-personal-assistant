$(document).ready(function() {
    // Light-Mode Dark-Mode switch toggle effect
    $('#theme_switch').on('click', function() {
        $('body').toggleClass('dark-mode');
        $('.right-side-wrapper').toggleClass('dark-mode');
        $('.left-side-wrapper').toggleClass('dark-mode');
        $('.help-btn').toggleClass('btn-dark-mode');
        $('h6, .message-textbox').toggleClass('txt-dark-mode');
        $('.message-textbox').toggleClass('m-txtbox-darkmode');
    })

    $('#message-send').on('click', function(event) {
        event.preventDefault();
        var date_time = "now";
        var msg = $('#message-box').val();
        if (msg.replace(/\s/g, '').length) {
            if ($('#theme_switch').is(":checked")) {
                var new_txt = "<div class='animate-txt-wrap user-msg'>\
                <div class='wrap-p-chat-txt send-txt-wrapper'><h6 class='p-chat-sent-txt'>" + msg + "</h6></div>"
                + "<h6 class='p-chat-snd-date-time txt-dark-mode'>" + "me <strong>" + date_time + "</strong></h6></div>";
            } else {
                var new_txt = "<div class='animate-txt-wrap user-msg'>\
                <div class='wrap-p-chat-txt send-txt-wrapper'><h6 class='p-chat-sent-txt'>" + msg + "</h6></div>"
                + "<h6 class='p-chat-snd-date-time'>" + "me <strong>" + date_time + "</strong></h6></div>";
            }
            var new_txt = "<div class='animate-txt-wrap user-msg'>\
                <div class='wrap-p-chat-txt send-txt-wrapper'><h6 class='p-chat-sent-txt'>" + msg + "</h6></div>"
                + "<h6 class='p-chat-snd-date-time'>" + "me <strong>" + date_time + "</strong></h6></div>";
            var $data = $(new_txt);
            $('.conversation-wrapper').append($data);
            $('.conversation-wrapper').animate({
                scrollTop: $('.conversation-wrapper').get(0).scrollHeight
            }, 1500);
            $data.animate({'margin-top': '10px'}, 230);
            $('#message-box').val("");
        }
    })
});