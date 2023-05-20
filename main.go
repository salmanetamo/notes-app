package main

import (
	"fmt"
	"notes/menu"
	"notes/notes"
	"notes/ui"
	"notes/utils"
)

func main() {
	fmt.Println("Welcome to the Notes app!")
loop:
	for {
		menu.Print()
		choice := menu.GetUserChoice()
		switch choice {
		case 1:
			viewNotes()
		case 2:
			viewNotes()
			viewNote()
		case 3:
			viewNotes()
			updateNote()
		case 4:
			viewNotes()
			deleteNote()
		case 5:
			createNote()
			viewNotes()
		case 0:
			break loop
		default:
			fmt.Println("Invalid option")
		}
	}

}

func viewNotes() {
	fmt.Println("\tNotes Table")
	headerRow := ui.Row{
		ui.Cell{Value: "ID", Length: ui.IdColumnLength, IsFirstInRow: true, IsLastInRow: false},
		ui.Cell{Value: "TITLE", Length: ui.TitleColumnLength, IsFirstInRow: false, IsLastInRow: false},
		ui.Cell{Value: "DESCRIPTION", Length: ui.DescriptionColumnLength, IsFirstInRow: false, IsLastInRow: false},
		ui.Cell{Value: "CREATED AT", Length: ui.DateColumnsLength, IsFirstInRow: false, IsLastInRow: false},
		ui.Cell{Value: "UPDATED AT", Length: ui.DateColumnsLength, IsFirstInRow: false, IsLastInRow: true},
	}
	ui.DrawHorizontalLine()
	headerRow.Build()
	notes.ViewNotes()
}

func viewNote() {
	fmt.Println("Please enter the id of the note you'd like to view:")
	noteId := utils.GetNumberUserInput()
	if noteId == -1 || !notes.IsValidNoteId(noteId) {
		fmt.Println("Invalid note id")
	} else {
		note := notes.GetNoteById(noteId)
		note.Print()
	}
}

func createNote() {
	fmt.Println("Please enter the title:")
	title, isTitleOk := utils.GetTextUserInput()

	fmt.Println("Please enter the description:")
	description, isDescriptionOk := utils.GetTextUserInput()

	if isTitleOk && isDescriptionOk {
		notes.CreateNote(title, description)
		fmt.Println("Note successfully added!")
	} else {
		fmt.Println("Please enter valid values")
	}
}

func updateNote() {
	fmt.Println("Please enter the id of the note you'd like to update:")
	noteId := utils.GetNumberUserInput()
	if noteId == -1 || !notes.IsValidNoteId(noteId) {
		fmt.Println("Invalid note id")
	} else {
		note := notes.GetNoteById(noteId)
		note.Print()

		fmt.Println("Please enter the new title (Press enter to skip):")
		updatedTitle, isTitleUpdated := utils.GetTextUserInput()

		fmt.Println("Please enter the new description (Press enter to skip):")
		updatedDescription, isDescriptionUpdated := utils.GetTextUserInput()

		updatedNote := note

		if isTitleUpdated {
			updatedNote.SetTitle(updatedTitle)
		}
		if isDescriptionUpdated {
			updatedNote.SetDescription(updatedDescription)
		}

		if isTitleUpdated || isDescriptionUpdated {
			updatedNote = notes.UpdateNote(noteId, updatedNote)
			fmt.Println("Note successfully updated!")
			updatedNote.Print()
		}
	}
}

func deleteNote() {
	fmt.Println("Please enter the id of the note you'd like to delete:")
	noteId := utils.GetNumberUserInput()
	if noteId == -1 || !notes.IsValidNoteId(noteId) {
		fmt.Println("Invalid note id")
	} else {
		note := notes.GetNoteById(noteId)
		note.Print()
		fmt.Println("Are you sure you want to delete this note? (Yes / No)")
		confirmDelete, ok := utils.GetBooleanUserInput()
		if ok {
			if confirmDelete {
				notes.DeleteNote(noteId)
				fmt.Println("Note deleted successfully!")
			} else {
				fmt.Println("Cancelling ....")
			}
		} else {
			fmt.Println("Sorry something went wrong when deleting the note, try again")
		}
	}
}
