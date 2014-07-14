package store

//TODO melhorar o storage

import (
	"log"
	"bolt"
	"erevna/document"
	"code.google.com/p/go-uuid/uuid"
)

const SEQUENCE string = "SEQUENCE"
const UUID string = "UUID"
const DOC_BUCKET string = "DOC_BUCKET"
const WORD_BUCKET string = "WORD_BUCKET"

type Storage interface {
	Open() Storage
	Persist(doc *document.Document) *document.Document
	Search(criteria string)
	Close()
}

type StorageKeyPair struct {
	db *bolt.DB
}

func (st *StorageKeyPair) Open() Storage {
	
	db, err := bolt.Open("erevna.db", 0644, nil)	
	handleError(err);	
	st.db = db	
	
	return st
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createBase(db bolt.DB) {
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(DOC_BUCKET))
		return err
	})
	handleError(err);
}

func (st *StorageKeyPair) Persist( doc *document.Document ) *document.Document {
	
	uuid_ := uuid.New()
	
    err := st.db.Update(func(tx *bolt.Tx) error {
        
        doc_bucket, err := tx.CreateBucketIfNotExists([]byte(DOC_BUCKET))
        if err != nil {
            return err
        }
        //bucket de relacao entre uuids e docs
        internal_bucket, err := doc_bucket.CreateBucketIfNotExists([]byte(doc.Name))
        if err != nil {
            return err
        }
        internal_bucket.Put([]byte(uuid_), []byte(""))//associa um uuid ao doc name bucket "" value pq o valor eh irrelevante
        
        //bucket do doc pelo uuid   
        bucket, errd := tx.CreateBucketIfNotExists([]byte(uuid_))
        if errd != nil {
            return err
        }
		for i := 0; i < len(doc.Fields); i++ {
        	err = bucket.Put([]byte(doc.Fields[i].Name), []byte(doc.Fields[i].DataAsString()))
        }
		
		doc.Uuid = uuid_
		
		//Save content
        bucket_word, errw := tx.CreateBucketIfNotExists([]byte(WORD_BUCKET))
		handleError(errw);
		
		var doc_word_bucket *bolt.Bucket 
		
		word := doc.Content.NextToken()
		for word != "" {
        	doc_word_bucket, errw = bucket_word.CreateBucketIfNotExists([]byte(word))
			doc_word_bucket.Put([]byte(uuid_), []byte(""))
			
			word = doc.Content.NextToken()	
			handleError(errw);
		}		
			
        return nil
    })	
	
	handleError(err);
	
	return doc
}

func (st *StorageKeyPair) Search(criteria string) {
	
}

func (st *StorageKeyPair) Close() {
	st.db.Close()
}
