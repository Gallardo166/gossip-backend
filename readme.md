# Gossip

Name: Cheneil Gallardo Lee

Link to deployed app: <https://gossip-frontend-nine.vercel.app/>

Frontend Codebase: <https://github.com/Gallardo166/gossip-frontend>

Image credits:
Photo by <a href="https://unsplash.com/@von_co?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash">Ivana Cajina</a> on <a href="https://unsplash.com/photos/milky-way-asuyh-_ZX54?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash">Unsplash</a>, Photo by <a href="https://unsplash.com/@damiano_baschiera?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash">Damiano Baschiera</a> on <a href="https://unsplash.com/photos/bed-of-orange-flowers-d4feocYfzAM?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash">Unsplash</a>

  --------------------------------------------------------

Or run locally:

## 1. Set up database

- Download PostgreSQL: <https://www.postgresql.org/download/>
- The setup installer will prompt you to enter a password
- After installing, open pgAdmin 4 and click on 'Servers' on the left tab
- Click on 'PostgreSQL 17' and enter the password from the previous step
- Right click 'Databases' and create a new database, entering the database name
- Open 'Schemas' and 'public', right click on 'Tables', and click 'Query Tool"
- Run the following script to create the tables:
  
     ```sql
     CREATE TABLE IF NOT EXISTS public.categories
     (
        id integer NOT NULL,
        name character varying(100) COLLATE pg_catalog."default" NOT NULL,
        CONSTRAINT categories_pkey PRIMARY KEY (id)
     )
      
     TABLESPACE pg_default;
      
     ALTER TABLE IF EXISTS public.categories
       OWNER to postgres;

     CREATE TABLE IF NOT EXISTS public.users
     (
        id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
        username character varying(100) COLLATE pg_catalog."default" NOT NULL,
        password character varying(100) COLLATE pg_catalog."default" NOT NULL,
        CONSTRAINT "Users_pkey" PRIMARY KEY (id)
     )
      
     TABLESPACE pg_default;
      
     ALTER TABLE IF EXISTS public.users
     OWNER to postgres;
     
     CREATE TABLE IF NOT EXISTS public.posts
     (
        id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
        title character varying(200) COLLATE pg_catalog."default" NOT NULL,
        body character varying(3000) COLLATE pg_catalog."default" NOT NULL,
        image_url text COLLATE pg_catalog."default",
        category_id integer NOT NULL,
        user_id integer NOT NULL,
        date timestamp without time zone NOT NULL,
        CONSTRAINT "Posts_pkey" PRIMARY KEY (id),
        CONSTRAINT "Posts_category_id_fkey" FOREIGN KEY (category_id)
            REFERENCES public.categories (id) MATCH SIMPLE
            ON UPDATE NO ACTION
            ON DELETE NO ACTION
            NOT VALID,
        CONSTRAINT "Posts_user_id_fkey" FOREIGN KEY (user_id)
            REFERENCES public.users (id) MATCH SIMPLE
            ON UPDATE NO ACTION
            ON DELETE NO ACTION
            NOT VALID
     )
        
     TABLESPACE pg_default;
        
     ALTER TABLE IF EXISTS public.posts
       OWNER to postgres;
    
    CREATE TABLE IF NOT EXISTS public.post_likes
    (
        post_id integer NOT NULL,
        user_id integer NOT NULL,
        CONSTRAINT post_likes_post_id_fkey FOREIGN KEY (post_id)
            REFERENCES public.posts (id) MATCH SIMPLE
            ON UPDATE NO ACTION
            ON DELETE NO ACTION,
        CONSTRAINT post_likes_user_id_fkey FOREIGN KEY (user_id)
            REFERENCES public.users (id) MATCH SIMPLE
            ON UPDATE NO ACTION
            ON DELETE NO ACTION
    )
        
    TABLESPACE pg_default;
        
    ALTER TABLE IF EXISTS public.post_likes
      OWNER to postgres;
    
     CREATE TABLE IF NOT EXISTS public.comments
     (
        id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
        body character varying(800) COLLATE pg_catalog."default" NOT NULL,
        user_id integer NOT NULL,
        post_id integer NOT NULL,
        date timestamp without time zone NOT NULL,
        parent_id integer,
        CONSTRAINT comments_pkey PRIMARY KEY (id),
        CONSTRAINT "Comments_post_id_fkey" FOREIGN KEY (post_id)
            REFERENCES public.posts (id) MATCH SIMPLE
            ON UPDATE NO ACTION
            ON DELETE NO ACTION
            NOT VALID,
        CONSTRAINT "Comments_user_id_fkey" FOREIGN KEY (user_id)
            REFERENCES public.users (id) MATCH SIMPLE
            ON UPDATE NO ACTION
            ON DELETE NO ACTION,
        CONSTRAINT comments_parent_id_fkey FOREIGN KEY (parent_id)
            REFERENCES public.comments (id) MATCH SIMPLE
            ON UPDATE NO ACTION
            ON DELETE NO ACTION
            NOT VALID
     )
      
     TABLESPACE pg_default;
      
     ALTER TABLE IF EXISTS public.comments
       OWNER to postgres;
     ```

- Right click 'Tables' and click 'Refresh' for the tables to appear
- Run the following scripts to create some categories:
  
     ```sql
     INSERT INTO public.categories(
       id, name)
     VALUES (1, 'Science');
     ```

     --------------------------------------------------------

     ```sql
     INSERT INTO public.categories(
       id, name)
     VALUES (2, 'Technology');
     ```

- Right click 'PostgreSQL 17', click on 'Properties', and under the 'Connection' tab note the host name/address

## 2. Set up backend

- Clone this repository: `git clone https://github.com/Gallardo166/gossip-backend.git`
- In the backend root directory, run `go get github.com/joho/godotenv` in the terminal
- Replace the code in `db.go` with

    ```go
    package initializers

    import (
      "log"
      "os"

      "github.com/jmoiron/sqlx"
      "github.com/joho/godotenv"
      _ "github.com/lib/pq"
    )

    var DB *sqlx.DB

    func ConnectDB() {
      envErr := godotenv.Load(".env")
      if envErr != nil {
        log.Fatalf("Error loading .env file: %s", envErr)
      }
      connStr := os.Getenv("CONNSTR")
      var dbErr error
      DB, dbErr = sqlx.Connect("postgres", connStr)
      if dbErr != nil {
        log.Fatalf("Error connecting to database: %s", dbErr)
      }
    }
    ```

- In the backend root directory, run `go mod tidy`
- Create a cloudinary account: <https://cloudinary.com/users/register_free>

- Skip introduction and go to Settings
- Under 'Product environment settings' click 'API Keys'
- Copy the API environment variable `cloudinary://<your_api_key>:<your_api_secret>@<your_cloud_name>`, replacing the necessary fields
- In the backend root directory, create an `.env` file with the following keys, replacing the necessary fields:

    ```js
    CONNSTR="user=postgres dbname=<your_db_name> host=<db_host_name> password=<your_pgadmin_password> sslmode=disable"
    JWTKEY="<any_string>"
    CLOUDINARY_URL="<your_api_key>:<your_api_secret>@<your_cloud_name>"
    ```

- In the backend root directory, run `go run .` for the app to listen

## 3. Set up frontend

- Clone the frontend repository: `git clone https://github.com/Gallardo166/gossip-backend.git`
- In the frontend root directory, run `npm install`
- In the frontend root directory, create a `.env` file with the following keys:

    ```js
    VITE_URL="http://localhost:3000"
    ```

- Run `npm run dev` and press `o + enter`
