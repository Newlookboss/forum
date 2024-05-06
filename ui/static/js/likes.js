
const PostLikeButtons = document.querySelectorAll(".likeButton")
const PostDislikeButtons = document.querySelectorAll(".dislikeButton")
const PostLikeCntElement = document.querySelectorAll(".likeCount");
const PostDislikeCntElement = document.querySelectorAll(".dislikeCount");
const PostLikeCnt = Array.from(PostLikeCntElement).map(element => +element.innerText);
const PostDislikeCnt = Array.from(PostDislikeCntElement).map(element => +element.innerText);


const commentLikeButtons = document.querySelectorAll(".commentLikeButton")
const commentDislikeButtons = document.querySelectorAll(".commentDislikeButton")
const commentLikeCntElement = document.querySelectorAll(".commentLikeCount");
const commentDislikeCntElement = document.querySelectorAll(".commentDislikeCount");
const commentLikeCnt = Array.from(commentLikeCntElement).map(element => +element.innerText);
const commentDislikeCnt = Array.from(commentDislikeCntElement).map(element => +element.innerText);



function submitLike(body, url, msg) {
    return new Promise((resolve, reject) => {
        fetch(url, {
            method: "POST",
            body: JSON.stringify(body),
            headers: {
                "Content-Type": "application/json"
            }
        })
        .then(response => {
            if (response.ok) {
                console.log("Successful post " + msg);
                resolve(); // Resolve the promise if the request is successful
            } else {
                console.error("Failed post " + msg);
                reject("Failed post " + msg); // Reject the promise if the request fails
            }
        })
        .catch(error => {
            console.error(error);
            reject(error); // Reject the promise if there's an error
        });
    });
}
document.addEventListener("DOMContentLoaded", function() {
    const PostLikeButtons = document.querySelectorAll(".likeButton");
    const PostDislikeButtons = document.querySelectorAll(".dislikeButton");

    PostLikeButtons.forEach(likeButton => {
        likeButton.addEventListener("click", () => {
            let likeCountElement = likeButton.querySelector(".likeCount");
            let dislikeButton = likeButton.nextElementSibling; // находим кнопку "дизлайк"
            let dislikeCountElement = dislikeButton.querySelector(".dislikeCount");

            let currentLikeCount = parseInt(likeCountElement.innerText);
            let currentDislikeCount = parseInt(dislikeCountElement.innerText);

            // Если кнопка "лайк" уже нажата, то сбрасываем ее
            if (likeButton.getAttribute("post-liked") === "true") {
                likeCountElement.innerText = currentLikeCount - 1;
                likeButton.querySelector(".likeIcon").src = "/static/img/like.png";
                likeButton.setAttribute("post-liked", "false");
            } else {
                likeCountElement.innerText = currentLikeCount + 1;
                likeButton.querySelector(".likeIcon").src = "/static/img/liked.png";
                likeButton.setAttribute("post-liked", "true");

                // Если кнопка "дизлайк" нажата, сбрасываем ее
                if (dislikeButton.getAttribute("post-disliked") === "true") {
                    dislikeCountElement.innerText = currentDislikeCount - 1;
                    dislikeButton.querySelector(".dislikeIcon").src = "/static/img/dislike.png";
                    dislikeButton.setAttribute("post-disliked", "false");
                }
            }
        });
    });

    PostDislikeButtons.forEach(dislikeButton => {
        dislikeButton.addEventListener("click", () => {
            let dislikeCountElement = dislikeButton.querySelector(".dislikeCount");
            let likeButton = dislikeButton.previousElementSibling; // находим кнопку "лайк"
            let likeCountElement = likeButton.querySelector(".likeCount");

            let currentDislikeCount = parseInt(dislikeCountElement.innerText);
            let currentLikeCount = parseInt(likeCountElement.innerText);

            // Если кнопка "дизлайк" уже нажата, то сбрасываем ее
            if (dislikeButton.getAttribute("post-disliked") === "true") {
                dislikeCountElement.innerText = currentDislikeCount - 1;
                dislikeButton.querySelector(".dislikeIcon").src = "/static/img/dislike.png";
                dislikeButton.setAttribute("post-disliked", "false");
            } else {
                dislikeCountElement.innerText = currentDislikeCount + 1;
                dislikeButton.querySelector(".dislikeIcon").src = "/static/img/disliked.png";
                dislikeButton.setAttribute("post-disliked", "true");

                // Если кнопка "лайк" нажата, сбрасываем ее
                if (likeButton.getAttribute("post-liked") === "true") {
                    likeCountElement.innerText = currentLikeCount - 1;
                    likeButton.querySelector(".likeIcon").src = "/static/img/like.png";
                    likeButton.setAttribute("post-liked", "false");
                }
            }
        });
    });
});



PostLikeButtons.forEach((button, index) => {
    let PostID = +button.getAttribute("post-id");
    let isPostLiked = button.getAttribute("post-liked");
    isPostLiked = (isPostLiked?.toLowerCase?.() === 'true');
    
    button.addEventListener("click", () => {
        isPostLiked = !isPostLiked;
        let body = {PostLikeCount: PostLikeCnt[index], PostID: PostID, isLiked: isPostLiked}
        let url = "/post/like"
        submitLike(body, url, "like");
    })
})

PostDislikeButtons.forEach((button, index) => {
    let PostID = +button.getAttribute("post-id");
    let isPostDisliked = button.getAttribute("post-disliked");
    isPostDisliked = (isPostDisliked?.toLowerCase?.() === 'true');
    
    button.addEventListener("click", () => {
        isPostDisliked = !isPostDisliked;

        let body = {PostDislikeCount: PostDislikeCnt[index], PostID: PostID, isDisliked: isPostDisliked}
        let url = "/post/dislike"
        submitLike(body, url, "dislike");
    })
})



function submitCommentLike(body, url, msg) {
    return new Promise((resolve, reject) => {
        fetch(url, {
            method: "POST",
            body: JSON.stringify(body),
            headers: {
                "Content-Type": "application/json"
            }
        })
        .then(response => {
            if (response.ok) {
                console.log("Successful comment " + msg);
                resolve(); // Resolve the promise if the request is successful
            } else {
                console.error("Failed comment " + msg);
                reject("Failed comment " + msg); // Reject the promise if the request fails
            }
        })
        .catch(error => {
            console.error(error);
            reject(error); // Reject the promise if there's an error
        });
    });
}
document.addEventListener("DOMContentLoaded", function() {
    const CommentLikeButtons = document.querySelectorAll(".commentLikeButton")
const CommentDislikeButtons = document.querySelectorAll(".commentDislikeButton")



CommentLikeButtons.forEach(likeButton => {
        likeButton.addEventListener("click", () => {
            let likeCountElement = likeButton.querySelector(".commentLikeCount");
            let dislikeButton = likeButton.nextElementSibling; // находим кнопку "дизлайк"
            let dislikeCountElement = dislikeButton.querySelector(".commentDislikeCount");

            let currentLikeCount = parseInt(likeCountElement.innerText);
            let currentDislikeCount = parseInt(dislikeCountElement.innerText);

            // Если кнопка "лайк" уже нажата, то сбрасываем ее
            if (likeButton.getAttribute("comment-liked") === "true") {
                likeCountElement.innerText = currentLikeCount - 1;
                likeButton.querySelector(".commentLikeIcon").src = "/static/img/like.png";
                likeButton.setAttribute("comment-liked", "false");
            } else {
                likeCountElement.innerText = currentLikeCount + 1;
                likeButton.querySelector(".commentLikeIcon").src = "/static/img/liked.png";
                likeButton.setAttribute("comment-liked", "true");

                // Если кнопка "дизлайк" нажата, сбрасываем ее
                if (dislikeButton.getAttribute("comment-disliked") === "true") {
                    dislikeCountElement.innerText = currentDislikeCount - 1;
                    dislikeButton.querySelector(".commentDislikeIcon").src = "/static/img/dislike.png";
                    dislikeButton.setAttribute("comment-disliked", "false");
                }
            }
        });
    });
    CommentDislikeButtons.forEach(dislikeButton => {
        dislikeButton.addEventListener("click", () => {
            let dislikeCountElement = dislikeButton.querySelector(".commentDislikeCount");
            let likeButton = dislikeButton.previousElementSibling; // находим кнопку "лайк"
            let likeCountElement = likeButton.querySelector(".commentLikeCount");

            let currentDislikeCount = parseInt(dislikeCountElement.innerText);
            let currentLikeCount = parseInt(likeCountElement.innerText);

            // Если кнопка "дизлайк" уже нажата, то сбрасываем ее
            if (dislikeButton.getAttribute("comment-disliked") === "true") {
                dislikeCountElement.innerText = currentDislikeCount - 1;
                dislikeButton.querySelector(".commentDislikeIcon").src = "/static/img/dislike.png";
                dislikeButton.setAttribute("comment-disliked", "false");
            } else {
                dislikeCountElement.innerText = currentDislikeCount + 1;
                dislikeButton.querySelector(".commentDislikeIcon").src = "/static/img/disliked.png";
                dislikeButton.setAttribute("comment-disliked", "true");

                // Если кнопка "лайк" нажата, сбрасываем ее
                if (likeButton.getAttribute("comment-liked") === "true") {
                    likeCountElement.innerText = currentLikeCount - 1;
                    likeButton.querySelector(".commentLikeIcon").src = "/static/img/like.png";
                    likeButton.setAttribute("comment-liked", "false");
                }
            }
        });
    });
});
commentLikeButtons.forEach((button, index) => {
    let PostID = +button.getAttribute("post-id");
    let commentID = +button.getAttribute("comment-id");
    let isCommentLiked = button.getAttribute("comment-liked");
    isCommentLiked = (isCommentLiked?.toLowerCase?.() === 'true');
    
    button.addEventListener("click", () => {
        isCommentLiked = !isCommentLiked;
        let body = {commentLikeCount: commentLikeCnt[index], commentID: commentID, postID: PostID, isCommentLiked: isCommentLiked}
        let url = "/post/commentLike"
        submitCommentLike(body, url, "like");
    })
})

commentDislikeButtons.forEach((button, index) => {
    let PostID = +button.getAttribute("post-id");
    let commentID = +button.getAttribute("comment-id");
    let isCommentDisliked = button.getAttribute("comment-disliked");
    isCommentDisliked = (isCommentDisliked?.toLowerCase?.() === 'true');
    
    button.addEventListener("click", () => {
        isCommentDisliked = !isCommentDisliked;
        let body = {commentDislikeCount: commentDislikeCnt[index], commentID: commentID, postID: PostID, isCommentDisliked: isCommentDisliked}
        let url = "/post/commentDislike"
        submitCommentLike(body, url, "dislike");
    })
})