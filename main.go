package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/iterator"

	"log"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	//"google.golang.org/api/option"
	"github.com/gofiber/fiber/v2/middleware/cors"

	
)
const (
	DDMMYYYYhhmmss = "2006-01-02 15:04:05"
)
type Test struct {
	Co string `json:"comp"`
	Te string `json:"test"`
}
type AllTestResponse struct {
	Tests []map[string]interface{}
}

func (box *AllTestResponse) AddItem(item map[string]interface{}) {
	box.Tests = append(box.Tests, item)
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	//https://brain-sugar.el.r.appspot.com/quiz/1420401705646271b29916b1684173234
	
	app.Get("/exam/:id", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		//msg := fmt.Sprintf("👴 %s is %s years old", c.Params("name"), c.Params("age"))

		ctx := context.Background()
		conf := &firebase.Config{ProjectID: "brain-sugar"}
		app, err := firebase.NewApp(ctx, conf)
		if err != nil {
			log.Fatalln(err)
		}

		client, err := app.Firestore(ctx)
		if err != nil {
			log.Fatalln(err)
		}

		if err != nil {
			log.Fatalln(err)
		}

		

		dsnap, err := client.Collection("quizz2").Doc(c.Params("id")).Get(ctx)
						if err != nil {
								
						}
						m := dsnap.Data()
						return c.JSON(m)

		

	})

	
	app.Get("/questions/:quizId", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		//msg := fmt.Sprintf("👴 %s is %s years old", c.Params("name"), c.Params("age"))

		ctx := context.Background()
		conf := &firebase.Config{ProjectID: "brain-sugar"}
		app, err := firebase.NewApp(ctx, conf)
		if err != nil {
			log.Fatalln(err)
		}

		client, err := app.Firestore(ctx)
		if err != nil {
			log.Fatalln(err)
		}

		if err != nil {
			log.Fatalln(err)
		}

		box := AllTestResponse{}

		iter := client.Collection("quizz2").Where("customer_id", "==", c.Params("studentId")).Documents(ctx)

		defer iter.Stop()

		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err == nil {
				fmt.Println(doc.Data())
				oneDoc := doc.Data()
				id := oneDoc["course_id"].(string)
				fmt.Println(id)
				s := []string{"course_id", "exam_end", "exam_start", "exam_time", "id", "num_retakes", "pass_mark", "retakes", "section_details", "status", "title", "total_point"}
				iter2 := client.Collection("quizz2").Select(s...).Where("course_id", "==", id).Documents(ctx)
				defer iter2.Stop()

				for {
					doc2, err2 := iter2.Next()
					if err2 == iterator.Done {
						break
					}
					if err2 == nil {

						box.AddItem(doc2.Data())

					} else {
						fmt.Println(err2)
					}

				}

				//course_id
				//box.AddItem(doc.Data())

			} else {
				fmt.Println(err)
			}

		}
		//fmt.Fprint(w, s)

		return c.JSON(box.Tests)

	})

	//https://brain-sugar.el.r.appspot.com/quiz/1420401705646271b29916b1684173234
	app.Get("/quiz/:studentId", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		//msg := fmt.Sprintf("👴 %s is %s years old", c.Params("name"), c.Params("age"))

		ctx := context.Background()
		conf := &firebase.Config{ProjectID: "brain-sugar"}
		app, err := firebase.NewApp(ctx, conf)
		if err != nil {
			log.Fatalln(err)
		}

		client, err := app.Firestore(ctx)
		if err != nil {
			log.Fatalln(err)
		}

		if err != nil {
			log.Fatalln(err)
		}

		box := AllTestResponse{}
		boxImportant := AllTestResponse{}

		iter := client.Collection("apply_course").Where("customer_id", "==", c.Params("studentId")).Documents(ctx)

		defer iter.Stop()

		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err == nil {
				fmt.Println(doc.Data())
				oneDoc := doc.Data()
				id := oneDoc["course_id"].(string)
				fmt.Println(id)
				//now := time.Now().UTC()
				//sT := now.Format(DDMMYYYYhhmmss)
				now := time.Now()
				var num int64
				num = now.UnixMilli()
				s := []string{"course_id", "exam_end", "exam_start", "exam_time", "id", "num_retakes", "pass_mark", "retakes", "section_details", "status", "title", "total_point","created_at"}
				iter2 := client.Collection("quizz2").Select(s...).Where("course_id", "==", id).Documents(ctx)
				defer iter2.Stop()

				for {
					doc2, err2 := iter2.Next()
					if err2 == iterator.Done {
						break
					}
					if err2 == nil {

						

						oneQuiz := doc2.Data()

						xmStart := oneQuiz["exam_start"].(int64)
						if xmStart<num {
							oneQuiz["hasStarted"] = true;

						}else{
							oneQuiz["hasStarted"] = false;
						}
						xmEnd := oneQuiz["exam_end"].(int64)
						if xmEnd>num {
							
							oneQuiz["hasEnded"] = false;

						}else{
							oneQuiz["hasEnded"] = true;
						}
						



							oneQuiz["can_take_exam"] = true;
							//oneQuiz["test_date"] = cfcf
							
							oneQuiz["quiz_id"] =  doc2.Ref.ID;
							oneQuiz["totalExamAppeadred"] = 0;
	
							//str := oneQuiz["exam_start"].(string)
							
							
							
							
	
							
							oneQuiz["ttt"] = num;
	
	
	
	
							
	
							dsnap, err := client.Collection("courses").Doc(id).Get(ctx)
							if err != nil {
									
							}
							m := dsnap.Data()
							course_title := m["course_title"].(string)
							oneQuiz["course_title"] = course_title;
	
	
							
							quiz_id :=  doc2.Ref.ID;
							query_d := "submitionCount-"+c.Params("studentId")
							
							
							docRef2 := client.Collection(query_d).Doc(quiz_id)
							doc3, err := docRef2.Get(ctx)
							if err != nil {
								
							}
							
							// Double handling (???)
							if doc3.Exists() {
								// Handle document existing here
								m2 := doc3.Data()
							
								exam_appreared := m2["count"].(int64)
								oneQuiz["totalExamAppeadred"] = exam_appreared;
								
								if oneQuiz["hasStarted"] == true {
									if oneQuiz["hasEnded"] == false {
										boxImportant.AddItem(oneQuiz)
									
									}else {
										box.AddItem(oneQuiz)
									}
								}else {
									box.AddItem(oneQuiz)
								}
								
							}else {
								if oneQuiz["hasStarted"]  == true {
									if oneQuiz["hasEnded"] == false {
										boxImportant.AddItem(oneQuiz)
									
									}else {
										box.AddItem(oneQuiz)
									}
								}else {
									box.AddItem(oneQuiz)
								}
							}

						
						
						
						
	
						
						
						

						


						

					} else {
						fmt.Println(err2)
					}

				}

				//course_id
				//box.AddItem(doc.Data())

			} else {
				fmt.Println(err)
			}

		}
		for i, s := range box.Tests {
			fmt.Println(i, s)
			boxImportant.AddItem(s)
		}
		
		return c.JSON(boxImportant.Tests)

	})
	app.Get("/results/:studentId", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		//msg := fmt.Sprintf("👴 %s is %s years old", c.Params("name"), c.Params("age"))

		ctx := context.Background()
		conf := &firebase.Config{ProjectID: "brain-sugar"}
		app, err := firebase.NewApp(ctx, conf)
		if err != nil {
			log.Fatalln(err)
		}

		client, err := app.Firestore(ctx)
		if err != nil {
			log.Fatalln(err)
		}

		if err != nil {
			log.Fatalln(err)
		}

		box := AllTestResponse{}
		s := []string{"title","total_point","exam_start","exam_end","created_at","student_id","obtainedMarks","pass_mark"}

		iter := client.Collection("submition").Select(s...).OrderBy("created_at",firestore.Desc).Documents(ctx)
		now := time.Now()
		var num int64
		num = now.UnixMilli()
		defer iter.Stop()

		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err == nil {

				oneDoc := doc.Data()
				oneDoc["id"] = doc.Ref.ID 
				student_id := oneDoc["student_id"].(string)
				endTime := oneDoc["exam_end"].(int64)
				if student_id == c.Params("studentId") {

					if endTime> num {
						oneDoc["exmTimeFinished"] = false;

					}else{
						oneDoc["exmTimeFinished"] = true;
					}
					oneDoc["sysTime"] = num;

					fmt.Println(doc.Data())
					box.AddItem(oneDoc)

				


				}
				

			} else {
				fmt.Println(err)
			}

		}
		//fmt.Fprint(w, s)

		return c.JSON(box.Tests)

	})
	app.Get("/submit/:id", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		//msg := fmt.Sprintf("👴 %s is %s years old", c.Params("name"), c.Params("age"))

		ctx := context.Background()
		conf := &firebase.Config{ProjectID: "brain-sugar"}
		app, err := firebase.NewApp(ctx, conf)
		if err != nil {
			log.Fatalln(err)
		}

		client, err := app.Firestore(ctx)
		if err != nil {
			log.Fatalln(err)
		}

		if err != nil {
			log.Fatalln(err)
		}

		dsnap, err := client.Collection("submition").Doc(c.Params("id")).Get(ctx)
		if err != nil {

		}
		m := dsnap.Data()
		return c.JSON(m)
		

		

	})

	// GET /api/register
	app.Get("/api/*", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ %s", c.Params("*"))
		return c.SendString(msg) // => ✋ register
	})

	// GET /flights/LAX-SFO
	app.Get("/flights/:from-:to", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("💸 From: %s, To: %s", c.Params("from"), c.Params("to"))
		return c.SendString(msg) // => 💸 From: LAX, To: SFO
	})

	// GET /dictionary.txt
	app.Get("/:file.:ext", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("📃 %s.%s", c.Params("file"), c.Params("ext"))
		return c.SendString(msg) // => 📃 dictionary.txt
	})


	app.Get("/:name", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
		return c.SendString(msg) // => Hello john 👋!
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	app.Listen(fmt.Sprintf(":%s", port))

}
