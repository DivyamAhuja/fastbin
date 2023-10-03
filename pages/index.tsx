import { useCallback, useEffect, useRef } from 'react'
import type { NextPage } from 'next'
import Head from 'next/head'
import { useRouter } from 'next/router'

import styles from '../styles/Editor.module.css'
import header_styles from '../styles/Header.module.css'

import Save from '@material-ui/icons/Save'
import NoteAdd from '@material-ui/icons/NoteAdd'

const Home: NextPage = () => {

  const codeRef = useRef<HTMLTextAreaElement>(null)
  const router = useRouter()

  const save = useCallback(() => {
    fetch('/api/new', {
      'method': 'POST',
      'headers': {
        'Content-Type': 'application/json'
      },
      'body': JSON.stringify({ data: codeRef.current?.value })
    }).then(res => res.json())
      .then(({ id }) => router.push(`/${id}`))
      .catch(() => router.push('/'))
  }, [router])

  useEffect(() => {
    const listener = (event : KeyboardEvent) => {
      if (event.code === "KeyS" && event.ctrlKey === true) {
        event.preventDefault()
        save()
      }
    }

    document.addEventListener('keydown', listener)

    return () => {
      document.removeEventListener('keydown', listener)
    }
  }, [save])

  return (
    <div className={styles.container}>
      <Head>
        <title>fastbin</title>
        <meta name="description" content="fastbin: sharing code made faster" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <div className={header_styles["header"]}>
        <div className={header_styles["logo"]}> fastbin </div>
        <div className={header_styles["buttons-container"]}>
          <div
            className={header_styles["buttons"]}
            onClick={() => router.push('/')}
          >
            <NoteAdd />
          </div>
          <div
            className={header_styles["buttons"]}
            onClick={save}
          >
            <Save />
          </div>
        </div>
      </div>

      <div className={styles.editor}>
        <span className={styles["line-numbers"]}>
          {">"}
        </span>
        <textarea
          autoFocus
          wrap="off"
          ref={codeRef}
          placeholder={"Type Someting Here...\nCtrl + S to Save Document\nCtrl + N for New Document\n:)"}
          className={styles["code-editor"]}>
        </textarea>
      </div>
    </div>
  )
}

export default Home
