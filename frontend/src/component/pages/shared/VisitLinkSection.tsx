import React, { Component, ReactElement } from 'react';

import './CreateShortLinkSection.scss';
import { TextField } from '../../form/TextField';
import { Button } from '../../ui/Button';
import { ShortLinkUsage } from './ShortLinkUsage';
import { Section } from '../../ui/Section';
import { ShortLink } from '../../../entity/ShortLink';
import { UIFactory } from '../../UIFactory';
import { validateCustomAliasFormat } from '../../../validators/CustomAlias.validator';
import { IAppState } from '../../../state/reducers';
import { Store } from 'redux';
import { GraphQLService } from '../../../service/GraphQL.service';
import { ShortLinkService } from '../../../service/ShortLink.service';
import { QrCodeService } from '../../../service/QrCode.service';

interface IProps {
  store: Store<IAppState>;
  uiFactory: UIFactory;
  graphQLService: GraphQLService;
  shortLinkService: ShortLinkService;
  qrCodeService: QrCodeService;
  onShortLinkCreated?: (shortLink: string) => void;
  onAuthenticationFailed?: () => void;
}

interface IState {
  alias?: string;
  longLink?: string;
  room?: string;
  user?: string;
  inputError?: string;
  isShortLinkPublic?: boolean;
  shouldShowUsage: boolean;
  createdShortLink: string;
  createdLongLink: string;
  qrCodeURL: string;
  club: string;
  link: string;
  status: string;
}

export class VisitLinkSection extends Component<IProps, IState> {
  private shortLinkTextField = React.createRef<TextField>();

  constructor(props: IProps) {
    super(props);
    this.state = {
      shouldShowUsage: false,
      createdShortLink: '',
      createdLongLink: '',
      qrCodeURL: '',
      club: 'green',
      link: '',
      status: ''
    };
  }

  render(): ReactElement {
    return (
      <Section title={''}>
        <div className={'control visit-short-link'}>
          <h1>
            <span className={this.state.club}>club</span>
            <span className={this.state.link}>l</span>
            <span className={this.state.club}>.</span>
            <span className={this.state.link}>ink </span>
            <span className={'sand'}>/</span>
          </h1>
          <div className={'text-field-wrapper'}>
            <TextField
              className="code"
              ref={this.shortLinkTextField}
              text={this.state.alias}
              placeHolder={'enter code'}
              onBlur={this.handleCustomAliasTextFieldBlur}
              onChange={this.handleAliasChange}
            />
            {this.state.status === 'success' ? (
              <button
                className={'rocket-button'}
                onClick={async () => {
                  let qrCodeURL = await this.props.qrCodeService.newQrCode(
                    `${this.state.longLink}`
                  );
                  window.location.assign(
                    `/published/?alias=${this.state.alias}&longLink=${this.state.longLink}&shortLink=http://clubl.ink/${this.state.alias}&qrCodeURL=${qrCodeURL}`
                  );
                }}
              >
                🚀
              </button>
            ) : (
              <span
                className={`emoji ${!this.state.alias && 'hidden'}`}
                aria-hidden="true"
              >
                😩
              </span>
            )}
          </div>
        </div>
        <div className={'input-error'}>{this.state.inputError}</div>
        <div className={'input-description'}>
          {this.state.status === '' && 'Enter the super-secret code and go 🚀'}
          {this.state.status === 'error' &&
            "Code doesn't exist, try entering another"}
          {this.state.status === 'success' && (
            <>
              <div>
                <p>Imagine a link impossible to remember: </p>
                <a href={this.state.longLink} target="_blank">
                  {this.state.longLink}
                </a>
              </div>
              <div>
                <p>
                  by: <span className="bold">{this.state.user}</span> for room:{' '}
                  <span className="bold">{this.state.room}</span>
                </p>
              </div>
            </>
          )}
        </div>
        {/* {this.props.uiFactory.createPreferenceTogglesSubSection({
          uiFactory: this.props.uiFactory,
          isShortLinkPublic: this.state.isShortLinkPublic,
          onPublicToggleClick: this.handlePublicToggleClick
        })} */}
        {this.state.shouldShowUsage && (
          <div className={'short-link-usage-wrapper'}>
            <ShortLinkUsage
              shortLink={this.state.createdShortLink}
              longLink={this.state.createdLongLink}
              qrCodeUrl={this.state.qrCodeURL}
            />
          </div>
        )}
      </Section>
    );
  }

  handleAliasChange = async (newAlias: string) => {
    if (newAlias === '') {
      this.setState({
        alias: newAlias,
        club: 'green',
        link: '',
        status: ''
      });
    } else {
      this.setState({
        alias: newAlias,
        club: 'sand',
        link: ''
      });
      await this.handleCodeValidation(newAlias);
    }
  };

  handleCustomAliasTextFieldBlur = () => {
    const { alias } = this.state;
    const err = validateCustomAliasFormat(alias);
    this.setState({
      inputError: err || undefined
    });
  };

  handleCodeValidation = async (alias: string) => {
    await this.props.graphQLService
      .query('http://localhost:8080/graphql', {
        query: `query {
          authQuery {
            shortLink(alias: "${alias}") {
              id
              alias
              longLink
              room
              expireAt
            }
          }
        }`,
        variables: {}
      })
      .then(results => {
        let queryData: any = results;
        this.setState({
          link: 'green',
          status: 'success',
          longLink: queryData.authQuery.shortLink.longLink,
          room: queryData.authQuery.shortLink.room
        });
      })
      .catch(error => {
        this.setState({
          link: 'error',
          status: 'error'
        });
      });
  };

  handleCreateShortLinkClick = () => {
    // const { alias, longLink } = this.state;
    // const shortLink: ShortLink = {
    //   longLink: longLink,
    //   alias: alias || ''
    // };
    // this.props.shortLinkService
    //   .createShortLink(shortLink, this.state.isShortLinkPublic)
    //   .then(async (createdShortLink: ShortLink) => {
    //     const shortLink = this.props.shortLinkService.aliasToFrontendLink(
    //       createdShortLink.alias!
    //     );
    //     const qrCodeURL = await this.props.qrCodeService.newQrCode(shortLink);
    //     this.setState({
    //       createdShortLink: shortLink,
    //       createdLongLink: longLink,
    //       qrCodeURL: qrCodeURL,
    //       shouldShowUsage: true
    //     });
    //     if (this.props.onShortLinkCreated) {
    //       this.props.onShortLinkCreated(shortLink);
    //     }
    //   })
    //   .catch(({ authenticationErr, createShortLinkErr }) => {
    //     if (authenticationErr) {
    //       if (this.props.onAuthenticationFailed) {
    //         this.props.onAuthenticationFailed();
    //       }
    //       return;
    //     }
    //     this.props.store.dispatch(
    //       raiseCreateShortLinkError(createShortLinkErr)
    //     );
    //   });
  };

  handlePublicToggleClick = (enabled: boolean) => {
    this.setState({
      isShortLinkPublic: enabled
    });
  };

  focusShortLinkTextField = () => {
    if (!this.shortLinkTextField.current) {
      return;
    }
    this.shortLinkTextField.current.focus();
  };
}
