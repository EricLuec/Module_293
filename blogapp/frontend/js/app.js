function blogApp() {
    return {
        posts: [],
        form: { id: '', title: '', content: '' },

        fetchPosts() {
            fetch('/api/posts')
                .then(res => res.json())
                .then(data => this.posts = data)
        },

        savePost() {
            const method = this.form.id ? 'PUT' : 'POST'
            const url = this.form.id ? `/api/posts/${this.form.id}` : '/api/posts'

            fetch(url, {
                method,
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ title: this.form.title, content: this.form.content })
            }).then(() => {
                this.fetchPosts()
                this.resetForm()
            })
        },

        editPost(post) {
            this.form = { ...post, id: post.id.$oid || post.id }
        },

        deletePost(id) {
            fetch(`/api/posts/${id}`, { method: 'DELETE' })
                .then(() => this.fetchPosts())
        },

        resetForm() {
            this.form = { id: '', title: '', content: '' }
        }
    }
}
