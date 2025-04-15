function postsApp() {
    return {
        posts: [],
        async fetchPosts() {
            const res = await fetch('http://localhost:8000/api/posts');
            const rawData = await res.json();
            console.log("Daten empfangen:", rawData);
            this.posts = rawData.map(p => ({
                id: p._id,
                title: p.title,
                content: p.content,
                date: new Date(p.date).toLocaleString()
            }));
        },

        async deletePost(id) {
            if (confirm("Diesen Post wirklich löschen?")) {
                await fetch(`http://localhost:8000/api/posts/${id}`, { method: 'DELETE' });
                this.fetchPosts();
            }
        },
        init() {
            this.fetchPosts();
        }
    };
}

function createPostForm() {
    return {
        title: '',
        content: '',
        async submitPost() {
            await fetch('http://localhost:8000/api/posts', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ title: this.title, content: this.content })
            });
            window.location.href = 'index.html';
        }
    };
}

function editPostForm() {
    return {
        id: '',
        title: '',
        content: '',
        async loadPost() {
            const url = new URLSearchParams(window.location.search);
            this.id = url.get('id');
            const res = await fetch(`/api/posts/${this.id}`);
            const data = await res.json();
            this.title = data.title;
            this.content = data.content;
        },
        async updatePost() {
            await fetch(`/api/posts/${this.id}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ title: this.title, content: this.content })
            });
            window.location.href = 'index.html';
        },
        init() {
            this.loadPost();
        }
    };
}
