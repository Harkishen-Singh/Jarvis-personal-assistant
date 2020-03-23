$(document).ready(function() {
    $('#message-send').on('click', function(event) {
        event.preventDefault();
        var date_time = "now";
        var msg = $('#message-box').val();
        if (msg.replace(/\s/g, '').length) {
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