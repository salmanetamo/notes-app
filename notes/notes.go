package notes

import "time"

func ViewNotes() {
	for _, note := range notes {
		noteRow := note.GetRow()
		noteRow.Build()
	}
}

func GetNoteById(id int) Note {
	for _, note := range notes {
		if note.id == id {
			return note
		}
	}
	return Note{}
}

func CreateNote(title, description string) {
	newNote := Note{id: len(notes) + 1, title: title, description: description, createdAt: time.Now(), updatedAt: time.Now()}
	notes = append(notes, newNote)
}

func UpdateNote(id int, updatedNote Note) Note {
	updatedNote.updatedAt = time.Now()
	var index int
	for idx, note := range notes {
		if note.id == id {
			index = idx
		}
	}
	notes[index] = updatedNote
	return updatedNote
}

func DeleteNote(id int) {
	var index int
	for idx, note := range notes {
		if note.id == id {
			index = idx
		}
	}
	notes = append(notes[:index], notes[index+1:]...)
}

func IsValidNoteId(id int) bool {
	for _, note := range notes {
		if note.id == id {
			return true
		}
	}
	return false
}
