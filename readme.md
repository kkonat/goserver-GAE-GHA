Creating a Go 1.12+ Runtime Environment on Google App Engine with simple two service app example

create project:
`gcloud projects create kk-goserver`

kk-goserver must be globally (for the whole gcloud) unique name

Initialize config, login et stuff
`gclloud init`

I have two configurations: mieczu and default. I want to use mieczu so I choose 1

```
Pick configuration to use:
 [1] Re-initialize this configuration [mieczu] with new settings
 [2] Create a new configuration
 [3] Switch to and re-initialize existing configuration: [default]
Please enter your numeric choice:  1
```

to pick the other configuration I would have to run `gcloud init` again and choose 3

Now login...

```
Choose the account you would like to use to perform operations for this
configuration:
 [1] krzysztof.konatowicz@gmail.com
 [2] Log in with a new account
Please enter your numeric choice:  1
```

And then pick the created poroject:

```
Pick cloud project to use:
 [1] ******************
 [2] ******************
 [3] kk-goserver
 [4] ******************
 [5] ******************
 [6] ******************
 [7] ******************
 [8] Enter a project ID
 [9] Create a new project
Please enter numeric choice or text value (must exactly match list item):  3
```

make sure the
Compute Engine API is enabled for your project on the
https://console.developers.google.com/apis page.

click on "+ Enable APIs and Services" on top and search for "Compute Engine API" then enable it.

It requires billing, so you have to specify your billing account for the project. It should not cost you money if you stay within the free tier.

After you select your billing account, you can enable the Compute Engine API.

App engine does not require credentials, so no need to click on the "Create credentials" button.
