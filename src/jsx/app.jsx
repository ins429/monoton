import React from 'react';

class Monoton extends React.Component {
    getInitialState() {
        return {};
    }

    componentWillMount() {
    }

    render() {
        return (
            <section>
                <div id="stage" className="container-fluid">
                    <Grid/>
                    <Form/>
                </div>
            </section>
        );
    }
}

class Grid extends React.Component {
    getInitialState() {
        return {};
    }

    componentWillMount() {
    }

    render() {
        let image = (
            <div className="image">
                <Img src=""/>
            </div>
        )
        return (
            <div className="images">
                {image}
            </div>
        );
    }
}

class Img extends React.Component {
    getInitialState() {
        return {};
    }

    componentWillMount() {
    }

    render() {
        return (
            <div>
                <img src="this.props.src" alt="" />
            </div>
        );
    }
}

class Form extends React.Component {
    getInitialState() {
        return {};
    }

    componentWillMount() {
    }

    render() {
        return (
            <div>
                <form action="/photos" method="post" encType="multipart/form-data">
                    <input type="file" name="photo" />
                    <input type="hidden" value="foo" name="foo" />
                    <input type="submit" value="Upload Image" name="submit" />
                </form>
            </div>
        );
    }
}

React.render(<Monoton/>, document.body);
