const Voting = {
  template: '#voting-page',
  data: () => ({
    votes: [],
    selectedId: ''
  }),
  methods: {
    fetchVotes: async function() {
      let voterCode = localStorage.getItem('voter-code')
      let response = await fetch('/api/voter/votes', {
        headers: {
          'voter-code': voterCode
        }
      })

      if (response.status == 200) {
        let data = await response.json()
        this.votes = data.data
      } else if (response.status == 401) {
        this.$router.push('/login')
      } else {
        UIkit.notification({
          message: `<span class="uk-text-light">Sorry! Something error</span>`,
          status: 'danger',
          pos: 'bottom-center',
          timeout: 3000
        })
      }
    },
    doVote: async function() {
      if (this.selectedId != '') {
        let voterCode = localStorage.getItem('voter-code')
        let response = await fetch('/api/voter/vote', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            'voter-code': voterCode
          },
          body: JSON.stringify({
            'selected_id': this.selectedId
          })
        })

        if (response.status == 200) {
          localStorage.removeItem('voter-code')
          UIkit.modal('#modal-vote').hide()
          // go to thankyou
          this.$router.push('/login')
        } else {
          UIkit.notification({
            message: `<span class="uk-text-light">Sorry! Something error</span>`,
            status: 'danger',
            pos: 'bottom-center',
            timeout: 3000
          })
        }
      }
    }
  },
  mounted() {
    this.fetchVotes()
  }
}
