package main

type JobSite struct {
	vacancies []string
	subscribers []Observer
}

func (job *JobSite) AddVacancies(vacancy string) {
	job.vacancies = append(job.vacancies, vacancy)
	job.SendAll()
}

func (job *JobSite) RemoveVacancy(vacancy string) {
	var newVacanciesList []string
	for _, iterator := range job.vacancies {
		if vacancy == iterator {
			continue
		}
		newVacanciesList = append(newVacanciesList, iterator)
	}
	job.vacancies = newVacanciesList
	job.SendAll()
}

func (job *JobSite) Subscribe(observer Observer) {
	job.subscribers = append(job.subscribers, observer)
	job.SendAll()
}

func (job *JobSite) Unsubscribe(observer Observer) {
	var newSubscribersList []Observer
	for _, subscriber := range job.subscribers {
		if observer == subscriber {
			continue
		}
		newSubscribersList = append(newSubscribersList, subscriber)
	}
	job.subscribers = newSubscribersList
	job.SendAll()
}

func (job *JobSite) SendAll() {
	for _, observer := range job.subscribers {
		observer.HandleEvent(job.vacancies)
	}
}