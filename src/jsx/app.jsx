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

React.render(<Monoton/>, document.body);
