package main

func main() {
	hhKz := JobSite{}
    var bob Person = "Bob"
    hhKz.AddVacancies("Senior HTML Developer")
    hhKz.Subscribe(bob)
    hhKz.AddVacancies("Java Junior Developer")
    hhKz.RemoveVacancy("Senior HTML Developer")
}