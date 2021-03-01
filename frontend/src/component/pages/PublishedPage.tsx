import React, { Component } from 'react';
import './PublishedPage.scss';
import { ShortLinkUsage } from './shared/ShortLinkUsage';
import { GraphQLService } from '../../service/GraphQL.service';
import { QrCodeService } from '../../service/QrCode.service';

interface Props {
  graphQLService: GraphQLService;
  qrCodeService: QrCodeService;
}

interface State {
  alias?: string;
  longLink?: string;
  qrCode?: string;
  room?: string;
  user?: string;
  status?: string;
}
export class PublishedPage extends Component<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = {};
  }

  componentDidMount() {
    let shortLink;
    let urlData: URLSearchParams = new URLSearchParams(window.location.search);
    this.props.graphQLService
      .query('http://localhost:8080/graphql', {
        query: `query {
        authQuery {
          shortLink(id:"${urlData.get('link')}"){
            alias
            longLink
            room
          }
          userByShortLink(id:"${urlData.get('link')}"){
            name
          }
        }
      }`,
        variables: {}
      })
      .then(async results => {
        let data: any = results;
        let qrCodeURL = await this.props.qrCodeService.newQrCode(
          `${data.authQuery.shortLink.longLink}`
        );
        this.setState({
          alias: data.authQuery.shortLink.alias,
          longLink: data.authQuery.shortLink.longLink,
          room: data.authQuery.shortLink.room,
          user: data.authQuery.userByShortLink.name,
          qrCode: qrCodeURL,
          status: 'success'
        });
      })
      .catch(error => {
        this.setState({
          status: 'error'
        });
      });
  }

  render = () => {
    return (
      <div className="published">
        {this.state.status === 'success' && (
          <div className="short-link-usage-wrapper">
            <ShortLinkUsage
              shortLink={`${this.state.alias}`}
              longLink={`${this.state.longLink}`}
              qrCodeUrl={`${this.state.qrCode}`}
            />
          </div>
        )}
        <a href={`${this.state.longLink}`} className="heading" target="_blank">
          <h1 aria-label="clublink/luffy">
            <span aria-hidden>🚀</span>
            <span className="lightGreen">club</span>
            <span className="darkGreen">l</span>
            <span className="lightGreen">.</span>
            <span className="darkGreen">ink</span>/{this.state.alias}
          </h1>
        </a>
        <p className="imagine">Imagine a link impossible to remember:</p>
        <a href={`${this.state.longLink}`} target="_blank">
          {this.state.longLink}
        </a>
        <p>
          by: <span className="bold">{this.state.user}</span> for room:{' '}
          <span className="bold">{this.state.room}</span>
        </p>
        <div className="buttons">
          <button
            onClick={() => {
              navigator.clipboard.writeText(`${this.state.alias}`).then(
                function() {
                  /* clipboard successfully set */
                  let button = document.querySelector('button');

                  if (button) {
                    button.innerHTML = 'copied';
                  }
                },
                function() {
                  /* clipboard write failed */
                }
              );
            }}
          >
            copy club-link
          </button>
          <a className="button" href="/a/create">
            create new
          </a>
        </div>
      </div>
    );
  };
}
