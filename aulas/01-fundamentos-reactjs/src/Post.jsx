// props: { author: "", content: "" }

export function Post(props) { 
    return (
        <>
            <strong>{ props.author }</strong>
            <p> { props.content } </p>
        </>
    )
}

// export default Post


// Default Exports vs Named Exports