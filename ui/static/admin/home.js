const Home = {
  template: '#home-page',
  data: () => ({
    votes: [],
    voters: [],
    input: {
      names: [],
      numberOfVoters: '',
      prefix: ''
    }
  }),
  computed: {
    votingExist: function() {
      return this.votes.length > 0 && this.voters.length > 0
    },
    validVotes: function() {
      let valid = 0
      this.voters.forEach(voter => {
        if (voter.status == 'voted') {
          valid++
        }
      })

      return valid
    },
    absentVotes: function() {
      let absent = 0
      this.voters.forEach(voter => {
        if (voter.status == 'init') {
          absent++
        }
      })

      return absent
    }
  },
  methods: {
    fetchVotes: async function() {
      let basicAuth = localStorage.getItem('basic-auth')
      let response = await fetch('/api/admin/votes', {
        headers: {
          'Authorization': `Basic ${basicAuth}`
        }
      })
      
      if (response.status == 200) {
        let data = await response.json()
        this.votes = data.data
      } else if (response.status == 401) {
        this.$router.push('/login')
      } else {
        UIkit.notification({
          message: `Sorry! Something error`,
          status: 'danger',
          pos: 'bottom-center',
          timeout: 3000
        })
      }
    },
    fetchVoters: async function() {
      let basicAuth = localStorage.getItem('basic-auth')
      let response = await fetch('/api/admin/voters', {
        headers: {
          'Authorization': `Basic ${basicAuth}`
        }
      })
      
      if (response.status == 200) {
        let data = await response.json()
        this.voters = data.data
      } else if (response.status == 401) {
        this.$router.push('/login')
      } else {
        UIkit.notification({
          message: `Sorry! Something error`,
          status: 'danger',
          pos: 'bottom-center',
          timeout: 3000
        })
      }
    },
    createVoting: async function() {
      let basicAuth = localStorage.getItem('basic-auth')
      let response = await fetch('/api/admin/create', {
        method: 'POST',
        headers: {
          'Authorization': `Basic ${basicAuth}`,
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          'names': this.input.names.split(/\r?\n/),
          'number_of_voters': parseInt(this.input.numberOfVoters),
          'prefix': this.input.prefix
        })
      })

      if (response.status == 200) {
        UIkit.modal('#modal-create').hide()
        this.fetchVotes()
        this.fetchVoters()
      } else {
        UIkit.notification({
          message: `Sorry! Something error`,
          status: 'danger',
          pos: 'bottom-center',
          timeout: 3000
        })
      }
    },
    resetVoting: async function() {
      let basicAuth = localStorage.getItem('basic-auth')
      let response = await fetch('/api/admin/reset', {
        method: 'POST',
        headers: {
          'Authorization': `Basic ${basicAuth}`
        }
      })

      if (response.status == 200) {
        UIkit.modal('#modal-reset').hide()
        this.votes = []
        this.voters = []
      } else {
        UIkit.notification({
          message: `Sorry! Something error`,
          status: 'danger',
          pos: 'bottom-center',
          timeout: 3000
        })
      }
    },
    logout: function() {
      localStorage.removeItem('basic-auth')
      this.$router.push('/login')
    }
  },
  mounted() {
    this.fetchVotes()
    this.fetchVoters()
  }
}
